package service

import (
	"Twitcher/pb"
	"context"
	"database/sql"
	"log"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

func (s *MainServer) YoutubeSaveStreamKey(ctx context.Context, in *pb.YoutubeStreamKey) (*google_protobuf.Empty, error) {

	// Connect to db and create user with corresponding stream key.
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE youtube_params SET stream_key=$1 WHERE user_id=1", in.Key)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) CheckYoutubeParams(ctx context.Context, in *google_protobuf.Empty) (*pb.YoutubeParams, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var streamKey, streamUrl string
	var enable bool

	err = db.QueryRow("SELECT stream_key, stream_url, enable FROM youtube_params WHERE user_id=1").Scan(&streamKey, &streamUrl, &enable)
	if err != nil {
		return nil, err
	}

	var isKeyActive bool

	if len(streamKey) > 0 {
		isKeyActive = true
	}

	return &pb.YoutubeParams{IsKeyActive: isKeyActive, Url: streamUrl, Enabled: enable}, nil
}

func (s *MainServer) DeleteYoutubeStreamKey(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE youtube_params SET stream_key=$1, enable=$2 WHERE user_id=1", "", 0)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) YoutubeSaveStreamUrl(ctx context.Context, in *pb.YoutubeStreamUrl) (*google_protobuf.Empty, error) {

	// Connect to db and create user with corresponding stream key.
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE youtube_params SET stream_url=$1 WHERE user_id=1", in.Url)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) DeleteYoutubeStreamUrl(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE youtube_params SET stream_url=$1 WHERE user_id=1", "")
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) ManageYoutube(ctx context.Context, in *pb.YoutubeParams) (*google_protobuf.Empty, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE youtube_params SET enable=$1 WHERE user_id=1", in.Enabled)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}
