package stream

import (
	"bytes"
	"database/sql"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang/freetype/truetype"
	"github.com/google/renameio"
	"github.com/nfnt/resize"
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
}

func Start() error {

	err := streamAudio()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// ffmpeg -re -i "Zeli - Sharks [NCS Release].mp3" -c copy -f mp3 udp://127.0.0.1:23000
func streamAudio() error {
	// ffmpeg -re -stream_loop -1 -i files/stream/stream.mp4 -f image2 -loop 1 -i files/stream/stream.png -filter_complex "overlay=40:20" -i rtsp://127.0.0.1:8554/audio  -c:v libx264 -r 33 -g 66 -keyint_min 66 -force_key_frames 'expr:gte(t,n_forced*2)' -probesize 1000 -analyzeduration 1 -movflags +faststart -reset_timestamps 1 -f flv "rtmp://bue01.contribute.live-video.net/app/live_891663522_CRTo8bzjBSLxlxGs0OU9gUDZSF5v8L"
	// ffmpeg -re -stream_loop -1 -i files/stream/stream.mp4 -f image2 -loop 1 -i files/stream/stream.png -filter_complex "overlay=40:20" -i rtsp://127.0.0.1:8554/audio  -c:v libx264 -r 33 -g 66 -keyint_min 66 -force_key_frames 'expr:gte(t,n_forced*2)' -probesize 1000 -analyzeduration 1 -movflags +faststart -reset_timestamps 1 -f rtsp -rtsp_transport tcp rtsp://127.0.0.1:8554/stream
	// Get master playlist.
	// songs, err := getPlaylistSongs()
	// if err != nil {
	// 	return err
	// }

	// // stream using ffmpeg concurrently. This way makes it possible to change the audio track and the overlay of the livestream.
	// args := []string{"-stream_loop", "-1", "-re", "-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", "-i", "files/stream/playlists/master/hls/master.m3u8", "-f", "rtsp", "-rtsp_transport", "tcp", "rtsp://127.0.0.1:8554/audio"}
	// var errb bytes.Buffer
	// cmd := exec.Command("ffmpeg", args...)
	// cmd.Stderr = &errb
	// cmd.Start()
	// // -f flv "rtmp://bue01.contribute.live-video.net/app/live_198642898_h7vzj8LGGrSS3UVIkMomDHKdWEf2VA"
	// fmt.Println(cmd.Args)
	// // -f image2 -loop 1 -i files/stream/stream.png
	// // Change cover.
	// for i := 0; i <= len(songs); i++ {
	// 	if i == len(songs) {
	// 		i = 0
	// 		continue
	// 	}
	// 	err := changeCover(songs[i].Name, songs[i].Author, songs[i].Page, songs[i].CoverFilename)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	// Wait for song to finish.
	// 	time.Sleep(time.Duration(songs[i].Duration) * time.Second)
	// }

	// cmd.Wait()
	// fmt.Println("ffmpeg error:", errb.String())
	// ffmpeg -re -stream_loop -1 -thread_queue_size 4096 -i files/stream/stream.mp4  -thread_queue_size 4096 -f image2 -loop 1 -i files/stream/stream.png -filter_complex "overlay=40:20" -thread_queue_size 4096 -i "udp://127.0.0.7:1234?timeout=50000000" -c:v libx264 -c:a copy -g 66 -keyint_min 66 -force_key_frames 'expr:gte(t,n_forced*2)' -f rtsp -rtsp_transport tcp rtsp://127.0.0.1:8554:/stream

	// chipibarijho and latest ffmpeg command tried.
	// ffmpeg -stream_loop -1 -i files/stream/stream.mp4 -f image2 -loop 1 -i files/stream/stream.png -filter_complex "overlay=40:20" -thread_queue_size 4096 -i "udp://127.0.0.7:1234?timeout=50000000&overrun_nonfatal=1&fifo_size=50000000" -c:v libx264 -c:a copy -r 33 -g 66 -keyint_min 66 -force_key_frames 'expr:gte(t,n_forced*2)' -analyzeduration 0 -probesize 32 -reset_timestamps 1 -f flv "rtmp://bue01.contribute.live-video.net/app/live_198642898_h7vzj8LGGrSS3UVIkMomDHKdWEf2VA"
	// Silence as input  -f lavfi -i anullsrc --> useful to use as a silence input to send to the main ffmpeg command

	for {

		songs, err := getSongs()
		if err != nil {
			return err
		}

		for _, song := range songs {

			// Change cover
			err := changeCover(song.Name, song.Author, song.Page, song.CoverFilename)
			if err != nil {
				return err
			}

			// Stream song.
			args := []string{"-re", "-i", "files/songs/" + song.AudioFilename, "-c:a", "copy", "-f", "mp3", "udp://127.0.0.1:1234"}
			var errb bytes.Buffer
			cmd := exec.Command("ffmpeg", args...)
			cmd.Stderr = &errb
			cmd.Start()

			fmt.Println(cmd.Args)

			// Change cover
			// time.Sleep(1 * time.Second)
			// err := changeCover(song.Name, song.Author, song.Page, song.CoverFilename)
			// if err != nil {
			// 	return err
			// }

			cmd.Wait()
			// fmt.Println("ffmpeg error:", errb.String())
		}

	}
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

	rows, err := db.Query("SELECT page, name, author, audio_filename, cover_filename FROM songs ORDER BY RANDOM()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		err = rows.Scan(&song.Page, &song.Name, &song.Author, &song.AudioFilename, &song.CoverFilename)
		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

// func getPlaylistSongs() ([]Song, error) {
// 	// Connect to db.
// 	db, err := sql.Open("sqlite3", "data.db")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer db.Close()

// 	// Get master playlist id.
// 	var playlistId int
// 	err = db.QueryRow("SELECT id FROM playlists WHERE master = 1").Scan(&playlistId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Retrieve the songs from the playlist.
// 	var songs []Song

// 	rows, err := db.Query("SELECT page, name, author, audio_filename, cover_filename, duration FROM songs INNER JOIN song_playlist ON song_playlist.song_id = songs.page INNER JOIN playlists ON playlists.id = $1;", playlistId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var song Song
// 		err = rows.Scan(&song.Page, &song.Name, &song.Author, &song.AudioFilename, &song.CoverFilename, &song.Duration)
// 		if err != nil {
// 			return nil, err
// 		}

// 		songs = append(songs, song)
// 	}

// 	return songs, nil
// }

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
	newWidth := 180
	newHeight := 180
	resized := resize.Resize(180, 180, img, resize.Lanczos2)

	// Create a new image with enough space to add text to the right
	textWidth := 1000
	textHeight := newHeight
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
	textY = 170
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
	file.Close()

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
