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
		// http.HandleFunc("/ws", serveWs)

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

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// const (
// 	// Time allowed to write a message to the peer.
// 	writeWait = 2 * time.Second

// 	// Time allowed to read the next pong message from the peer.
// 	pongWait = 5 * time.Second

// 	// Send pings to peer with this period. Must be less than pongWait.
// 	pingPeriod = 1 * time.Second

// 	// Maximum message size allowed from peer.
// 	maxMessageSize = 512
// )

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

// type Conn struct {
// 	// The websocket connection.
// 	ws *websocket.Conn

// 	// Buffered channel of outbound messages.
// 	send chan []byte
// }

// // readPump pumps messages from the websocket connection to the hub.
// func (c *Conn) readPump() {
// 	defer func() {

// 		c.ws.Close()
// 	}()
// 	c.ws.SetReadLimit(maxMessageSize)
// 	c.ws.SetReadDeadline(time.Now().Add(pongWait))
// 	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
// 	for {
// 		_, message, err := c.ws.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
// 				log.Printf("error: %v", err)
// 			}
// 			break
// 		}
// 		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
// 		// hub.broadcast <- message
// 		fmt.Println(message)
// 	}
// }

// // write writes a message with the given message type and payload.
// func (c *Conn) write(mt int, payload []byte) error {
// 	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
// 	return c.ws.WriteMessage(mt, payload)
// }

// // writePump pumps messages from the hub to the websocket connection.
// func (c *Conn) writePump() {
// 	ticker := time.NewTicker(pingPeriod)
// 	defer func() {
// 		ticker.Stop()
// 		c.ws.Close()
// 	}()
// 	for {
// 		select {
// 		case message, ok := <-c.send:
// 			if !ok {
// 				c.write(websocket.CloseMessage, []byte{})
// 				return
// 			}

// 			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
// 			w, err := c.ws.NextWriter(websocket.TextMessage)
// 			if err != nil {
// 				return
// 			}
// 			w.Write(message)

// 			// Add queued chat messages to the current websocket message.
// 			n := len(c.send)
// 			for i := 0; i < n; i++ {
// 				w.Write(newline)
// 				w.Write(<-c.send)
// 			}

// 			if err := w.Close(); err != nil {
// 				return
// 			}
// 		case <-ticker.C:
// 			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}
// 	}
// }

// // serveWs handles websocket requests from the peer.
// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println("connection")
// 	conn := &Conn{send: service.AudioDataRes, ws: ws}
// 	go conn.writePump()
// 	conn.readPump()
// }

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
