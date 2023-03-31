package stream

import (
	"database/sql"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	Bitrate       int
}

func Twitch(streamUrl string) error {
	// Pipe the songs to the stdin of the ffmpeg instance streaming to Twitch.
	pr, pw, err := os.Pipe()
	if err != nil {
		return err
	}

	// Infinite loop through songs and writing the output of this ffmpeg instance to the "pw" pipe.
	go func() error {

		for {

			songs, err := getSongs()
			if err != nil {
				return err
			}

			for i := 0; i < len(songs); i++ {
				// Change cover
				go changeCover(songs[i].Name, songs[i].Author, songs[i].Page, songs[i].CoverFilename)

				// Command that pipes song from the loop to "pw".
				cmd := exec.Command("ffmpeg",
					"-re", "-hide_banner",
					"-i", "files/songs/"+songs[i].AudioFilename,
					"-map", "0:a",
					"-acodec", "copy",
					"-map_metadata", "-1",
					"-fflags", "+bitexact", "-flags:v", "+bitexact", "-flags:a", "+bitexact",
					"-f", "mp3",
					"-",
				)

				cmd.Stdout = pw // Pipe the output of this first ffmpeg instance, to later be used in the second ffmpeg instance.

				cmd.Run()

				continue

			}
		}

	}()

	// Stream to Twitch's rtmp server.
	// This ffmpeg instance takes the "pr" pipe as stdin.
	cmd := exec.Command("ffmpeg",
		"-re",
		"-stream_loop", "-1",
		"-i", "files/stream/sunset-720p.mp4",
		"-f", "image2", "-loop", "1", "-i", "files/stream/stream.png", "-filter_complex", "[0:v]overlay=5:10",
		"-thread_queue_size", "4096", "-i", "pipe:0",
		"-c:v", "libx264",
		"-acodec", "aac",
		"-r", "25",
		"-g", "50",
		"-keyint_min", "50", "-force_key_frames", "expr:gte(t,n_forced*2)",
		"-use_wallclock_as_timestamps", "1",
		"-f", "flv", "-flvflags", "no_duration_filesize", streamUrl,
	)

	cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

	cmd.Stdin = pr // Stdout pipe from the looped songs ffmpeg instance. Piped into the Stdin of the second ffmpeg instance.

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

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
	time.Sleep(3 * time.Second)
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
