package twitchApi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type ValidateAccessToken struct {
	ExpiresIn int    `json:"expires_in"` // The expiration time of the token is sent in seconds.
	Message   string `json:"message"`
}

// Validate access token.
func ValidateToken() (ValidateAccessToken, error) {
	// Check if access token exists in database.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return ValidateAccessToken{}, err
	}

	defer db.Close()

	var accessToken string
	err = db.QueryRow("SELECT access_token FROM users WHERE id = 1").Scan(&accessToken)
	if err == sql.ErrNoRows {
		err := errors.New("access_token not found. Generate it by running: twitcher api")
		return ValidateAccessToken{}, err
	} else if err != nil {
		return ValidateAccessToken{}, err
	}

	// Validate token with Twitch.
	url := "https://id.twitch.tv/oauth2/validate"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ValidateAccessToken{}, err
	}

	// Set Authorization header with bearer token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ValidateAccessToken{}, err
	}
	defer resp.Body.Close()

	// Read response body
	var responseBody ValidateAccessToken
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return ValidateAccessToken{}, err
	}

	// If the access token is invalid try to refresh it. If everything is OK this should give a new access token.
	if resp.StatusCode == 401 && responseBody.Message == "invalid access token" {
		token, err := RefreshToken()
		if err != nil {
			return ValidateAccessToken{}, err
		}

		return ValidateAccessToken{token.ExpiresIn, ""}, err

	} else if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("error. Couldn't validate access token. Status code: %v\n", resp.StatusCode)
		err := errors.New(errMsg)
		return ValidateAccessToken{}, err
	}

	log.Println("Access token successfully validated.")

	return responseBody, nil
}

// Save client id, secret key, access token and refresh token to database.
func SaveClient(username, clientId, secret, code string) error {

	// Post request to get access token and refresh token
	type ResponseBody struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// Create HTTP POST request
	url := "https://id.twitch.tv/oauth2/token?client_id=" + clientId + "&client_secret=" + secret + "&grant_type=authorization_code&code=" + code + "&redirect_uri=http://localhost:9696"

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

	// resp.Body.Close()

	// Get user id from Twitch
	type TwitchUser struct {
		Id string `json:"id"`
	}
	type TwitchResponse struct {
		Data []TwitchUser `json:"data"`
	}

	url = "https://api.twitch.tv/helix/users?login=" + username

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

	// resp.Body.Close()

	// Connect to db and create user with corresponding data.
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO users (client_id, secret, access_token, refresh_token, twitch_user_name, twitch_user_id) VALUES ($1,$2,$3,$4,$5,$6)", clientId, secret, responseBody.AccessToken, responseBody.RefreshToken, username, twitchResp.Data[0].Id)
	if err != nil {
		return err
	}

	return nil
}

// Refresh access token.
// The access token needs to be refreshed every 3 hours.
func RefreshToken() (ValidateAccessToken, error) {

	// Get refresh token, client id and client secret
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return ValidateAccessToken{}, err
	}

	defer db.Close()
	var clientId, secret, refreshToken string

	err = db.QueryRow("SELECT client_id, secret, refresh_token FROM users WHERE id = 1").Scan(&clientId, &secret, &refreshToken)
	if err != nil {
		return ValidateAccessToken{}, err
	}

	// Create HTTP POST request to get new access and refresh tokens.
	type ResponseBody struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}

	url := "https://id.twitch.tv/oauth2/token?grant_type=refresh_token&refresh_token=" + refreshToken + "&client_id=" + clientId + "&client_secret=" + secret

	var body []byte

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return ValidateAccessToken{}, err
	}

	// Set correct Content-Type.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ValidateAccessToken{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("couldn't create new access token and refresh token. Response status code: %v\n", resp.StatusCode)
		err := errors.New(errMsg)
		return ValidateAccessToken{}, err
	}

	// Read response body
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return ValidateAccessToken{}, err
	}

	resp.Body.Close()

	// Update rows with new data from the POST request
	_, err = db.Exec("UPDATE users SET access_token = $1, refresh_token = $2 WHERE id = 1", responseBody.AccessToken, responseBody.RefreshToken)
	if err != nil {
		return ValidateAccessToken{}, err
	}

	log.Println("Saved new access token and refresh token.")

	return ValidateAccessToken{responseBody.ExpiresIn, ""}, nil
}

func CreateEventSubs(sessionId, URL string) error {

	var twitchUserId, accessToken, clientId string

	// Get twitch_user_id, access_token and client_id from the database
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.QueryRow("SELECT twitch_user_id, access_token, client_id FROM users WHERE id = 1").Scan(&twitchUserId, &accessToken, &clientId)
	if err != nil {
		return err
	}

	// Create eventSub subscription for channel.follow
	// Create request body
	requestBody := bytes.NewBuffer([]byte(`
	{
		"type": "channel.follow", 
		"version": "2",
		"condition":{
			"broadcaster_user_id":"` + twitchUserId + `",
			"moderator_user_id":"` + twitchUserId + `"
		},
		"transport":{
			"method":"websocket",
			"session_id":"` + sessionId + `"
		}
	}`))

	req, err := http.NewRequest("POST", URL, requestBody)
	if err != nil {
		return err
	}

	// Set Authorization header with bearer token and set the Client-Id header.
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	type ResponseBody struct {
		Error   string `json:"error"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		err := errors.New(responseBody.Message)
		return err
	}

	// Create eventSub subscription for channel.subscribe
	// Create request body
	requestBody = bytes.NewBuffer([]byte(`
	{
		"type": "channel.subscribe", 
		"version": "1",
		"condition":{
			"broadcaster_user_id":"` + twitchUserId + `",
			"moderator_user_id":"` + twitchUserId + `"
		},
		"transport":{
			"method":"websocket",
			"session_id":"` + sessionId + `"
		}
	}`))
	req, err = http.NewRequest("POST", URL, requestBody)
	if err != nil {
		return err
	}

	// Set Authorization header with bearer token and set the Client-Id header.
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	if resp.StatusCode > 300 {
		errMsg := fmt.Sprintf("couldn't create channel.subscribe eventSub subscription. HTTP status code: %v", resp.StatusCode)
		err := errors.New(errMsg)
		return err
	}
	return nil
}
