package main

import (
	"Twitcher/stream"
	"Twitcher/twitchApi"
	// "Twitcher/twitchApi"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"database/sql"

	"github.com/gocolly/colly/v2"
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
	Bitrate    int
	SampleRate int
}

func main() {

	twitchPtr := flag.String("twitch", "", "a string")

	flag.Parse()

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

	// Save client id and client secret to database. Ask user for input.
	if flag.Arg(0) == "api" {
		scanner := bufio.NewScanner(os.Stdin)

		// Get Twitch's username
		fmt.Print("Enter your Twitch's username: ")
		scanner.Scan()
		username := scanner.Text()

		// Get client id from user input.
		fmt.Print("Enter your client id: ")
		scanner.Scan()
		clientId := scanner.Text()

		// Get secret from user input.
		fmt.Print("Enter your secret: ")
		scanner.Scan()
		secret := scanner.Text()

		fmt.Println("Go to: https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=" + clientId + "&redirect_uri=http://localhost:9696&scope=moderator%3Aread%3Afollowers+channel%3Aread%3Asubscriptions+user%3Aread%3Aemail")

		// Get code from user input.
		fmt.Printf("Enter your code: ")
		scanner.Scan()
		code := scanner.Text()

		err := twitchApi.SaveClient(username, clientId, secret, code)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	// Twitch stream
	if *twitchPtr != "" {
		err := stream.Twitch(*twitchPtr)
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

			selection := e.DOM
			childNodes := selection.Children().Nodes

			song.ReleaseDate = childNodes[5].FirstChild.Data

			// Crawl for the download link and cover image on the song's individual page.
			coverSrc, downloadLink := crawlSongPage(song.Page, c)

			song.CoverSrc = coverSrc
			if downloadLink != "" {
				song.DownloadLink = downloadLink
			} else {
				song.DownloadLink = "https://ncs.io/track/download/" + e.ChildAttr("td a", "data-tid")
			}

			if song.CoverSrc != "" && song.DownloadLink != "" {
				err := saveSong(song)
				if err != nil {
					fmt.Println(err)
				}
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
	case 15: // Peaceful
		switch tag {
		case "Drum & Bass", "Dubstep", "Future House", "Melodic Dubstep", "Drumstep":
			return false
		}
	case 3, 4: // Dreamy and Epic
		switch tag {
		case "Drum & Bass", "Dubstep", "Future House", "Melodic Dubstep", "Drumstep", "Future Bass", "Bass House", "Glitch Hop", "Hardstyle", "Weird", "Bass", "Trap", "Sexy", "EDM", "House", "N/A", "Phonk":
			return false
		}
	}

	return true
}

func crawlSongPage(url string, c *colly.Collector) (coverSrc, downloadLink string) {
	var cover, download string

	c.OnHTML("div.site-bg > img", func(e *colly.HTMLElement) {
		cover = e.Attr("src")
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		if e.Attr("data-label") == "freedownloads" {
			download = e.Attr("href")
		}
	})

	c.Visit(url)

	return cover, download
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

	// Get audio bitrate and sample rate.
	audioFileData, err := audioData(song.DownloadLink)
	if err != nil {
		return err
	}

	// Download audio file.
	audioFilename, err := downloadSong(song.DownloadLink, song.Name, song.Author, audioFileData.Bitrate, audioFileData.SampleRate)
	if err != nil {
		return err
	}

	// Download cover image.
	coverFilename, err := downloadFile(song.CoverSrc, "files/covers")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO songs (page, name, genre, author, release_date, audio_filename, cover_filename, bitrate) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", song.Page, song.Name, song.Genre, song.Author, song.ReleaseDate, audioFilename, coverFilename, audioFileData.Bitrate)
	if err != nil {
		return err
	}

	return nil
}

func downloadSong(url, name, author string, bitrate, sampleRate int) (string, error) {
	// Download song using ffmpeg. This download method will check that the audio file is not corrupted. Also it'll change the sample rate to 44.1 kHz if needed.
	if sampleRate == 44100 {
		cmd := exec.Command("ffmpeg", "-hide_banner", "-i", url, "-map", "a:0", "-b:a", strconv.Itoa(bitrate)+"k", "files/songs/"+author+" - "+name+".mp3", "-y")
		cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

		err := cmd.Run()
		if err != nil {
			return "", err
		}

	} else {
		cmd := exec.Command("ffmpeg", "-hide_banner", "-i", url, "-map", "a:0", "-b:a", strconv.Itoa(bitrate)+"k", "-af", "aresample=resampler=soxr", "-ar", "44100", "files/songs/"+author+" - "+name+".mp3", "-y")
		cmd.Stderr = os.Stderr // ffmpeg logs everything to stderr.

		err := cmd.Run()
		if err != nil {
			return "", err
		}

	}

	filename := author + " - " + name + ".mp3"
	return filename, nil
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

		filename = params["filename"]
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

func audioData(file string) (AudioProbe, error) {
	var audioData AudioProbe

	// Get sampel rate.
	cmd := exec.Command("ffprobe", "-hide_banner", "-select_streams", "a", "-show_streams", file)

	stdOut, _ := cmd.StdoutPipe()

	cmd.Start()

	// Scan line by line stderr for ffprobe result. Get line that contains "sample_rate" and " bit_rate"
	scanner := bufio.NewScanner(stdOut)
	for scanner.Scan() {
		m := scanner.Text()

		// Get sample rate
		if strings.HasPrefix(m, "sample_rate") {
			_, sampleRateLine, _ := strings.Cut(m, "sample_rate=")
			sampleRateInt, err := strconv.Atoi(sampleRateLine)
			if err != nil {
				return audioData, err
			}

			audioData.SampleRate = sampleRateInt
		}

		// Get bitrate
		if strings.HasPrefix(m, "bit_rate") {
			_, bitrateLine, _ := strings.Cut(m, "bit_rate=")
			bitrateInt, err := strconv.Atoi(bitrateLine)
			if err != nil {
				return audioData, err
			}

			audioData.Bitrate = bitrateInt
		}

	}

	if audioData.Bitrate > 400000 {
		err := errors.New("bitrate is higher than 400kbps")
		return audioData, err
	}

	cmd.Wait()

	return audioData, nil
}
