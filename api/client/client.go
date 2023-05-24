package main

import (
	"log"

	"Twitcher/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewSongsManagementClient(conn)

	response, err := c.CreatePlaylist(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	// log.Printf("Response from server: %s", response.GetSongs()[0:len(response.GetSongs())-1])
	log.Printf("Response from server: %s", response.GetSongs()[0])
}
