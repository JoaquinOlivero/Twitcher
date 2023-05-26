package stream

import (
	"Twitcher/twitchApi"
	"bufio"
	"database/sql"
	"encoding/json"
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
	"github.com/sacOO7/gowebsocket"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
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

// Twitch's websocket response.
type Response struct {
	Metadata struct {
		MessageId        string `json:"message_id"`
		MessageType      string `json:"message_type"`
		MessageTimestamp string `json:"message_timestamp"`
	} `json:"metadata"`

	Payload struct {
		Session struct {
			Id                      string `json:"id"`
			Status                  string `json:"status"`
			KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
			ReconnectURL            string `json:"reconnect_url"`
		} `json:"session"`

		Subscription struct {
			Id     string `json:"id"`
			Status string `json:"status"`
			Type   string `json:"type"`
		} `json:"subscription"`

		Event struct {
			UserId    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"event"`
	} `json:"payload"`
}

type SubscriptionChannel struct {
	Type     string
	Username string
}

func audio(wg *sync.WaitGroup) error {
	// Named pipe
	audioPipePath := "files/stream/audio"

	// Remove the named pipes if they already exists.
	err := os.Remove(audioPipePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipes.
	err = syscall.Mkfifo(audioPipePath, 0644)
	if err != nil {
		panic(err)
	}

	wgCountdown := 1

	for {

		// Get songs from the database in random order.
		songs, err := getSongs()
		if err != nil {
			return err
		}

		// Loop through the songs in real-time.
		for _, song := range songs {
			// Open named audio pipe
			audioPipe, err := os.OpenFile("files/stream/audio", os.O_RDWR, os.ModeNamedPipe)
			if err != nil {
				panic(err)
			}

			// Stop waiting after the silent audio file has been played.
			if wgCountdown == 1 {
				wg.Done()
				wgCountdown--
			}

			// Change cover
			go changeCover(song.Name, song.Author, song.Page, song.CoverFilename)

			// Buffer audio file in real-time to named pipe.
			file, err := os.Open("files/songs/" + song.AudioFilename)
			if err != nil {
				return err
			}
			defer file.Close()

			// Buffer with a size corresponding to the sample rate of the audio file which is 44100 Hz. All audio files have been normalize to 44100 Hz.
			r := bufio.NewReader(file)
			buffer := make([]byte, 44100*2)

			for {
				n, err := io.ReadFull(r, buffer[:cap(buffer)])
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

				// process buf
				_, err = audioPipe.Write(buffer)
				if err != nil {
					return err
				}

			}

			file.Close()
		}

		log.Println("Songs loop ending. Starting a new one")
	}
}

// func video() error {
// 	// Named pipe
// 	videoPipePath := "files/stream/video"

// 	err := os.Remove(videoPipePath)
// 	if err != nil && !os.IsNotExist(err) {
// 		panic(err)
// 	}

// 	err = syscall.Mkfifo(videoPipePath, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return nil
// }

func output(wg *sync.WaitGroup) error {
	// Named pipe
	outputPipePath := "files/stream/output"

	err := os.Remove(outputPipePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = syscall.Mkfifo(outputPipePath, 0644)
	if err != nil {
		panic(err)
	}

	wgCountdown := 1

	for {
		// Open named output pipe.
		outputPipe, err := os.OpenFile("files/stream/output", os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("ffmpeg",
			"-hide_banner",
			"-y", "-re",
			"-stream_loop", "-1",
			"-i", "files/stream/sunset-720p.mp4", // Background video
			"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
			"-f", "image2", "-loop", "1", "-i", "files/stream/alerts/frames/%d.png",
			"-filter_complex", "[0][1]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60", // Filter that actually places the overlays over the video.
			"-i", "files/stream/audio", // Audio input pipe.
			"-f", "fifo", // Fifo muxer implemented to recover stream in case of failure.
			"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header", "-tag:v", "7", "-tag:a", "2",
			"-g", "50",
			"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
			"-c:v", "libx264", // Encode new video with overlays.
			"-c:a", "copy", // Copy the single audio stream.
			// "-loglevel", "warning",
			"-f", "mpegts", "-", // Pipe the result to the ffmpeg instance stdout.
		)

		cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

		cmd.Stdout = outputPipe // Write the stdout the the "output" pipe.
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGKILL,
		}

		// Stop waiting once the ffmpeg instance starts.
		if wgCountdown == 1 {
			wg.Done()
			wgCountdown--
		}

		cmd.Run()
		// err = cmd.Start()
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return err
		// }

		// err = cmd.Wait()
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return err
		// }
	}

}

func Preview(preview bool) error {
	os.RemoveAll("files/stream/preview")
	os.MkdirAll("files/stream/preview", 0777)

	if preview {
		var (
			wgAudio  sync.WaitGroup
			wgOutput sync.WaitGroup
		)

		wgAudio.Add(1)
		wgOutput.Add(1)

		// First wait until the audio named pipe is ready so that it can later be used in the ffmpeg instance that generates the final output for the stream.
		go audio(&wgAudio)
		// wgAudio.Wait()

		// Finally wait until the ffmpeg instance that composes the whole stream is ready to be previewed.
		go output(&wgOutput)
		// wgOutput.Wait()
	}
	time.Sleep(5 * time.Second)

	// Ffmpeg instance that streams the piped input to Twitch's rtmp servers.
	cmd := exec.Command("ffmpeg", "-hide_banner",
		"-re", "-stream_loop", "1",
		"-i", "files/stream/output",
		// "-loglevel", "warning",
		"-f", "fifo", // Fifo muxer implemented to recover stream in case of failure.
		"-map", "0:v", "-map", "0:a",
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header", "-tag:v", "7", "-tag:a", "2",
		"-c", "copy",
		"-hls_time", "4",
		"-hls_list_size", "10",
		"-hls_flags", "delete_segments",
		"-f", "hls",
		"files/stream/preview/master.m3u8",
	)

	// cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Run()

	return nil
}

func Twitch(streamUrl string) error {
	// Named pipes
	audioPipePath := "files/stream/audio"
	videoPipePath := "files/stream/video"

	// Remove the named pipes if they already exists.
	err := os.Remove(audioPipePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(videoPipePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Create named pipes.
	err = syscall.Mkfifo(audioPipePath, 0777)
	if err != nil {
		panic(err)
	}

	err = syscall.Mkfifo(videoPipePath, 0777)
	if err != nil {
		panic(err)
	}

	// Infinite loop through songs and buffering the audio files in real-time to the "pw" pipe.
	go func() error {

		for {

			// Get songs from the database in random order.
			songs, err := getSongs()
			if err != nil {
				return err
			}

			// Loop through the songs in real-time.
			for _, song := range songs {
				// Open named audio pipe
				audioPipe, err := os.OpenFile("files/stream/audio", os.O_RDWR, os.ModeNamedPipe)
				if err != nil {
					panic(err)
				}

				// 2 second silent audio.
				anullsrc, err := os.Open("files/songs/silence.mp3")
				if err != nil {
					return err
				}

				defer anullsrc.Close()

				// Buffer with a size corresponding to the sample rate of the audio file which is 44100 Hz. All audio files have been normalize to 44100 Hz.
				r := bufio.NewReader(anullsrc)
				buffer := make([]byte, 44100)

				for {
					n, err := io.ReadFull(r, buffer[:cap(buffer)])
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

					// process buf
					_, err = audioPipe.Write(buffer)
					if err != nil {
						return err
					}

				}
				r.Reset(r)
				anullsrc.Close()

				// Change cover
				go changeCover(song.Name, song.Author, song.Page, song.CoverFilename)

				// Buffer audio file in real-time to pipe.
				file, err := os.Open("files/songs/" + song.AudioFilename)
				if err != nil {
					return err
				}
				defer file.Close()

				// Buffer with a size corresponding to the sample rate of the audio file which is 44100 Hz. All audio files have been normalize to 44100 Hz.
				r = bufio.NewReader(file)
				buffer = make([]byte, 44100)

				for {
					n, err := io.ReadFull(r, buffer[:cap(buffer)])
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

					// process buf
					_, err = audioPipe.Write(buffer)
					if err != nil {
						return err
					}

				}
				r.Reset(r)
				file.Close()
			}

			log.Println("Songs loop ending. Starting a new one")
		}

	}()

	// Go routine to handle alert notifications. The notifications come from Twitch's websocket api.
	var wgAlerts sync.WaitGroup
	wgAlertsI := 1
	wgAlerts.Add(wgAlertsI)

	alertFrame := 0

	go func() error {
		var wg sync.WaitGroup

		// Go routine to validate access token and refresh it if necessary.
		// To subscribe to Twitch's websocket events it is required to have a valid user access token.
		wgI := 1
		wg.Add(wgI)

		go func() {
			for {

				token, err := twitchApi.ValidateToken()
				if err != nil {
					log.Println(err)
				}

				// Once there is a valid access token. Decrement the waitgroup count to zero and unblock the connection to Twitch's websocket server.
				if wgI == 1 {
					wg.Done()
					wgI--
				}

				// Hold loop until 10 minutes before the access token expires and then refresh the token.
				if token.ExpiresIn > 10*60 {
					sleepDuration := time.Duration(token.ExpiresIn - (10 * 60))
					time.Sleep(sleepDuration * time.Second)
					log.Println("Access token will expire in ten minutes. Getting a new one.")
				}

				twitchApi.RefreshToken()
			}
		}()

		// wait until there is a valid acces token availabe to use.
		wg.Wait()

		// channels that will contain the type of notification and the username.
		channel := make(chan SubscriptionChannel, 300)
		exit := make(chan struct{}, 1)

		// Go routine to catch notifications coming from the websocket connection to Twitch.
		go func() {
		alert:
			for {
				select {
				case notification := <-channel:

					// new follower alert. Draw the username on top of the alert with a fade in effect.
					if notification.Type == "channel.follow" {
						cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/follower-empty.webm",
							"-filter_complex", "[0:v]drawtext=fontfile=../../../Poppins-Bold.ttf:text='"+notification.Username+"':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2",
							"files/stream/alerts/frames/%d.png")

						cmd.Run()
					}

					// new sub alert. Draw the username on top of the alert with a fade in effect.
					if notification.Type == "channel.subscribe" {
						cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/sub-empty.webm",
							"-filter_complex", "[0:v]drawtext=fontfile=../../../Poppins-Bold.ttf:text='"+notification.Username+"':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2",
							"files/stream/alerts/frames/%d.png")

						cmd.Run()
					}

					// Frame where the alert should stop and go back to showing a transparent png image as the alert placeholder.
					for {
						if alertFrame >= 235 {
							break
						}
					}

					// Transparent placeholder.
					cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/empty.webm", "files/stream/alerts/frames/%d.png")

					cmd.Run()

				case <-exit:
					break alert
				}
			}
		}()

		var wsURL string
		wsURL = "wss://eventsub.wss.twitch.tv/ws"
		// wsURL = "ws://127.0.0.1:8080/ws"
		reconnect := make(chan struct{}, 1)

		for {

			//Create a client instance
			socket := gowebsocket.New(wsURL)

			socket.OnConnected = func(socket gowebsocket.Socket) {
				log.Println("Connected to server: ", wsURL)
			}

			socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
				log.Println("Received connect error ", err)
			}

			// Read received messages.
			socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
				var v Response
				err = json.Unmarshal([]byte(message), &v)
				if err != nil {
					log.Println(err)
				}

				// When connection to wss://eventsub.wss.twitch.tv/ws is established, Twitch replies with a welcome message that contains a session_id needed to subscribe to events.
				if v.Metadata.MessageType == "session_welcome" {
					// Send POST requests to create EventSub subscriptions for channel.follow and channel.subscribe.
					err := twitchApi.CreateEventSubs(v.Payload.Session.Id, "https://api.twitch.tv/helix/eventsub/subscriptions")
					// err := twitchApi.CreateEventSubs(v.Payload.Session.Id, "http://127.0.0.1:8080/eventsub/subscriptions")
					if err != nil {
						log.Println(err)
					}

					if wgAlertsI == 1 {
						wgAlerts.Done()
						wgAlertsI--
					}
				}

				// Handling of notifications that come from the event channel.follow and channel.subscribe.
				// Go channels are used to handle the incoming notifications as a queue, one by one.
				if v.Metadata.MessageType == "notification" {
					var data SubscriptionChannel

					data.Type = v.Payload.Subscription.Type
					data.Username = v.Payload.Event.UserName

					channel <- data
				}

				// Twitch sends notification if the edge server the client is connected to needs to be swapped.
				if v.Metadata.MessageType == "session_reconnect" {
					wsURL = v.Payload.Session.ReconnectURL
					log.Println("Request from Twitch to reconnect websocket client to: ", wsURL)
					reconnect <- struct{}{}
				}
			}

			// This will send websocket handshake request to socketcluster-server
			socket.Connect()

			// Close socket connection and continue loop to reconnect to new server address.
			// This channel waiting also blocks the infinite loop from continuing, allowing to maintain the websocket connection and create a new one when needed.
			<-reconnect
			log.Println("reconnecting...")
			socket.Close()
			continue
		}
	}()

	wgAlerts.Wait()

	// Pipe the video and audio ready to be streamed.
	spr, spw, err := os.Pipe()
	if err != nil {
		return err
	}

	// Go routine that runs ffmpeg to encode video, overlays and audio together.
	// This ffmpeg instance pipes its output to stdout to later be used in the actual ffmpeg instance that has the job to stream the input to Twitch's rtmp server.
	go func() error {

		for {

			cmd := exec.Command("ffmpeg",
				"-hide_banner",
				"-y", "-re",
				"-stream_loop", "-1",
				"-i", "files/stream/sunset-720p.mp4", // Background video
				"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
				"-f", "image2", "-loop", "1", "-i", "files/stream/alerts/frames/%d.png",
				"-filter_complex", "[0][1]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60", // Filter that actually places the overlays over the video.
				"-i", "files/stream/audio", // Audio input pipe.
				"-c:v", "libx264", // Encode new video with overlays.
				"-c:a", "copy", // Copy the single audio stream.
				"-loglevel", "warning",
				"-f", "mpegts", "-", // Pipe the result to the ffmpeg instance stdout.
			)

			cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

			cmd.Stdout = spw // Write the stdout the the "spw" pipe.

			err = cmd.Start()
			if err != nil {
				return err
			}

			go func() {
				for {
					if alertFrame == 250 {
						alertFrame = 1
						continue
					}
					alertFrame += 1
					time.Sleep(40 * time.Millisecond)
				}
			}()

			err = cmd.Wait()
			if err != nil {
				log.Fatalln(err)
				// return err
			}
		}

	}()

	// Ffmpeg instance that streams the piped input to Twitch's rtmp servers.
	cmd := exec.Command("ffmpeg", "-hide_banner",
		"-re", "-stream_loop", "-1",
		"-i", "pipe:0",
		"-r", "25",
		"-g", "50",
		"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
		"-loglevel", "warning",
		"-f", "fifo", "-fifo_format", "flv", // Fifo muxer implemented to recover stream in case a failure occurs.
		"-map", "0:v", "-map", "0:a",
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header", "-tag:v", "7", "-tag:a", "2",
		"-c", "copy",
		"-flvflags", "no_duration_filesize",
		streamUrl,
	)

	cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.
	cmd.Stdin = spr        // Pipe that contains the video and audio ready to be streamed.

	cmd.Run()
	return nil
}

func getSongs() ([]Song, error) {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Retrieve available songs.
	var songs []Song

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

		songs = append(songs, song)
	}

	db.Close()

	return songs, nil
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
