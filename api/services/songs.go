package service

import (
	"Twitcher/pb"
	"bufio"
	"context"
	"database/sql"
	"errors"
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
	"sync"
	"syscall"
	"time"

	"github.com/gocolly/colly"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SongDetails struct {
	Name         string
	Author       string
	Genre        string
	Page         string
	AltPage      string
	ReleaseDate  string
	DownloadLink string
	CoverSrc     string
}

type Mood struct {
	Id    int
	NcsId int
	Name  string
}

var songs []SongDetails

func (s *MainServer) FindNewSongsNCS(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {

	if s.status.finding {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("Finding new songs. Please wait..."),
		)
	}

	defer log.Println("Finished finding new songs")

	s.mu.Lock()
	s.status.finding = true
	s.mu.Unlock()

	moods, err := getMoods()
	if err != nil {
		log.Fatalln(err)
	}

	for _, mood := range moods {
		err := crawlByMood(mood.NcsId)
		if err != nil {
			log.Println(err)
		}
	}

	maxDownloads := 4
	guard := make(chan struct{}, maxDownloads)
	var wg sync.WaitGroup

	for _, song := range songs {
		guard <- struct{}{}
		wg.Add(1)
		go func(song SongDetails) {

			c := colly.NewCollector(
				colly.AllowedDomains("ncs.io", "ncs.lnk.to"),
				colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
			)

			songURL := getSongURL(song.Page, c)
			song.DownloadLink = songURL

			coverURL := getCoverURL(song.AltPage, c)
			song.CoverSrc = coverURL

			if song.CoverSrc != "" && song.DownloadLink != "" {
				err := saveSong(song)
				if err != nil {
					log.Println(err)
					<-guard
					wg.Done()
					return
				}

				tempFilename := song.Author + " - " + song.Name + ".mp3"
				filename := song.Author + " - " + song.Name + ".opus"

				cmd := exec.Command("ffmpeg", "-hide_banner", "-y",
					"-i", "files/songs/"+tempFilename,
					"-c:a", "libopus", "-b:a", "128k", "-vbr", "on", "-compression_level", "10", "-frame_duration", "60", "-application", "audio", "-page_duration", "500",
					"files/songs/"+filename,
				)

				cmd.SysProcAttr = &syscall.SysProcAttr{
					Pdeathsig: syscall.SIGKILL,
				}

				err = cmd.Run()
				if err != nil {
					// Remove from DB if error.
					db, err := sql.Open("sqlite3", "data.db")
					if err != nil {
						log.Println(err)
					}

					defer db.Close()

					_, err = db.Exec("DELETE FROM songs WHERE page=$1", song.Page)
					if err != nil {
						log.Fatalln(err)
					}
					cmd.Process.Signal(syscall.SIGKILL)
					cmd.Wait()
				}

				os.Remove("files/songs/" + tempFilename)

			} else {
				log.Println("Skipping: ", song.Name)
			}

			<-guard
			wg.Done()

		}(song)
	}

	// wait for all songs to be processed.
	wg.Wait()

	s.mu.Lock()
	s.status.finding = false
	s.mu.Unlock()

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) StatusNCS(ctx context.Context, in *google_protobuf.Empty) (*pb.StatusNCSResponse, error) {
	return &pb.StatusNCSResponse{Active: s.status.finding}, nil
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
		colly.Async(true),
		colly.AllowedDomains("ncs.io", "ncs.lnk.to"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 8,
		RandomDelay: 5 * time.Second,
	})

	// On every a element which has href attribute call callback
	c.OnHTML("table.tablesorter > tbody tr", func(e *colly.HTMLElement) {
		var song SongDetails
		var ok bool

		// Check for unwanted tags. If unwanted tags are present, ok will be set to false and the below foreach loop will break. Also the rest of the crawling will be skipped.
		// Also, there is a call to the function "noSong" that returns false if the song is already in the database or it has already been found previously crawling for a different mood (this can happen because songs can have multiple moods). If the song is already in the db, ok will be set to false and then break the foreach loop. This will also skip the rest of the crawling because the boolean "ok" will be false.
		e.ForEachWithBreak("td a", func(_ int, elem *colly.HTMLElement) bool {
			if elem.Attr("class") == "tag" {
				ok = checkTags(moodId, elem.Text)

				if !ok {
					return false
				}
			}

			// There is an "a" html element that contains the href to the song's individual page. However this "a" html element doesn't have a class, while the other "a" html element does and they should be looped separately to get the needed href.
			if elem.Index == 1 && elem.Attr("class") != "tag" {
				song.Page = "https://ncs.io" + elem.Attr("href")
				song.AltPage = "https://ncs.lnk.to" + elem.Attr("href")
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

			songs = append(songs, song)
		}
	})

	c.OnHTML("a.page-link", func(e *colly.HTMLElement) {
		next_page := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	c.Visit("https://ncs.io/music-search?q=&genre=&mood=" + strconv.Itoa(moodId))

	c.Wait()

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

func getSongURL(url string, c *colly.Collector) (downloadLink string) {
	var downloadURL string

	c.OnHTML("div.waveform", func(e *colly.HTMLElement) {
		if e.Attr("id") == "player" {
			id := e.Attr("data-tid")
			downloadURL = "https://ncs.io/track/download/" + id
		}
	})

	c.Visit(url)
	return downloadURL
}

func getCoverURL(url string, c *colly.Collector) (coverSrc string) {
	var cover string

	c.OnHTML("div.site-bg > img", func(e *colly.HTMLElement) {
		cover = e.Attr("src")
	})

	c.Visit(url)
	return cover
}

// Todo: handle errors in this function.
func songDoesNotExist(page string) bool {
	// First check if the song has been found on a previous crawl.
	for i := 0; i < len(songs); i++ {
		if songs[i].Page == page {
			return false
		}
	}

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
	audioFilename, err := downloadSong(song.DownloadLink, song.Name, song.Author)
	if err != nil {
		return err
	}

	// Check that the bitrate is below 400kbps
	err = audioBitrate("files/songs/" + song.Author + " - " + song.Name + ".mp3")
	if err != nil {
		return err
	}

	// Download cover image.
	coverFilename, err := downloadFile(song.CoverSrc, "files/covers")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO songs (page, name, genre, author, release_date, audio_filename, cover_filename, bitrate) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", song.Page, song.Name, song.Genre, song.Author, song.ReleaseDate, audioFilename, coverFilename, 128)
	if err != nil {
		return err
	}

	return nil
}

func downloadSong(url, name, author string) (string, error) {
	tempFilename := author + " - " + name + ".mp3"

	log.Println("Downloading: ", tempFilename)
	out, err := os.Create("files/songs/" + tempFilename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Downloaded: ", tempFilename)

	filename := author + " - " + name + ".opus"

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

	log.Println("Downloading: " + filename)
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Downloaded: " + filename)

	return filename, nil
}

func audioBitrate(file string) error {
	var bitrate int

	// Get sample rate.
	cmd := exec.Command("ffprobe", "-hide_banner", "-select_streams", "a", "-show_streams", file)

	stdOut, _ := cmd.StdoutPipe()

	cmd.Start()

	// Scan line by line stderr for ffprobe result. Get line that contains "sample_rate" and " bit_rate"
	scanner := bufio.NewScanner(stdOut)
	for scanner.Scan() {
		m := scanner.Text()

		// Get bitrate
		if strings.HasPrefix(m, "bit_rate") {
			_, bitrateLine, _ := strings.Cut(m, "bit_rate=")
			bitrateInt, err := strconv.Atoi(bitrateLine)
			if err != nil {
				return err
			}

			bitrate = bitrateInt
		}

	}

	cmd.Wait()

	if bitrate > 400000 {
		err := errors.New("bitrate is higher than 400kbps")
		return err
	}

	return nil
}
