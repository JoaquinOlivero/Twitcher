package service

import (
	"Twitcher/pb"
	"Twitcher/stream"
	"context"
	"database/sql"
)

type SongsManagementServer struct {
	pb.UnimplementedSongsManagementServer
	// playlist *pb.Playlist
}

func (s *SongsManagementServer) CreatePlaylist(ctx context.Context, in *pb.Empty) (*pb.Playlist, error) {
	// Connect to db.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Retrieve available songs.
	var songs []*pb.Song

	rows, err := db.Query("SELECT page, name, author, audio_filename, cover_filename, bitrate FROM songs ORDER BY RANDOM()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song stream.Song
		err = rows.Scan(&song.Page, &song.Name, &song.Author, &song.AudioFilename, &song.CoverFilename, &song.Bitrate)
		if err != nil {
			return nil, err
		}

		protoSong := pb.Song{
			Name:    song.Name,
			Page:    song.Page,
			Author:  song.Author,
			Audio:   song.AudioFilename,
			Cover:   song.CoverFilename,
			Bitrate: int32(song.Bitrate),
		}

		songs = append(songs, &protoSong)
	}

	db.Close()

	return &pb.Playlist{
			Songs: songs,
		},
		nil
}
