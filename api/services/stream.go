package service

import (
	"Twitcher/pb"
	"Twitcher/twitchApi"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/go-zeromq/zmq4"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataChannelMsg struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
}

type Song struct {
	Page          string
	Name          string
	Author        string
	AudioFilename string
	CoverFilename string
	Duration      float64
	Bitrate       int
}

type Status struct {
	audio   bool
	preview bool
	stream  bool
	finding bool
}

type StreamParams struct {
	preset string
	width  int
	height int
	fps    int
	volume float64
}

type Channels struct {
	swapBackgroundVideo    chan struct{}
	sendOverlayDataChannel chan string
	sendMsgDataChannel     chan string
	sdpFromClient          chan string
	sdpForClient           chan string
	stopPreview            chan struct{}
	stopAudio              chan struct{}
	stopStream             chan struct{}
}

type MainServer struct {
	pb.UnimplementedMainServer
	playlist               pb.SongPlaylist
	currentSong            pb.Song
	currentBackgroundVideo pb.BackgroundVideo
	overlays               []OverlayObject
	mu                     sync.Mutex
	status                 Status
	streamParams           StreamParams
	channels               Channels
}

func (s *MainServer) Status(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Audio: s.status.audio, Preview: s.status.preview, Stream: s.status.stream}, nil
}

func (s *MainServer) StartPreview(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusResponse, error) {
	params, err := getStreamParams()
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.streamParams = params
	s.channels.swapBackgroundVideo = make(chan struct{})
	s.channels.sendOverlayDataChannel = make(chan string, 1)
	s.channels.sendMsgDataChannel = make(chan string, 300)
	s.channels.sdpFromClient = make(chan string, 300)
	s.channels.sdpForClient = make(chan string, 300)
	s.channels.stopPreview = make(chan struct{})
	s.channels.stopAudio = make(chan struct{})
	s.mu.Unlock()

	err = manageNamedPipes()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// audio
	go s.initAudio(&wg)

	// background video
	bgR, bgW := io.Pipe()
	go s.initBackgroundVideo(bgW, &wg)

	// Stream go routine.
	go s.initPreview(bgR, &wg)

	wg.Wait()

	// webRTC connection
	go s.Broadcast(s.channels.sdpFromClient, s.channels.sdpForClient, s.channels.stopPreview)

	return &pb.StatusResponse{Preview: s.status.preview, Audio: s.status.audio, Stream: s.status.stream}, nil
}

func (s *MainServer) StopPreview(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusResponse, error) {
	log.Println("stop preview")
	s.mu.Lock()
	s.status.preview = false
	s.status.stream = false
	s.status.audio = false
	s.mu.Unlock()

	close(s.channels.stopPreview)
	s.channels.stopAudio <- struct{}{}
	close(s.channels.swapBackgroundVideo)
	close(s.channels.sendOverlayDataChannel)
	close(s.channels.sdpFromClient)
	close(s.channels.sdpForClient)

	return &pb.StatusResponse{Preview: s.status.preview, Audio: s.status.audio, Stream: s.status.stream}, nil
}

func (s *MainServer) Preview(ctx context.Context, in *pb.SDP) (*pb.SDP, error) {
	if !s.status.audio {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio not available to show preview"),
		)
	}

	s.channels.sdpFromClient <- in.Sdp

	sdpForClient := <-s.channels.sdpForClient

	return &pb.SDP{Sdp: sdpForClient}, nil
}

func (s *MainServer) StartStream(ctx context.Context, in *google_protobuf.Empty) (*pb.StreamResponse, error) {
	if s.status.stream {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("stream has already started."),
		)
	}

	params, err := getStreamParams()
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.status.audio = true
	s.streamParams = params
	s.channels.swapBackgroundVideo = make(chan struct{})
	s.channels.sendOverlayDataChannel = make(chan string, 1)
	s.channels.sendMsgDataChannel = make(chan string, 300)
	s.channels.sdpFromClient = make(chan string, 300)
	s.channels.sdpForClient = make(chan string, 300)
	s.channels.stopAudio = make(chan struct{})
	s.channels.stopStream = make(chan struct{})
	s.mu.Unlock()

	// Enable alert notifications
	var wgAlerts sync.WaitGroup
	wgAlerts.Add(1)

	exitAlerts := make(chan struct{})

	go twitchApi.Alerts(&wgAlerts, exitAlerts)

	wgAlerts.Wait()
	// Get ingest server link
	twitchStreamLink, err := twitchIngestLink()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(
				codes.FailedPrecondition,
				fmt.Sprintln("stream key not found."),
			)
		}

		return nil, err
	}

	err = manageNamedPipes()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// audio
	go s.initAudio(&wg)

	// Background video.
	bgR, bgW := io.Pipe()
	go s.initBackgroundVideo(bgW, &wg)

	// Stream go routine.
	go s.initStream(bgR, twitchStreamLink, exitAlerts, &wg)

	wg.Wait()

	// webRTC connection
	go s.Broadcast(s.channels.sdpFromClient, s.channels.sdpForClient, s.channels.stopStream)

	return &pb.StreamResponse{Volume: s.streamParams.volume, Status: &pb.StatusResponse{Stream: s.status.stream, Audio: s.status.audio, Preview: s.status.preview}}, nil
}

func (s *MainServer) StopStream(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusResponse, error) {
	s.mu.Lock()
	s.status.preview = false
	s.status.stream = false
	s.status.audio = false
	s.mu.Unlock()

	close(s.channels.stopStream)
	s.channels.stopAudio <- struct{}{}
	close(s.channels.swapBackgroundVideo)
	close(s.channels.sendOverlayDataChannel)
	close(s.channels.sdpFromClient)
	close(s.channels.sdpForClient)

	return &pb.StatusResponse{Stream: s.status.stream, Audio: s.status.audio, Preview: s.status.preview}, nil
}

func (s *MainServer) CreateSongPlaylist(ctx context.Context, in *google_protobuf.Empty) (*pb.SongPlaylist, error) {

	playlist, err := s.generateRandomPlaylist()
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (s *MainServer) CurrentSongPlaylist(ctx context.Context, in *google_protobuf.Empty) (*pb.SongPlaylist, error) {
	return &s.playlist, nil
}

func (s *MainServer) UpdateSongPlaylist(ctx context.Context, in *pb.SongPlaylist) (*google_protobuf.Empty, error) {

	s.mu.Lock()

	s.playlist.Songs = nil
	s.playlist.Songs = append(s.playlist.Songs, in.Songs...)

	s.mu.Unlock()

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) GetOverlays(ctx context.Context, in *google_protobuf.Empty) (*pb.Overlays, error) {
	if len(s.overlays) == 0 {
		return &pb.Overlays{}, nil
	}

	var overlays []*pb.Overlay

	for _, overlay := range s.overlays {
		overlays = append(overlays, &pb.Overlay{
			Id:         overlay.Id,
			Type:       overlay.Type,
			Width:      int32(overlay.Width),
			Height:     int32(overlay.Height),
			PointX:     int32(overlay.PointX),
			PointY:     int32(overlay.PointY),
			Show:       overlay.Show,
			CoverId:    overlay.CoverId,
			Text:       overlay.Text,
			FontFamily: overlay.FontFamily,
			FontSize:   int32(overlay.FontSize),
			LineHeight: overlay.LineHeight,
			TextColor:  overlay.TextColor,
			TextAlign:  overlay.TextAlign,
		})
	}

	return &pb.Overlays{Overlays: overlays}, nil
}

func (s *MainServer) BackgroundVideos(ctx context.Context, in *google_protobuf.Empty) (*pb.BackgroundVideosResponse, error) {
	var videos []*pb.BackgroundVideo

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, active FROM background_videos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video pb.BackgroundVideo

		err = rows.Scan(&video.Id, &video.Name, &video.Active)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		videos = append(videos, &pb.BackgroundVideo{Id: video.Id, Name: video.Name, Active: video.Active})
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.BackgroundVideosResponse{Videos: videos}, nil
}

func (s *MainServer) SwapBackgroundVideo(ctx context.Context, in *pb.BackgroundVideo) (*google_protobuf.Empty, error) {
	s.mu.Lock()
	s.currentBackgroundVideo = *in
	s.mu.Unlock()

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE background_videos SET active = 0 WHERE active = 1")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE background_videos SET active = 1 WHERE id = $1", in.Id)
	if err != nil {
		return nil, err
	}

	if s.status.stream || s.status.preview {
		s.channels.swapBackgroundVideo <- struct{}{}
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) UploadVideo(stream pb.Main_UploadVideoServer) error {
	var (
		file     *os.File
		fileName string
	)

	defer file.Close()

	i := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Create tables
			db, err := sql.Open("sqlite3", "data.db")
			if err != nil {
				log.Println(err)
				return err
			}

			defer db.Close()
			_, err = db.Exec(`
				INSERT INTO background_videos (name, active) VALUES ($1, 0)
			`, fileName)
			if err != nil {
				return err
			}

			return stream.SendAndClose(&pb.UploadVideoResponse{Id: 1})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if i == 0 {
			fileName = req.GetInfo().GetFileName()

			file, err = os.OpenFile("files/stream/background-videos/"+fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0744)
			if err != nil {
				file.Close()
				if errors.Is(err, fs.ErrExist) {
					log.Println(fileName + " already exists.")
					return status.Error(codes.AlreadyExists, err.Error())
				}

				return err
			}

			i += 1
		} else {
			_, err = file.Write(req.GetChunk())
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}
	}

}

func (s *MainServer) DeleteBackgroundVideo(ctx context.Context, in *pb.BackgroundVideo) (*google_protobuf.Empty, error) {
	if s.status.preview || s.status.stream && s.currentBackgroundVideo.Id == in.Id {
		return &google_protobuf.Empty{}, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("Can't delete active background video while preview and/or stream are on."),
		)
	}

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	defer db.Close()

	if in.Active {
		_, err := db.Exec("UPDATE background_videos SET active = 1 WHERE id != $1 LIMIT 1", in.Id)
		if err != nil {
			return nil, err
		}
	}

	_, err = db.Exec("DELETE FROM background_videos WHERE id = $1", in.Id)
	if err != nil {
		return nil, err
	}

	err = os.Remove("files/stream/background-videos/" + in.Name)
	if err != nil {
		return nil, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) StreamParameters(ctx context.Context, in *google_protobuf.Empty) (*pb.StreamParametersResponse, error) {
	if s.streamParams.preset == "" {
		params, err := getStreamParams()
		if err != nil {
			return nil, err
		}

		s.mu.Lock()
		s.streamParams = params
		s.mu.Unlock()
	}

	if s.streamParams.volume == 0 {
		s.streamParams.volume = 101
	}

	return &pb.StreamParametersResponse{Width: int32(s.streamParams.width), Height: int32(s.streamParams.height), Fps: int32(s.streamParams.fps), Preset: s.streamParams.preset, Volume: s.streamParams.volume}, nil
}

func (s *MainServer) initPreview(r *io.PipeReader, wg *sync.WaitGroup) error {
	defer log.Println("preview stopped")

	cmd := exec.Command("ffmpeg", "-hide_banner", "-y",
		"-r", strconv.Itoa(s.streamParams.fps),
		"-stream_loop", "-1",
		"-i", "pipe:0",
		"-i", "files/stream/audio",

		"-map", "1:0",
		"-c:a", "copy", "-page_duration", "2000", "-f", "ogg", "files/stream/previewAudio",

		"-map", "0:0",
		"-c:v", "copy", "-f", "h264", "files/stream/previewVideo",
	)

	cmd.Stdin = r
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	wg.Done()

	s.mu.Lock()
	s.status.preview = true
	s.mu.Unlock()

	// Wait until stop command is received.
	<-s.channels.stopPreview

	s.mu.Lock()
	s.status.preview = false
	s.status.audio = false
	s.mu.Unlock()

	log.Println("stopping preview")
	err := cmd.Process.Signal(syscall.SIGKILL)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	r.Close()

	cmd.Wait()

	return nil
}

func (s *MainServer) initStream(r *io.PipeReader, twitchStreamLink string, exitAlerts chan<- struct{}, wg *sync.WaitGroup) error {
	defer log.Println("stream stopped")

	cmd := exec.Command("ffmpeg", "-hide_banner", "-y",
		"-r", strconv.Itoa(s.streamParams.fps),
		"-stream_loop", "-1",
		"-i", "pipe:0",
		"-thread_queue_size", "8", "-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
		"-thread_queue_size", "8", "-i", "files/stream/alert",
		"-filter_complex", `[0][1]overlay=0:0[v1];[v1][2]overlay=W-w+10:H-h+60,zmq[vout]`,
		"-i", "files/stream/audio",

		"-map", "3:0",
		"-c:a", "copy", "-page_duration", "2000", "-f", "ogg", "files/stream/previewAudio",

		"-map", "0:0",
		"-c:v", "copy", "-f", "h264", "files/stream/previewVideo",

		"-map", "[vout]",
		"-c:v", "libx264",
		"-fps_mode", "passthrough",
		"-preset", s.streamParams.preset, "-maxrate", "6M", "-bufsize", "3M",
		"-map", "3:0",
		"-c:a", "libmp3lame", "-q:a", "0",
		"-af", `volume=`+strconv.FormatFloat(s.streamParams.volume, 'f', -1, 64)+`,volume@foo,azmq=bind_address=tcp\\\://0.0.0.0\\\:5554`,
		"-f", "fifo", "-fifo_format", "flv", // Fifo muxer implemented to recover stream in case a failure occurs.
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header",
		"-g", strconv.Itoa(s.streamParams.fps*2),
		"-keyint_min", strconv.Itoa(s.streamParams.fps*2), "-force_key_frames", "expr:gte(t,n_forced*2)",
		"-flvflags", "no_duration_filesize",
		twitchStreamLink,
		// "/dev/null",
	)

	cmd.Stdin = r
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	wg.Done()

	s.mu.Lock()
	s.status.stream = true
	s.mu.Unlock()

	// Wait until stop command is received.
	<-s.channels.stopStream

	log.Println("disconnecting from Twitch's websocket server")
	exitAlerts <- struct{}{}

	s.mu.Lock()
	s.status.stream = false
	s.status.audio = false
	s.mu.Unlock()

	log.Println("stopping stream")
	cmd.Process.Signal(syscall.SIGKILL)

	r.Close()

	cmd.Wait()

	return nil
}

func (s *MainServer) initAudio(wg *sync.WaitGroup) error {
	defer log.Println("audio stopped")

	// Init song overlay
	err := s.InitOverlay()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	i := 0
	for {
		exitRoutine := make(chan struct{})

		s.currentSong = *s.playlist.Songs[0]

		// Pop the song from the queue.
		go func() {
			s.mu.Lock()
			_, s.playlist.Songs = s.playlist.Songs[0], s.playlist.Songs[1:]
			s.mu.Unlock()

			// Generate new playlist when there are ten songs left.
			if len(s.playlist.Songs) == 10 {
				// this function appends new songs to the playlist method in the MainServer struct.
				_, err := s.generateRandomPlaylist()
				if err != nil {
					log.Println(err)
				}

			}
		}()

		// Change song stream overlay
		go s.changeSongOverlay(true)

		cmd := exec.Command("ffmpeg", "-y",
			"-re",
			"-i", "files/songs/"+s.currentSong.Audio,
			"-c:a", "copy",
			"-f", "ogg", "files/stream/audio",
		)

		cmd.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGKILL,
		}

		cmd.Start()

		if i == 0 {
			wg.Done()
			i += 1
			s.mu.Lock()
			s.status.audio = true
			s.mu.Unlock()
		}

		// kill ffmpeg instance when channels.stopAudio is called. Otherwise break the for loop when exit is called and let the goroutine finish.
		go func(exitRoutine chan struct{}) {
		routine:
			for {
				select {
				case <-exitRoutine:
					break routine
				case <-s.channels.stopAudio:
					log.Println("Killing audio process and breaking out of audio loop")

					s.mu.Lock()
					s.overlays = nil
					s.status.audio = false
					s.mu.Unlock()

					cmd.Process.Signal(syscall.SIGKILL)
				}
			}

		}(exitRoutine)

		cmd.Wait()

		// Once the song ends close exitRoutine to cancel the previous go routine that listens for channels.stopAudio.
		close(exitRoutine)

		if len(s.playlist.Songs) == 0 || !s.status.audio {
			break
		}
	}

	return nil
}

func (s *MainServer) initBackgroundVideo(w *io.PipeWriter, wg *sync.WaitGroup) error {
	defer log.Println("bg video stopped")

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.QueryRow("SELECT id, name, active FROM background_videos WHERE active=1").Scan(&s.currentBackgroundVideo.Id, &s.currentBackgroundVideo.Name, &s.currentBackgroundVideo.Active)
	if err != nil {
		log.Println(err)
		return err
	}
	db.Close()

	i := 0
	var exit bool
	for {
		cmd := exec.Command("ffmpeg",
			"-hide_banner",
			"-y",
			"-re",
			"-stream_loop", "-1",
			"-i", "files/stream/background-videos/"+s.currentBackgroundVideo.Name,
			"-c:v", "copy",
			"-an",
			"-f", "h264", "-",
		)

		cmd.Stdout = w

		cmd.Start()

		if i == 0 {
			wg.Done()
			i++
		}

	inner:
		for {
			select {
			case <-s.channels.stopStream:
			case <-s.channels.stopPreview:
				exit = true
				cmd.Process.Signal(syscall.SIGKILL)
				w.Close()
				break inner
			case <-s.channels.swapBackgroundVideo:
				cmd.Process.Signal(syscall.SIGKILL)
				break inner
			}
		}

		cmd.Wait()

		if exit {
			break
		}
	}

	return nil
}

// This function is required in CreateSongPlaylist() and Audio()
func (s *MainServer) generateRandomPlaylist() (*pb.SongPlaylist, error) {

	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Retrieve available songs.
	rows, err := db.Query("SELECT page, name, author, audio_filename, cover_filename, bitrate FROM songs ORDER BY RANDOM()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		err = rows.Scan(&song.Page, &song.Name, &song.Author, &song.AudioFilename, &song.CoverFilename, &song.Bitrate)
		if err != nil {
			return nil, err
		}

		protoSong := pb.Song{
			Name:    song.Name,
			Page:    song.Page,
			Author:  song.Author,
			Audio:   song.AudioFilename,
			Cover:   song.CoverFilename,
			Bitrate: int32(song.Bitrate),
		}

		s.playlist.Songs = append(s.playlist.Songs, &protoSong)
	}

	db.Close()

	return &s.playlist, nil
}

func (s *MainServer) manageDataChannelMessage(msg []byte) {
	type Body struct {
		Type   string        `json:"type"`
		Object OverlayObject `json:"object"`
		Volume string        `json:"volume"`
	}

	var body Body

	err := json.Unmarshal(msg, &body)
	if err != nil {
		panic(err)
	}

	switch body.Type {
	case "overlay":
		s.changeOverlay(body.Object)
	case "volume":
		req := zmq4.NewReq(context.Background())
		defer req.Close()

		err := req.Dial("tcp://0.0.0.0:5554")
		if err != nil {
			log.Fatalf("could not dial: %v", err)
		}

		err = req.Send(zmq4.NewMsgString("volume@foo volume " + body.Volume))
		if err != nil {
			log.Fatalf("could not send command: %v", err)
		}

		volume, err := strconv.ParseFloat(body.Volume, 64)
		if err != nil {
			log.Fatalln(err)
		}

		s.streamParams.volume = volume
	case "volumeDb":
		volume, err := strconv.ParseFloat(body.Volume, 64)
		if err != nil {
			log.Fatalln(err)
		}
		saveVolume(volume)
	}
}

func saveVolume(volume float64) {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	_, err = db.Exec("UPDATE users SET volume = $1", volume)
	if err != nil {
		log.Println(err)
	}
}

func manageNamedPipes() error {
	// Named audio pipe
	audio := "files/stream/audio"

	// Remove the named pipe if it already exists.
	err := os.Remove(audio)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create named pipe.
	err = syscall.Mkfifo(audio, 0777)
	if err != nil {
		return err
	}

	// Named pipe
	streamOutput := "files/stream/streamOutput"

	// Remove the named pipe if it already exists.
	err = os.Remove(streamOutput)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipe.
	err = syscall.Mkfifo(streamOutput, 0777)
	if err != nil {
		panic(err)
	}

	// Named pipe
	previewAudioPipePath := "files/stream/previewAudio"

	// Remove the named pipe if it already exists.
	err = os.Remove(previewAudioPipePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipe.
	err = syscall.Mkfifo(previewAudioPipePath, 0777)
	if err != nil {
		panic(err)
	}

	previewOutput := "files/stream/previewVideo"

	// Remove the named pipe if it already exists.
	err = os.Remove(previewOutput)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipe.
	err = syscall.Mkfifo(previewOutput, 0777)
	if err != nil {
		panic(err)
	}

	streamVideo := "files/stream/streamVideo"

	// Remove the named pipe if it already exists.
	err = os.Remove(streamVideo)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipe.
	err = syscall.Mkfifo(streamVideo, 0777)
	if err != nil {
		panic(err)
	}

	return nil
}

func getStreamParams() (StreamParams, error) {
	var streamParams StreamParams

	// Get stream parameters -> preset, width, height, fps
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return StreamParams{}, err
	}

	err = db.QueryRow("SELECT preset, width, height, fps, volume FROM users WHERE id=1").Scan(&streamParams.preset, &streamParams.width, &streamParams.height, &streamParams.fps, &streamParams.volume)
	if err != nil {
		fmt.Println(err)
		return StreamParams{}, err
	}

	db.Close()

	return streamParams, nil
}

func twitchIngestLink() (string, error) {
	type ResponseBody struct {
		Ingests []struct {
			URL string `json:"url_template"`
		} `json:"ingests"`
	}

	req, err := http.NewRequest("GET", "https://ingest.twitch.tv/ingests", nil)
	if err != nil {
		return "", err
	}

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Read response body
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return "", err
	}

	resp.Body.Close()

	// Get stream key
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return "", err
	}

	var streamKey string
	err = db.QueryRow("SELECT stream_key FROM users WHERE id=1").Scan(&streamKey)
	if err != nil {
		return "", err
	}

	db.Close()

	ingestLink := strings.Replace(responseBody.Ingests[0].URL, "{stream_key}", streamKey, 1)

	return ingestLink, nil
}
