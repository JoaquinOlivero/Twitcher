package service

import (
	"Twitcher/pb"
	"Twitcher/twitchApi"
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/google/renameio"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Song struct {
	Page          string
	Name          string
	Author        string
	AudioFilename string
	CoverFilename string
	Duration      float64
	Bitrate       int
}

type StreamManagementServer struct {
	pb.UnimplementedStreamManagementServer
	playlist pb.SongPlaylist
	mu       sync.Mutex
	audioOn  bool
	outputOn bool
	streamOn bool
}

var (
	audioDataRes = make(chan string, 300)

	sdp                 = make(chan string, 300)
	sdpForClientChannel = make(chan string, 300)

	streamChan = make(chan struct{})

	sr, sw = io.Pipe()
)

func (s *StreamManagementServer) Audio(in *pb.Empty, stream pb.StreamManagement_AudioServer) error {
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
		panic(err)
	}

	// Create named pipe.
	err = syscall.Mkfifo(audioPipePath, 0644)
	if err != nil {
		panic(err)
	}

	i := 1
	for {

		song := s.playlist.Songs[0]

		// Pop the song from the queue and send message to let the client know using webRTC.
		go func() {
			s.mu.Lock()
			_, s.playlist.Songs = s.playlist.Songs[0], s.playlist.Songs[1:]
			s.mu.Unlock()

			audioDataRes <- "pop"

			// Generate new playlist when there are ten songs left and let the client know using webRTC.
			if len(s.playlist.Songs) == 10 {
				_, err := s.generateRandomPlaylist() // this functions appends new songs to the playlist method in the StreamManagementServer struct.
				if err != nil {
					log.Println(err)
				}

				audioDataRes <- "extended"
			}
		}()

		// Change cover
		go changeCover(song.Name, song.Author, song.Page, song.Cover)

		// Open named audio pipe
		audioPipe, err := os.OpenFile("files/stream/audio", os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("ffmpeg",
			"-re",
			"-i", "files/songs/"+song.Audio,
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
			// Let client know that the audio is ready to be used in the final output.
			stream.Send(&pb.AudioStream{Ready: true})
			i--
		}

		time.Sleep(10 * time.Second)

		cmd.Process.Kill()

		cmd.Wait()

		if len(s.playlist.Songs) == 0 {
			break
		}
	}

	return nil

}

func (s *StreamManagementServer) Output(in *pb.Empty, stream pb.StreamManagement_OutputServer) error {
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
	stream.Send(&pb.OutputResponse{Ready: true})

	go Broadcast(sdp, sdpForClientChannel, audioDataRes)

	go func() {
		r := bufio.NewReader(stdout)
		buffer := make([]byte, 2*1024*1024)

		for {
			select {
			case <-streamChan:
				r.Reset(r)
				_, err := io.Copy(sw, stdout)
				if err != nil {
					panic(err)
				}

			default:
				r.Discard(r.Size())
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

	cmd.Wait()

	return nil
}

func (s *StreamManagementServer) Preview(in *pb.SDP, stream pb.StreamManagement_PreviewServer) error {

	if !s.audioOn || !s.outputOn {
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("audio or output video not available to show preview"),
		)
	}

	sdp <- in.Sdp

	sdpForClient := <-sdpForClientChannel
streamLoop:
	for {
		err := stream.Send(&pb.SDP{Sdp: sdpForClient})
		if err != nil {
			break streamLoop
		}
	}

	return nil
}

func (s *StreamManagementServer) StartTwitch(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	if s.streamOn {
		return &pb.Empty{}, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("stream has already started."),
		)
	}

	s.mu.Lock()
	s.streamOn = true
	s.mu.Unlock()

	// Enable alert notifications
	var wg sync.WaitGroup
	wg.Add(1)

	go twitchApi.Alerts(&wg)

	wg.Wait()

	cmd := exec.Command("ffmpeg", "-hide_banner",
		"-re", "-stream_loop", "-1",
		"-i", "pipe:0",
		"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
		"-thread_queue_size", "256", "-i", "files/stream/alert",
		"-filter_complex", "[0][1]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60[vout]", // Filter that actually places the overlays over the video.
		"-f", "fifo", "-fifo_format", "flv", // Fifo muxer implemented to recover stream in case a failure occurs.
		"-map", "[vout]",
		"-map", "0:a",
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header",
		"-g", "50",
		"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
		"-c:v", "libx264",
		"-acodec", "libmp3lame", "-q:a", "0",
		"-flvflags", "no_duration_filesize",
		"rtmp://bue01.contribute.live-video.net/app/live_198642898_QgYIiTqK8yCQu2sXd1jIOv79oOJBhf",
	)

	cmd.Stdin = sr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	streamChan <- struct{}{}

	cmd.Wait()

	return &pb.Empty{}, nil
}

func (s *StreamManagementServer) CreateSongPlaylist(ctx context.Context, in *pb.Empty) (*pb.SongPlaylist, error) {

	playlist, err := s.generateRandomPlaylist()
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (s *StreamManagementServer) CurrentSongPlaylist(ctx context.Context, in *pb.Empty) (*pb.SongPlaylist, error) {
	return &s.playlist, nil
}

func (s *StreamManagementServer) UpdateSongPlaylist(ctx context.Context, in *pb.SongPlaylist) (*pb.Empty, error) {

	s.mu.Lock()

	s.playlist.Songs = nil
	s.playlist.Songs = append(s.playlist.Songs, in.Songs...)

	s.mu.Unlock()

	return &pb.Empty{}, nil
}

// This function is required in CreateSongPlaylist() and Audio()
func (s *StreamManagementServer) generateRandomPlaylist() (*pb.SongPlaylist, error) {

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

func changeCover(name, author, page, cover string) error {
	// Open the original image
	file, err := os.Open("files/covers/" + cover)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Decode the image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return err
	}

	// Resize the image to a new width of 400 pixels
	newWidth := 250
	newHeight := 250
	resized := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos2)

	// Create a new image with enough space to add text to the right
	textWidth := 1280
	textHeight := 720
	withText := image.NewRGBA(image.Rect(0, 0, newWidth+textWidth, textHeight))
	draw.Draw(withText, withText.Bounds(), resized, image.Point{0, 0}, draw.Src)

	// Add song name to the right of the image
	text := name
	fontData, err := os.ReadFile("Poppins-Bold.ttf")
	if err != nil {
		fmt.Println("Error opening font file:", err)
		return err
	}
	fontSize := 36.0
	textX := newWidth + 20
	textY := int(fontSize)
	textColor := image.White
	if err := addText(withText, text, fontData, fontSize, textX, textY, textColor); err != nil {
		fmt.Println("Error adding text:", err)
		return err
	}

	// Add author
	text = author
	fontData, err = os.ReadFile("Poppins-Light.ttf")
	if err != nil {
		fmt.Println("Error opening font file:", err)
		return err
	}
	fontSize = 20.0
	textY = 36 + 30
	if err := addText(withText, text, fontData, fontSize, textX, textY, textColor); err != nil {
		fmt.Println("Error adding text:", err)
		return err
	}

	// Add ncs.io link
	text = "https://ncs.io/" + filepath.Base(page)
	fontSize = 15.0
	textX = 5
	textY = 700
	if err := addText(withText, text, fontData, fontSize, textX, textY, textColor); err != nil {
		fmt.Println("Error adding text:", err)
		return err
	}

	// Save the new image to a file
	output, err := os.Create("files/stream/next.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer output.Close()
	if err := png.Encode(output, withText); err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	time.Sleep(2 * time.Second)

	overlay, err := os.ReadFile("files/stream/next.png")
	if err != nil {
		return err
	}

	renameio.WriteFile("files/stream/stream.png", overlay, 0644)

	file.Close()
	output.Close()

	return nil
}

func addText(img *image.RGBA, text string, fontFamily []byte, fontSize float64, x, y int, textColor image.Image) error {
	// Parse font
	fontFace, err := truetype.Parse(fontFamily)
	if err != nil {
		return err
	}

	// Set the font options
	fontOptions := &truetype.Options{
		Size: fontSize,
		DPI:  72,
	}

	// Draw the text onto the image
	d := &font.Drawer{
		Dst:  img,
		Src:  textColor,
		Face: truetype.NewFace(fontFace, fontOptions),
		Dot:  fixed.P(x, y),
	}
	d.DrawString(text)

	return nil
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
