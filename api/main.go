package main

import (
	"Twitcher/pb"
	service "Twitcher/services"
	"errors"
	"net"

	"flag"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()

	if len(flag.Args()) > 1 {
		err := errors.New("can't use more than one argument")
		log.Fatalln(err)
	}

	if flag.Arg(0) == "gui" {

		// serve static files
		go func() {
			http.Handle("/covers/", http.StripPrefix("/covers/", http.FileServer(http.Dir("files/covers"))))
			http.Handle("/preview/", http.StripPrefix("/preview/", http.FileServer(http.Dir("files/stream/preview"))))

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

	// Save client id and client secret to database. Ask user for input.
	// if flag.Arg(0) == "api" {
	// // client id
	// // nmt0hburg9hhvc7txcv3qlofl8lp01
	// // client secret
	// // zpl1ancbp3qbsq43g7jj9cix1zj6iz

	// // https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=nmt0hburg9hhvc7txcv3qlofl8lp01&redirect_uri=http://localhost:3000/twitch&scope=moderator%3Aread%3Afollowers+channel%3Aread%3Asubscriptions+user%3Aread%3Aemail

	// scanner := bufio.NewScanner(os.Stdin)

	// // Get client id from user input.
	// fmt.Print("Enter your client id: ")
	// scanner.Scan()
	// clientId := scanner.Text()

	// // Get secret from user input.
	// fmt.Print("Enter your secret: ")
	// scanner.Scan()
	// secret := scanner.Text()

	// fmt.Println("Go to: https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=" + clientId + "&redirect_uri=http://localhost:3000/twitch&scope=moderator%3Aread%3Afollowers+channel%3Aread%3Asubscriptions+user%3Aread%3Aemail")

	// // Get code from user input.
	// fmt.Printf("Enter your code: ")
	// scanner.Scan()
	// code := scanner.Text()

	// err := twitchApi.SaveClient(clientId, secret, code)
	// if err != nil {
	// 	log.Println(err)
	// 	panic(err)
	// }
	// }

}
