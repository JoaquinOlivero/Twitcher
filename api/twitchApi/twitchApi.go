package twitchApi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/sacOO7/gowebsocket"
)

// Twitch's websocket response.
type Response struct {
	Metadata struct {
		MessageId        string `json:"message_id"`
		MessageType      string `json:"message_type"`
		MessageTimestamp string `json:"message_timestamp"`
	} `json:"metadata"`

	Payload struct {
		Session struct {
			Id                      string `json:"id"`
			Status                  string `json:"status"`
			KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
			ReconnectURL            string `json:"reconnect_url"`
		} `json:"session"`

		Subscription struct {
			Id     string `json:"id"`
			Status string `json:"status"`
			Type   string `json:"type"`
		} `json:"subscription"`

		Event struct {
			UserId    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"event"`
	} `json:"payload"`
}

type SubscriptionChannel struct {
	Type     string
	Username string
}

type ValidateAccessToken struct {
	ExpiresIn int    `json:"expires_in"` // The expiration time of the token is sent in seconds.
	Message   string `json:"message"`
}

func Alerts(wgAlerts *sync.WaitGroup) error {
	// Go routine to handle alert notifications. The notifications come from Twitch's websocket api.
	wgAlertsI := 1

	// Go routine to validate access token and refresh it if necessary.
	// To subscribe to Twitch's websocket events it is required to have a valid user access token.
	var wg sync.WaitGroup
	wgI := 1
	wg.Add(wgI)

	go func() {
		for {

			token, err := validateToken()
			if err != nil {
				log.Println(err)
			}

			// Once there is a valid access token. Decrement the waitgroup count to zero and unblock the connection to Twitch's websocket server.
			if wgI == 1 {
				wg.Done()
				wgI--
			}

			// Hold loop until 10 minutes before the access token expires and then refresh the token.
			if token.ExpiresIn > 10*60 {
				sleepDuration := time.Duration(token.ExpiresIn - (10 * 60))
				time.Sleep(sleepDuration * time.Second)
				log.Println("Access token will expire in ten minutes. Getting a new one.")
			}

			refreshToken()
		}
	}()

	// wait until there is a valid acces token availabe to use.
	wg.Wait()

	// channels that will contain the type of notification and the username.
	channel := make(chan SubscriptionChannel, 300)
	stop := make(chan struct{}, 1)
	exit := make(chan struct{}, 1)

	// Go routine to catch notifications coming from the websocket connection to Twitch.
	go func() {
		// alert named pipe
		alertPipePath := "files/stream/alert"

		// Remove the named pipes if they already exists.
		err := os.Remove(alertPipePath)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}

		// Create named pipes.
		err = syscall.Mkfifo(alertPipePath, 0644)
		if err != nil {
			panic(err)
		}

		// Open alert named pipe to write to it.
		// Open named audio pipe
		alertPipe, err := os.OpenFile(alertPipePath, os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			panic(err)
		}

	alert:
		for {
			select {
			case notification := <-channel:

				// new follower alert. Draw the username on top of the alert with a fade in effect.
				if notification.Type == "channel.follow" {
					cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/follower-empty.webm",
						"-filter_complex", "[0:v]drawtext=fontfile=../../../Poppins-Bold.ttf:text='"+notification.Username+"':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2",
						"-c:v", "png", "-f", "image2pipe", "-")
					cmd.Stdout = alertPipe
					cmd.SysProcAttr = &syscall.SysProcAttr{
						Pdeathsig: syscall.SIGKILL,
					}

					cmd.Run()
				}

				// new sub alert. Draw the username on top of the alert with a fade in effect.
				if notification.Type == "channel.subscribe" {
					cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/sub-empty.webm",
						"-filter_complex", "[0:v]drawtext=fontfile=../../../Poppins-Bold.ttf:text='"+notification.Username+"':fontsize=16:fontcolor=ffffff:alpha='if(lt(t,0.5),0,if(lt(t,1.5),(t-0.5)/1,if(lt(t,10.5),1,if(lt(t,11),(0.5-(t-10.5))/0.5,0))))':x=(w-text_w)/2:y=(h-text_h)/2",
						"-c:v", "png", "-f", "image2pipe", "-")
					cmd.Stdout = alertPipe
					cmd.SysProcAttr = &syscall.SysProcAttr{
						Pdeathsig: syscall.SIGKILL,
					}

					cmd.Run()

				}

			case <-exit:
				break alert

			default:
				cmd := exec.Command("ffmpeg", "-hide_banner", "-re", "-stream_loop", "-1", "-c:v", "libvpx-vp9", "-i", "files/stream/alerts/empty.webm", "-c:v", "png", "-f", "image2pipe", "-")
				cmd.Stdout = alertPipe
				cmd.SysProcAttr = &syscall.SysProcAttr{
					Pdeathsig: syscall.SIGKILL,
				}

				// cmd.Run()
				cmd.Start()
			placeholder:
				for {
					select {
					case <-stop:
						cmd.Process.Signal(syscall.SIGKILL)
						break placeholder
					}
				}
				// cmd.Wait()
			}
		}
	}()

	var wsURL string
	wsURL = "wss://eventsub.wss.twitch.tv/ws"
	// wsURL = "ws://127.0.0.1:8080/ws"
	reconnect := make(chan struct{}, 1)

	for {

		//Create a client instance
		socket := gowebsocket.New(wsURL)

		socket.OnConnected = func(socket gowebsocket.Socket) {
			log.Println("Connected to server: ", wsURL)
		}

		socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
			log.Println("Received connect error ", err)
		}

		// Read received messages.
		socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
			var v Response
			err := json.Unmarshal([]byte(message), &v)
			if err != nil {
				log.Println(err)
			}

			// When connection to wss://eventsub.wss.twitch.tv/ws is established, Twitch replies with a welcome message that contains a session_id needed to subscribe to events.
			if v.Metadata.MessageType == "session_welcome" {
				// Send POST requests to create EventSub subscriptions for channel.follow and channel.subscribe.
				err := createEventSubs(v.Payload.Session.Id, "https://api.twitch.tv/helix/eventsub/subscriptions")
				// err := createEventSubs(v.Payload.Session.Id, "http://127.0.0.1:8080/eventsub/subscriptions")
				if err != nil {
					log.Println(err)
				}

				if wgAlertsI == 1 {
					fmt.Println("session_welcome")
					wgAlerts.Done()
					wgAlertsI--
				}
			}

			// Handling of notifications that come from the event channel.follow and channel.subscribe.
			// Go channels are used to handle the incoming notifications as a queue, one by one.
			if v.Metadata.MessageType == "notification" {
				var data SubscriptionChannel

				data.Type = v.Payload.Subscription.Type
				data.Username = v.Payload.Event.UserName
				stop <- struct{}{}
				channel <- data
			}

			// Twitch sends notification if the edge server the client is connected to needs to be swapped.
			if v.Metadata.MessageType == "session_reconnect" {
				wsURL = v.Payload.Session.ReconnectURL
				log.Println("Request from Twitch to reconnect websocket client to: ", wsURL)
				reconnect <- struct{}{}
			}
		}

		// This will send websocket handshake request to socketcluster-server
		socket.Connect()

		// Close socket connection and continue loop to reconnect to new server address.
		// This channel waiting also blocks the infinite loop from continuing, allowing to maintain the websocket connection and create a new one when needed.
		<-reconnect
		log.Println("reconnecting...")
		socket.Close()
		continue
	}
}

// Validate access token.
func validateToken() (ValidateAccessToken, error) {
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
		token, err := refreshToken()
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
func refreshToken() (ValidateAccessToken, error) {

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

func createEventSubs(sessionId, URL string) error {

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
