package main

import (
	"Twitcher/playlist"
	"Twitcher/stream"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"database/sql"

	"github.com/gocolly/colly/v2"
	"github.com/hajimehoshi/go-mp3"
	_ "github.com/mattn/go-sqlite3"
)

type SongDetails struct {
	Name         string
	Author       string
	Genre        string
	Page         string
	ReleaseDate  string
	DownloadLink string
	CoverSrc     string
}

type Mood struct {
	Id    int
	NcsId int
	Name  string
}

type AudioProbe struct {
	Format struct {
		Duration string `json:"duration"`
	}
}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		err := errors.New("this program needs at least one argument to run")
		log.Fatalln(err)
	}

	if len(flag.Args()) > 1 {
		err := errors.New("can't use more than one argument")
		log.Fatalln(err)
	}

	if flag.Arg(0) == "search" {

		moods, err := getMoods()
		if err != nil {
			log.Fatalln(err)
		}

		for _, mood := range moods {
			err := crawlByMood(mood.NcsId)
			if err != nil {
				fmt.Println(err)
			}
		}

		return
	}

	if flag.Arg(0) == "stream" {
		err := stream.Start()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if flag.Arg(0) == "playlist" {
		err := playlist.Create()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getMoods() ([]Mood, error) {
	// Query db to retrieve moods.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var moods []Mood

	sqlQuery := "SELECT * FROM moods"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mood Mood
		err = rows.Scan(&mood.Id, &mood.NcsId, &mood.Name)
		if err != nil {
			return nil, err
		}

		moods = append(moods, mood)
	}

	return moods, nil
}

func crawlByMood(moodId int) error {
	c := colly.NewCollector(
		colly.AllowedDomains("ncs.io", "ncs.lnk.to"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("table.tablesorter > tbody tr", func(e *colly.HTMLElement) {
		var song SongDetails
		var ok bool

		// Check for unwanted tags. If unwanted tags are present, ok will be set to false and the below foreach loop will break. Also the rest of the crawling will be skipped.
		// Also, there is a call to the function "noSong" that returns false if the song is already in the database. If the song is already in the db, ok will be set to false and then break the foreach loop. This will also skip the rest of the crawling because the boolean "ok" will be false.
		e.ForEachWithBreak("td a", func(_ int, elem *colly.HTMLElement) bool {
			if elem.Attr("class") == "tag" {
				ok = checkTags(moodId, elem.Text)

				if !ok {
					return false
				}
			}

			// There is an "a" html element that contains the href to the song's individual page. However this "a" html element doesn't have a class, while the other "a" html element does and they should be looped separately to get the needed href.
			if elem.Index == 1 && elem.Attr("class") != "tag" {
				song.Page = "https://ncs.lnk.to" + elem.Attr("href")
				ok = songDoesNotExist(song.Page)

				if !ok {
					return false
				}
			}

			return true
		})

		// if ok is true it means that the song being scraped does not contain any unwanted tags AND is not in the database.
		if ok {

			song.Genre = strings.Split(e.ChildAttr("td span.genre", "title"), ",")[0]
			song.Name = e.ChildText("td a p")
			song.Author = e.ChildText("td span")
			song.DownloadLink = "https://ncs.io/track/download/" + e.ChildAttr("td a", "data-tid")

			selection := e.DOM
			childNodes := selection.Children().Nodes

			song.ReleaseDate = childNodes[5].FirstChild.Data

			song.CoverSrc = crawlSongPage(song.Page, c)

			// songs = append(songs, song)
			err := saveSong(song)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	c.OnHTML("a.page-link", func(e *colly.HTMLElement) {
		next_page := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.Visit("https://ncs.io/music-search?q=&genre=&mood=" + strconv.Itoa(moodId))

	return nil
}

func checkTags(moodId int, tag string) bool {
	switch moodId {
	case 13: // Laid Back
		switch tag {
		case "Dark", "Dubstep", "Future House":
			return false
		}

	case 17: // Relaxing
		switch tag {
		case "Drum & Bass", "Dubstep", "Trap", "Future House":
			return false
		}
	case 19: // Romantic
		switch tag {
		case "Drum & Bass", "Dubstep", "Future House":
			return false
		}
	}

	return true
}

func crawlSongPage(url string, c *colly.Collector) string {
	var imageUrl string

	c.OnHTML("div.site-bg > img", func(e *colly.HTMLElement) {
		imageUrl = e.Attr("src")
	})

	c.Visit(url)

	return imageUrl
}

// Todo: handle errors in this function.
func songDoesNotExist(page string) bool {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return false
	}

	defer db.Close()

	// Query at least one row to check if it returns a value or an error. If it returns sql.ErrNoRows it means that the song is not in the database and therefore should return true.
	var name string

	sqlQuery := "SELECT name FROM songs WHERE page = ?"
	err = db.QueryRow(sqlQuery, page).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
	}

	return false
}

func saveSong(song SongDetails) error {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}

	defer db.Close()

	// Download audio file.
	audioFilename, err := downloadFile(song.DownloadLink, "files/songs")
	if err != nil {
		return err
	}

	// Download cover image.
	coverFilename, err := downloadFile(song.CoverSrc, "files/covers")
	if err != nil {
		return err
	}

	// Get song duration.
	duration, err := getAudioLength(audioFilename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("INSERT INTO songs (page, name, genre, author, release_date, audio_filename, cover_filename, duration) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", song.Page, song.Name, song.Genre, song.Author, song.ReleaseDate, audioFilename, coverFilename, duration)
	if err != nil {
		return err
	}

	return nil
}

func downloadFile(url, saveDir string) (string, error) {

	// Make HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Get the filename from the Content-Disposition header if it exists.
	var filename string

	disposition := resp.Header["Content-Disposition"]
	if len(disposition) > 0 {
		_, params, err := mime.ParseMediaType(disposition[0])
		if err != nil {
			return "", err
		}
		// For some reason some audio files from ncs.io don't have an extension.
		if filepath.Ext(params["filename"]) != ".mp3" {
			filename = params["filename"] + ".mp3"

		} else {
			filename = params["filename"]
		}
	}

	// If the filename is still empty, extract it from the URL
	if filename == "" {
		filename = strings.Split(url, "/")[5] + filepath.Ext(url)
	}

	// Create the file
	out, err := os.Create(saveDir + "/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	fmt.Println("Downloading: " + filename)
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Downloaded: " + filename)

	return filename, nil
}

func getAudioLength(filename string) (int64, error) {
	f, err := os.Open("files/songs/" + filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return 0, err
	}
	f.Close()

	const sampleSize = 4                           // From documentation.
	samples := d.Length() / sampleSize             // Number of samples.
	audioLength := samples / int64(d.SampleRate()) // Audio length in seconds.

	return audioLength, nil
}

// Genres --> skip
// 11 Indie
// 10 House
// 12 Melodic Dubstep
// 2 Chill --> skip if Dark

// Moods --> skip
// 13 Laid Back ---> skip if Dark | Dubstep  tag
// 17 Relaxing ---> skip if Drum & Bass | Dubstep | Trap | Future House
// 19 Romantic
// INSERT INTO moods (ncsid, name) VALUES(19, 'romantic');
