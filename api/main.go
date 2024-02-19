package main

import (
	"Twitcher/pb"
	service "Twitcher/services"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"

	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {

	// check database.
	err := checkDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// serve static files
	go func() {
		http.Handle("/covers/", http.StripPrefix("/covers/", http.FileServer(http.Dir("files/covers"))))

		log.Println("Static file server listening at :9001")
		if err := http.ListenAndServe(":9001", nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register service methods
	pb.RegisterMainServer(grpcServer, &service.MainServer{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port 9000: %v", err)
	}

}

func checkDatabase() error {
	log.Println("Checking database")

	// Check if database already exists.
	file, err := os.OpenFile("files/data.db", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0744)
	if err != nil {
		defer file.Close()

		if errors.Is(err, fs.ErrExist) {
			log.Println("data.db file already exists. Skipping creation.")
			return nil
		}

		return err
	}

	file.Close()
	// First time starting the program.
	// Create necessary directories.
	err = os.MkdirAll("files/covers", 0744)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll("files/songs", 0744)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll("files/stream/background-videos", 0744)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll("files/stream/alerts", 0744)
	if err != nil {
		log.Fatalln(err)
	}

	// Create tables
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE songs (
			page TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			genre TEXT NOT NULL,
			author TEXT NOT NULL,
			release_date TEXT NOT NULL,
			audio_filename TEXT NOT NULL,
			cover_filename TEXT NOT NULL,
			bitrate INT NOT NULL,
			UNIQUE(page)
		)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE moods (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			ncsid INTEGER NOT NULL,
			name TEXT NOT NULL,
			UNIQUE(name),
			UNIQUE(ncsid)
		)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			preset TEXT CHECK (preset IN ("ultrafast", "superfast", "veryfast", "faster","fast", "medium", "slow", "slower", "veryslow")) DEFAULT "medium",
			width INTEGER DEFAULT 1280,
    		height INTEGER DEFAULT 720,
    		fps INTEGER DEFAULT 25,
			volume REAL DEFAULT 0.50
		)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE background_videos (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			active BOOLEAN NOT NULL CHECK (active IN (0, 1))
		)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE youtube_params (
			user_id INTEGER,
			enable INTEGER NOT NULL DEFAULT 0,
			stream_key VARCHAR(255) NOT NULL DEFAULT "",
			stream_url VARCHAR(255) NOT NULL DEFAULT "",
			CONSTRAINT FK_user FOREIGN KEY (user_id) REFERENCES users (id)
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE twitch_params (
			user_id INTEGER,
			enable INTEGER NOT NULL DEFAULT 0,
			twitch_user_id VARCHAR(255),
			twitch_user_name VARCHAR(255),
			client_id VARCHAR(255),
			secret VARCHAR(255),
			access_token VARCHAR(255),
			refresh_token VARCHAR(255),
			stream_key VARCHAR(255),
			CONSTRAINT FK_user FOREIGN KEY (user_id) REFERENCES users (id)
		)
	`)
	if err != nil {
		return err
	}

	// Insert default user data.
	_, err = db.Exec(`
		INSERT INTO users (id) VALUES (1)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO youtube_params (user_id,enable) VALUES (1,0)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO twitch_params (user_id,enable) VALUES (1,0)
	`)

	if err != nil {
		return err
	}

	// Insert moods data
	_, err = db.Exec(`
		INSERT INTO moods (ncsid, name) VALUES (13, 'laid back')
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO moods (ncsid, name) VALUES (17, 'relaxing')
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO moods (ncsid, name) VALUES (19, 'romantic')
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO moods (ncsid, name) VALUES (15, 'peaceful')
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO moods (ncsid, name) VALUES (4, 'epic')
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE overlays (
			id TEXTNULL PRIMARY KEY,
			type TEXT NOT NULL CHECK (type IN ("img", "textbox")),
			width INTEGER NOT NULL,
			height INTEGER NOT NULL,
			point_x INTEGER NOT NULL,
			point_y INTEGER NOT NULL,
			show BOOLEAN NOT NULL CHECK (show IN (0, 1)),
			font_family TEXT,
			font_size INTEGER,
			line_height REAL,
			text_color TEXT,
			text_align TEXT CHECK (text_align IN ("left", "center", "right")),
			UNIQUE(id)
		)
	`)

	if err != nil {
		return err
	}

	// insert default overlays into the db.
	_, err = db.Exec(`
		INSERT INTO overlays (id, type, width, height, point_x, point_y, show) VALUES ("cover", "img", 250, 250, 5, 5, 1)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO overlays (id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align) VALUES ("song_name", "textbox", 1000, 0, 275, 5, 1, "Poppins-Bold.ttf", 36, 1.16, "255 255 255", "left")
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO overlays (id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align) VALUES ("song_author", "textbox", 1000, 0, 275, 50, 1, "Poppins-Light.ttf", 20, 1.16, "255 255 255", "left")
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO overlays (id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align) VALUES ("song_page", "textbox", 1000, 0, 5, 700, 1, "Poppins-Light.ttf", 15, 1.16, "255 255 255", "left")
	`)
	if err != nil {
		return err
	}

	return nil
}
