package service

import (
	"Twitcher/pb"
	"Twitcher/twitchApi"
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

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

type MainServer struct {
	pb.UnimplementedMainServer
	playlist    pb.SongPlaylist
	currentSong pb.Song
	overlays    []OverlayObject
	mu          sync.Mutex
	audioOn     bool
	outputOn    bool
	streamOn    bool
	findingOn   bool
}

var (
	SendChannelData    = make(chan string, 10)
	ReceiveChannelData = make(chan string, 10)

	sdp                 = make(chan string, 300)
	sdpForClientChannel = make(chan string, 300)

	stopOutputChan = make(chan struct{})

	stopAudioChan = make(chan struct{})

	streamStopChan = make(chan struct{})

	sr, sw = io.Pipe()
)

func (s *MainServer) StartAudio(ctx context.Context, in *google_protobuf.Empty) (*pb.AudioResponse, error) {
	if s.audioOn {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio already started"),
		)
	}

	if len(s.playlist.Songs) == 0 {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("cannot play audio. Playlist is empty"),
		)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go s.Audio(&wg)

	wg.Wait()

	return &pb.AudioResponse{Ready: true}, nil
}

func (s *MainServer) Audio(wg *sync.WaitGroup) error {
	if s.audioOn {
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio already started"),
		)
	}

	if len(s.playlist.Songs) == 0 {
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("cannot play audio. Playlist is empty"),
		)
	}

	s.mu.Lock()
	s.audioOn = true
	s.mu.Unlock()

	// Named pipe
	audioPipePath := "files/stream/audio"

	// Remove the named pipe if it already exists.
	err := os.Remove(audioPipePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create named pipe.
	err = syscall.Mkfifo(audioPipePath, 0644)
	if err != nil {
		return err
	}

	exitChan := make(chan struct{})

	// Init song overlay
	err = s.InitOverlay()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	i := 1
	for {
		// song := s.playlist.Songs[0]
		s.currentSong = *s.playlist.Songs[0]

		// Pop the song from the queue.
		go func() {
			s.mu.Lock()
			_, s.playlist.Songs = s.playlist.Songs[0], s.playlist.Songs[1:]
			s.mu.Unlock()

			// Generate new playlist when there are ten songs left.
			if len(s.playlist.Songs) == 10 {
				// this functions appends new songs to the playlist method in the MainServer struct.
				_, err := s.generateRandomPlaylist()
				if err != nil {
					log.Println(err)
				}

			}
		}()

		// Change song stream overlay
		// go s.changeSongOverlay(song.Name, song.Author, song.Page, song.Cover)
		go s.changeSongOverlay(true)

		// Open named audio pipe
		audioPipe, err := os.OpenFile("files/stream/audio", os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			return err
		}

		cmd := exec.Command("ffmpeg",
			"-re",
			"-i", "files/songs/"+s.currentSong.Audio,
			// "-i", "pipe:0",
			// "-c:a", "libopus", "-page_duration", "1000",
			// "-f", "ogg", "-",
			"-c:a", "copy",
			"-f", "mp3", "-",
		)

		cmd.Stdout = audioPipe
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGKILL,
		}

		cmd.Start()

		if i > 0 {
			wg.Done()
			i--
		}

		// kill ffmpeg instance when stopAudioChan is called. Otherwise break the for loop when exit is called and let the goroutine finish.
		go func(exit chan struct{}) {
		routine:
			for {
				select {
				case <-exit:
					break routine
				case <-stopAudioChan:
					log.Println("Killing audio process and breaking out of audio loop")

					s.mu.Lock()
					s.overlays = nil
					s.audioOn = false
					s.mu.Unlock()

					cmd.Process.Signal(syscall.SIGKILL)
				}
			}

		}(exitChan)

		cmd.Wait()

		// Once the song ends call exitChan to cancel the previous go routine that listens for stopAudioChan.
		exitChan <- struct{}{}

		if len(s.playlist.Songs) == 0 || !s.audioOn {
			break
		}
	}

	return nil
}

func (s *MainServer) StartOutput(ctx context.Context, in *pb.OutputRequest) (*pb.OutputResponse, error) {
	if !s.audioOn {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio is not available to start output"),
		)
	}

	if s.outputOn {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("output already started"),
		)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go s.Output(&wg, in.Mode)

	wg.Wait()

	return &pb.OutputResponse{Ready: true}, nil
}

func (s *MainServer) Output(wg *sync.WaitGroup, mode string) error {
	if !s.audioOn {
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio is not available to start output"),
		)
	}

	if s.outputOn {
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("output already started"),
		)
	}

	s.mu.Lock()
	s.outputOn = true
	s.mu.Unlock()

	exitBrChan := make(chan struct{})

	err := manageNamedPipes()
	if err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg",
		"-hide_banner",
		"-y", "-re",
		"-stream_loop", "-1",
		"-i", "files/stream/test.mp4", // Background video
		"-i", "files/stream/audio", // Audio input pipe.
		"-af", `volume@foo,azmq=bind_address=tcp\\\://0.0.0.0\\\:5554`,
		"-f", "fifo", "-fifo_format", "tee", // Fifo muxer implemented to recover stream in case of failure.
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header",
		"-map", "0:v", "-map", "1:a",
		"-c:v", "copy", // Encode new video with overlays.
		"-c:a", "libopus",
		"-b:a", "128k", "-vbr", "on", "-compression_level", "10", "-frame_duration", "60",
		"-f", "tee",
		`[select=\'a:0\':page_duration=500:f=ogg]files/stream/previewAudio
		|
		[select=\'v:0\':f=h264]files/stream/previewVideo
		|
		[f=mpegts:select=\'v:0,a\']pipe:1`,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	// webRTC connection
	go s.Broadcast(sdp, sdpForClientChannel, exitBrChan)

	// This go routine handles the stdout ("pipe:1") from the ffmpeg instance depending if it is needed for a preview in the web or a livestream.
	go func() {
		defer log.Println("closing ffmpeg stdout go routine")
		defer stdout.Close()

		switch mode {
		case "stream":
			_, err := io.Copy(sw, stdout)
			if err != nil {
				log.Println(err)
			}

		// "preview" case is used when the preview has been started but not the livestream.
		// It reads the stdout from ffmpeg to a buffer, otherwise ffmpeg will hang forever until its stdout is being read.
		case "preview":
			r := bufio.NewReader(stdout)
			buffer := make([]byte, 1*1024*1024)

		preview:
			for {
				if len(buffer) == 0 {
					r.Reset(r)
					break preview
				}

				n, err := io.ReadFull(r, buffer[:cap(buffer)-cap(buffer)/2])
				buffer = buffer[:n]

				if err != nil {
					// This is an expected EOF error because it's thrown when no more input is available and it's been made to signal a graceful end of input.
					// Basically the file has been completely read and therefore everything is OK.
					if err == io.EOF {
						break
					}

					// Unlike the previous error, an unexpected EOF means that an EOF was encountered in the middle of reading a fixed-size block or data structure.
					if err != io.ErrUnexpectedEOF {
						fmt.Fprintln(os.Stderr, err)
						break
					}

				}
			}

		}
	}()

	wg.Done()

	for v := range stopOutputChan {
		log.Println("Killing output process ", v)

		s.mu.Lock()
		s.outputOn = false
		s.streamOn = false
		s.mu.Unlock()

		exitBrChan <- struct{}{}

		break
	}

	err = cmd.Process.Signal(syscall.SIGKILL)
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Wait()

	return nil
}

func (s *MainServer) Status(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Audio: s.audioOn, Output: s.outputOn, Stream: s.streamOn}, nil
}

func (s *MainServer) StopOutput(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {
	stopOutputChan <- struct{}{}
	stopAudioChan <- struct{}{}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) Preview(ctx context.Context, in *pb.SDP) (*pb.SDP, error) {

	if !s.audioOn || !s.outputOn {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio or output video not available to show preview"),
		)
	}

	sdp <- in.Sdp

	sdpForClient := <-sdpForClientChannel

	return &pb.SDP{Sdp: sdpForClient}, nil
}

func (s *MainServer) StartStream(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {
	if s.streamOn {
		return &google_protobuf.Empty{}, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("stream has already started."),
		)
	}

	defer log.Println("Twitch stream stopped")

	s.mu.Lock()
	s.streamOn = true
	s.mu.Unlock()

	// Enable alert notifications
	var wg sync.WaitGroup
	wg.Add(1)

	exitAlerts := make(chan struct{})

	go twitchApi.Alerts(&wg, exitAlerts)

	wg.Wait()

	// Get ingest server link
	type ResponseBody struct {
		Ingests []struct {
			URL string `json:"url_template"`
		} `json:"ingests"`
	}

	req, err := http.NewRequest("GET", "https://ingest.twitch.tv/ingests", nil)
	if err != nil {
		log.Fatalln(err)
		return &google_protobuf.Empty{}, err
	}

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return &google_protobuf.Empty{}, err
	}

	// Read response body
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatalln(err)
		return &google_protobuf.Empty{}, err
	}

	resp.Body.Close()

	// Get stream key
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
		return &google_protobuf.Empty{}, err
	}

	defer db.Close()

	var streamKey string
	err = db.QueryRow("SELECT stream_key FROM users WHERE id=1").Scan(&streamKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return &google_protobuf.Empty{}, status.Errorf(
				codes.FailedPrecondition,
				fmt.Sprintln("stream key not found."),
			)
		}
		log.Fatalln(err)
		return &google_protobuf.Empty{}, err
	}

	ingestLink := strings.Replace(responseBody.Ingests[0].URL, "{stream_key}", streamKey, 1)

	cmd := exec.Command("ffmpeg", "-hide_banner",
		"-re", "-stream_loop", "-1",
		"-thread_queue_size", "256", "-i", "pipe:0",
		"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
		"-i", "files/stream/alert",
		// "-filter_complex", `[0][1]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60,zmq[vout]`,
		"-filter_complex", `[1:v]scale@cover=-1:-2[ovrl];[0:v][ovrl]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60,zmq[vout]`,
		"-f", "fifo", "-fifo_format", "flv", // Fifo muxer implemented to recover stream in case a failure occurs.
		"-map", "[vout]",
		"-map", "0:a",
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header",
		"-g", "50",
		"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
		"-c:v", "libx264",
		"-acodec", "libmp3lame", "-q:a", "0",
		"-flvflags", "no_duration_filesize",
		ingestLink,
	)

	cmd.Stdin = sr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	// Wait until stop command is received.
	<-streamStopChan

	log.Println("disconnecting from Twitch's websocket server")
	exitAlerts <- struct{}{}

	log.Println("stopping stream")
	err = cmd.Process.Signal(syscall.SIGKILL)
	if err != nil {
		log.Fatalln(err)
	}

	err = sr.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = sw.Close()
	if err != nil {
		log.Fatalln(err)
	}

	s.mu.Lock()
	s.outputOn = false
	s.streamOn = false
	s.mu.Unlock()

	cmd.Wait()

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) StopStream(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {
	streamStopChan <- struct{}{}
	return &google_protobuf.Empty{}, nil
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
	}

	var body Body

	err := json.Unmarshal(msg, &body)
	if err != nil {
		panic(err)
	}

	switch body.Type {
	case "overlay":
		s.changeOverlay(body.Object)
	}
}

func manageNamedPipes() error {
	// Named pipe
	streamOutput := "files/stream/streamOutput"

	// Remove the named pipe if it already exists.
	err := os.Remove(streamOutput)
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

	return nil
}
