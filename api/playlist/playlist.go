package playlist

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Song struct {
	Page          string
	Name          string
	Author        string
	AudioFilename string
	CoverFilename string
	Duration      float64
}

func Create() error {

	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Get songs to create new playlist.
	songs, err := GetSongs()
	if err != nil {
		return err
	}

	// Insert new row into the playlist table and get the id back.
	playlist, err := db.Exec("INSERT INTO playlists (master) VALUES ($1)", 0)
	if err != nil {
		return err
	}

	// playlist row id.
	playlistId, err := playlist.LastInsertId()
	if err != nil {
		return err
	}

	// Create playlist directory.
	err = os.MkdirAll("files/stream/playlists/"+strconv.Itoa(int(playlistId))+"/hls", 0777)
	if err != nil {
		return err
	}

	// Create new playlist in db.
	for _, song := range songs {

		_, err := db.Exec("INSERT INTO song_playlist (song_id, playlist_id) VALUES ($1,$2)", song.Page, playlistId)
		if err != nil {
			return err
		}
		// ’
		// Append song name to playlist list file.
		if !strings.Contains(song.AudioFilename, "'") {
			f, err := os.OpenFile("files/songs/"+strconv.Itoa(int(playlistId))+".txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()

			if _, err := f.WriteString("file '" + song.AudioFilename + "'\n"); err != nil {
				log.Println(err)
			}
		}

	}

	// Concatenate audio files using ffmpeg and the list.txt file.
	// ffmpeg -f concat -safe 0 -i 5.txt -c copy stream.mp3
	args := []string{"-f", "concat", "-safe", "0", "-i", strconv.Itoa(int(playlistId)) + ".txt", "-c", "copy", "../stream/playlists/" + strconv.Itoa(int(playlistId)) + "/stream.mp3"}
	var errb bytes.Buffer
	cmd := exec.Command("ffmpeg", args...)
	cmd.Dir = "files/songs"
	cmd.Stderr = &errb
	cmd.Start()

	fmt.Println(cmd.Args)
	// Do stuff...

	cmd.Wait()
	fmt.Println("ffmpeg error:", errb.String())

	// ffmpeg -i "stream.mp3" -map 0:a -profile:v baseline -level 3.0 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls hls/out.m3u8
	args = []string{"-i", "stream.mp3", "-map", "0:a", "-profile:v", "baseline", "-level", "3.0", "-start_number", "0", "-hls_time", "60", "-hls_list_size", "0", "-f", "hls", "hls/master.m3u8"}
	cmd = exec.Command("ffmpeg", args...)
	cmd.Dir = "files/stream/playlists/" + strconv.Itoa(int(playlistId))
	cmd.Stderr = &errb
	cmd.Start()

	fmt.Println(cmd.Args)
	// Do stuff...

	cmd.Wait()
	fmt.Println("ffmpeg error:", errb.String())

	return nil
}

func GetSongs() ([]Song, error) {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Retrieve available songs.
	var songs []Song

	rows, err := db.Query("SELECT page, name, author, audio_filename, cover_filename FROM songs")
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

	// Randomize the order of the songs
	rand.Seed(time.Now().UnixNano())

	// Shuffle the slice
	rand.Shuffle(len(songs), func(i, j int) {
		songs[i], songs[j] = songs[j], songs[i]
	})

	return songs, nil
}
