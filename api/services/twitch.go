package service

import (
	"Twitcher/pb"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

func (s *MainServer) TwitchSaveStreamKey(ctx context.Context, in *pb.TwitchStreamKey) (*google_protobuf.Empty, error) {

	// Connect to db and create user with corresponding stream key.
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE twitch_params SET stream_key=$1 WHERE user_id=1", in.Key)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) CheckTwitchStreamKey(ctx context.Context, in *google_protobuf.Empty) (*pb.TwitchStreamKey, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var streamKey string

	err = db.QueryRow("SELECT stream_key FROM twitch_params WHERE user_id=1").Scan(&streamKey)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return &pb.TwitchStreamKey{Active: false}, nil
		}
		return nil, err
	}

	return &pb.TwitchStreamKey{Active: true}, nil

}

func (s *MainServer) DeleteTwitchStreamKey(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE twitch_params SET stream_key=$1 WHERE user_id=1", nil)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) SaveTwitchDevCredentials(ctx context.Context, in *pb.DevCredentials) (*google_protobuf.Empty, error) {
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return &google_protobuf.Empty{}, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE twitch_params SET client_id=$1, secret=$2 WHERE user_id=1", in.ClientId, in.Secret)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) CheckTwitchDevCredentials(ctx context.Context, in *google_protobuf.Empty) (*pb.DevCredentials, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return &pb.DevCredentials{}, err
	}

	defer db.Close()

	var clientId string

	err = db.QueryRow("SELECT client_id FROM twitch_params WHERE user_id=1").Scan(&clientId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return &pb.DevCredentials{Active: false}, nil
		}
		return nil, err
	}

	return &pb.DevCredentials{ClientId: clientId, Active: true}, nil
}

func (s *MainServer) DeleteTwitchDevCredentials(ctx context.Context, in *google_protobuf.Empty) (*google_protobuf.Empty, error) {

	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE twitch_params SET client_id=$1, secret=$1, access_token=$1, refresh_token=$1 WHERE user_id=1", nil)
	if err != nil {
		log.Println(err)
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

func (s *MainServer) TwitchAccessToken(ctx context.Context, in *pb.UserAuth) (*google_protobuf.Empty, error) {
	// Get client id and secret from db
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return &google_protobuf.Empty{}, err
	}

	defer db.Close()

	var clientId, secret string

	err = db.QueryRow("SELECT client_id, secret FROM twitch_params WHERE user_id=1").Scan(&clientId, &secret)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}

	err = saveClient(clientId, secret, in.Code)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}

	return &google_protobuf.Empty{}, nil
}

// Save client id, secret key, access token and refresh token to database.
func saveClient(clientId, secret, code string) error {

	// Post request to get access token and refresh token
	type ResponseBody struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// Create HTTP POST request
	url := "https://id.twitch.tv/oauth2/token?client_id=" + clientId + "&client_secret=" + secret + "&grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:3000/twitch"

	var body []byte

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}

	// Get user id from Twitch
	type TwitchUser struct {
		Id       string `json:"id"`
		Username string `json:"login"`
	}
	type TwitchResponse struct {
		Data []TwitchUser `json:"data"`
	}

	url = "https://api.twitch.tv/helix/users"

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set Authorization header with bearer token
	req.Header.Set("Authorization", "Bearer "+responseBody.AccessToken)
	req.Header.Set("Client-Id", clientId)

	// Send HTTP request
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return err
	}

	// Read response body
	var twitchResp TwitchResponse

	err = json.NewDecoder(resp.Body).Decode(&twitchResp)
	if err != nil {
		return err
	}

	// Connect to db and create user with corresponding data.
	db, err := sql.Open("sqlite3", "files/data.db")
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE twitch_params SET access_token=$1, refresh_token=$2, twitch_user_name=$3, twitch_user_id=$4 WHERE user_id=1", responseBody.AccessToken, responseBody.RefreshToken, twitchResp.Data[0].Username, twitchResp.Data[0].Id)
	if err != nil {
		return err
	}

	return nil
}
