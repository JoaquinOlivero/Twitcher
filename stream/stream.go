package stream

import (
	"Twitcher/twitchApi"
	"bufio"
	"context"
	"database/sql"
	"errors"
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
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/google/renameio"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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

func Twitch(streamUrl string) error {
	// Pipe the songs to the stdin of the ffmpeg instance that combines audio and video.
	pr, pw, err := os.Pipe()
	if err != nil {
		return err
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
					_, err = pw.Write(buffer)
					if err != nil {
						return err
					}

				}
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
					_, err = pw.Write(buffer)
					if err != nil {
						return err
					}

				}
				file.Close()

			}

			fmt.Println("Songs loop ending. Starting a new one")
		}

	}()

	// Go routine to handle alert notifications. The notifications come from Twitch's websocket api.
	go func() error {
		var wg sync.WaitGroup

		// Go routine to validate access token and refresh it if necessary.
		// To subscribe to Twitch's websocket events it is required to have a valid user access token.
		i := 1
		wg.Add(i)
		go func() error {
			for {
				token, err := twitchApi.ValidateToken()
				if err != nil {
					log.Println(err)
					return err
				}

				// Once there is a valid access token. Decrement the waitgroup count to zero and unblock the connection to Twitch's websocket server.
				if i == 1 {
					wg.Done()
					i--
				}

				// Hold loop until 10 minutes before the access token expires and then refresh the token.
				if token.ExpiresIn > 10*60 {
					sleepDuration := time.Duration(token.ExpiresIn - (10 * 60))
					time.Sleep(sleepDuration * time.Second)
				}

				twitchApi.RefreshToken()
			}

		}()

		// wait until there is a valid acces token availabe to use.
		wg.Wait()

		// Needs to be in a for loop because the connection to the websocket server could fail and needs to be tried again.
		var wsURL string
		wsURL = "wss://eventsub.wss.twitch.tv/ws"

	websocket:
		for {
			// channels that will contain the type of notification and the username.
			channel := make(chan SubscriptionChannel, 300)
			exit := make(chan bool, 300)

			// Go routine to catch notifications coming from the websocket connection to Twitch.
			go func() {
			alert:
				for {
					select {
					case alert := <-channel:
						err := changeAlert(alert)
						if err != nil {
							log.Println(err)
						}

						time.Sleep(9 * time.Second)

						alert.Type = "reset"
						err = changeAlert(alert)
						if err != nil {
							log.Println(err)
						}

					case <-exit:
						break alert
					}
				}
			}()

			// Connect to Twitch's ws server and receive real-time notifications.
			ctx := context.Background()

			c, _, err := websocket.Dial(ctx, wsURL, nil)
			if err != nil {
				log.Println("Couldn't connect to " + wsURL + " will retry in 60 seconds.")
				log.Println("error: ", err)
				time.Sleep(60 * time.Second)

				continue
			}
			defer c.Close(websocket.StatusInternalError, "the sky is falling")

			log.Println("Connected to: ", wsURL)
			// For loop used to constantly read for websocket incoming connections.
		notifications:
			for {
				var v Response
				err = wsjson.Read(ctx, c, &v)
				if err != nil {
					if err.Error()[len(err.Error())-50:] == "WebSocket closed: failed to read frame header: EOF" { // This error happens when the websocket connection has been closed server side. For example when Twitch's websocket servers are down.
						log.Println("Couln't read from websocket. Will try to reconnect in 10 seconds.")
						log.Println("error: ", err)
						time.Sleep(10 * time.Second)
						break notifications
					}
					time.Sleep(10 * time.Second)
					log.Println(err)
					continue notifications
				}

				// When connection to wss://eventsub.wss.twitch.tv/ws is established, Twitch replies with a welcome message that contains a session_id needed to subscribe to events.
				if v.Metadata.MessageType == "session_welcome" {
					// Send POST requests to create EventSub subscriptions for channel.follow and channel.subscribe.
					err := twitchApi.CreateEventSubs(v.Payload.Session.Id, "https://api.twitch.tv/helix/eventsub/subscriptions")
					if err != nil {
						time.Sleep(10 * time.Second)
						c.Close(websocket.StatusNormalClosure, "")
						exit <- true
						log.Println(err)
						break notifications
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
					c.Close(websocket.StatusNormalClosure, "")
					continue websocket
				}

			}

			c.Close(websocket.StatusNormalClosure, "")
		}
	}()

	time.Sleep(1 * time.Second)

	// Pipe the video and audio ready to be streamed.
	spr, spw, err := os.Pipe()
	if err != nil {
		return err
	}

	// Go routine that runs ffmpeg to encode video, overlays and audio together.
	// This ffmpeg instance pipes its output to stdout to later be used in the actual ffmpeg instance that has the job to stream the input to Twitch's rtmp server.
	go func() error {

		cmd := exec.Command("ffmpeg",
			"-hide_banner",
			"-y", "-re",
			"-stream_loop", "-1",
			"-i", "files/stream/sunset-720p.mp4", // Background video
			"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", // Overlay that shows the song's cover. The "stream.png" file will be atomically changed according to the song that is being currently played.
			"-stream_loop", "-1", "-c:v", "libvpx-vp9", "-f", "concat", "-i", "files/stream/alerts/list1.txt", // Overlay that shows the alerts. Alerts are video files that use the ".webm" format.
			"-filter_complex", "[0][1]overlay=5:5[v1];[v1][2]overlay=W-w+10:H-h+60", // Filter that actually places the overlays over the video.
			"-i", "pipe:0", // Audio input pipe.
			"-c:v", "libx264", // Encode new video with overlays.
			"-c:a", "copy", // Copy the single audio stream.
			"-g", "50",
			"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
			"-f", "flv", "-flvflags", "no_duration_filesize", "-", // Pipe the result to the ffmpeg instance stdout.
		)

		// cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

		cmd.Stdin = pr   // Pipe that contains the buffered songs.
		cmd.Stdout = spw // Write the stdout the the "spw" pipe.

		err = cmd.Start()
		if err != nil {
			return err
		}

		err = cmd.Wait()
		if err != nil {
			return err
		}

		return nil
	}()

	// Ffmpeg instance that streams the piped input to Twitch's rtmp servers.
	cmd := exec.Command("ffmpeg", "-hide_banner",
		"-re", "-stream_loop", "-1",
		"-i", "pipe:0",
		"-flvflags", "no_duration_filesize",
		"-f", "fifo", "-fifo_format", "flv", // Fifo muxer implemented to recover stream in case a failure occurs.
		"-map", "0:v", "-map", "0:a",
		"-attempt_recovery", "1", "-recover_any_error", "1", "-recovery_wait_time", "1", "-flags", "+global_header", "-tag:v", "7", "-tag:a", "2",
		"-c", "copy",
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

// Change the location of the file being used as an alert in list1.txt and list2.txt files.
func changeAlert(alert SubscriptionChannel) error {

	if alert.Type == "channel.follow" {
		err := createAlert(alert.Username, "files/stream/alerts/follower-empty.webm", "files/stream/alerts/follower.webm")
		if err != nil {
			return err
		}

		err = replaceLine("files/stream/alerts/list1.txt", "follower.webm")
		if err != nil {
			return err
		}

		err = replaceLine("files/stream/alerts/list2.txt", "follower.webm")
		if err != nil {
			return err
		}

		return nil
	}

	if alert.Type == "channel.subscribe" {
		err := createAlert(alert.Username, "files/stream/alerts/sub-empty.webm", "files/stream/alerts/sub.webm")
		if err != nil {
			return err
		}
		err = replaceLine("files/stream/alerts/list1.txt", "sub.webm")
		if err != nil {
			return err
		}

		err = replaceLine("files/stream/alerts/list2.txt", "sub.webm")
		if err != nil {
			return err
		}

		return nil
	}

	// Reset the list1.txt and list2.txt files to use the default empty.webm. This empty.webm file works as a transparent placeholder for the actual alerts.
	if alert.Type == "reset" {
		err := replaceLine("files/stream/alerts/list1.txt", "empty.webm")
		if err != nil {
			return err
		}

		err = replaceLine("files/stream/alerts/list2.txt", "empty.webm")
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

// FFmpeg command that inserts text into video.
func createAlert(username, input, output string) error {
	// Add ellipsis to username if it is too long. Max 17 characters.
	maxLength := 17
	if len(username) > maxLength {
		username = username[:maxLength] + "..."
	}

	// ffmpegCommand := `-c:v libvpx-vp9 -i ` + input + ` -filter_complex "[0:v]drawtext=fontfile=Poppins-Bold.ttf:text='` + username + `':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2" ` + output + ` -y`
	// ffmpeg -c:v libvpx-vp9 -i sub.webm -filter_complex "[0:v]drawtext=fontfile=../../../Poppins-Bold.ttf:text='TestUserTwitchAAA...':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2"
	cmd := exec.Command("ffmpeg",
		"-c:v", "libvpx-vp9", "-i", input,
		"-filter_complex", "[0:v]drawtext=fontfile=Poppins-Bold.ttf:text='"+username+"':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2",
		output, "-y",
	)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func replaceLine(filePath, gif string) error {
	// Open the text file for reading and writing
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a buffer to store the lines of the file
	var lines []string

	// Read the lines from the file and store them in the buffer
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check if the file has at least two lines
	if len(lines) < 2 {
		err := errors.New("file does not have at least two lines")
		return err
	}

	// Replace the second line with a new line
	lines[1] = "file " + "'" + gif + "'"

	// Truncate the file to 0 bytes
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	// Rewind the file to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	// Write the updated lines to the file
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// Flush the writer to write the changes to disk
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
