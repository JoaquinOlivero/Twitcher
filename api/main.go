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
	file, err := os.OpenFile("data.db", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		defer file.Close()

		if errors.Is(err, fs.ErrExist) {
			log.Println("data.db file already exists. Skipping creation.")
			return nil
		}

		return err
	}

	file.Close()

	// Create tables
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE songs (
			page VARCHAR NOT NULL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			release_date VARCHAR(255) NOT NULL,
			audio_filename VARCHAR(255) NOT NULL,
			cover_filename VARCHAR(255) NOT NULL,
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
			name VARCHAR(255) NOT NULL,
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
			twitch_user_id VARCHAR(255),
			twitch_user_name VARCHAR(255),
			client_id VARCHAR(255),
			secret VARCHAR(255),
			access_token VARCHAR(255),
			refresh_token VARCHAR(255),
			stream_key VARCHAR(255)
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

	return nil
}
