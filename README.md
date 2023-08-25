# Twitcher

Twitcher is a 24/7 Twitch broadcasting software with a web GUI. It can play video, sound and overlay layouts on top of the video.

## Project status

This project is at a very early stage as it only fits my current needs. However, I am currently working in improving its current status and adding new features.

## Tech used

- I decided to use Go and SQLite for the backend and Next.js (TypeScript) for the frontend. However, I really wanted to learn and experiment with grpc and protocol buffers, which resulted in using Next.js' server side rendering to communicate with the Go backend through grpc using protobufs due to the lack of grpc browser support. On the other hand, web browsers do support webRTC which is used in this case to showcase a live preview of the livestream.
- FFMPEG is used to handle the video, audio and image overlays for the encoding of the livestream and live preview.
- The Go backend communicates with Twitch's API to authenticate the user using OAuth 2.0 and to subscribe and then listen to WebSocket events coming from Twitch such as new follows or new subscribers.

## Current features

- Web GUI.
- 24/7 livestream.
- Live preview of livestream.
- Start preview without starting a livestream.
- Scrape ncs.io to download copyright free music.
- Automatic creation of a playlist when starting preview or livestream.
- Automatic additon of songs to the playlist that is about to end.
- Move songs around in the playlist queue using drag and drop.
- Image overlay that contains the song's name, cover image and the author's name. This overlay changes automatically when a new song plays and its elements can be customized (change image width, height, x and y position, font color, font size and text alignment).
- Swap active background video during livestream.
- Upload background video files.
- OAuth2 implemented to access Twitch api and receive alerts from Twitch that can be displayed on the livestream.

## Project Screenshot
![Twitcher GUI screenshot](/images/twitcher_screenshot.png)

## GUI Example Clips
#### Start stream and swap background video.
<img src="images/1b.gif" alt="drawing" width="350"/> <img src="images/1a.gif" alt="drawing" width="350"/>

#### Customize default song overlay.
<img src="images/3b.gif" alt="drawing" width="350"/> <img src="images/3a.gif" alt="drawing" width="350"/>

<img src="images/6b.gif" alt="drawing" width="350"/> <img src="images/6a.gif" alt="drawing" width="350"/>

## Installation and Setup Instructions

NOTE: This project has only been tested on Linux amd64/arm64.

Clone this repository.

### Ffmpeg build
A build of ffmpeg with libzmq is required to run the program. The script "ffmpeg_compile.sh" creates a build of ffmpeg with libzmq.
```
sudo chmod +x ffmpeg_compile.sh
```

```
./ffmpeg_compile.sh 
```

### Install
```
cd Twitcher && cd api && go get ./... && cd ../src && npm install && npm run build && cd ../
```

### Run

#### Golang backend - Twitcher/api directory.
The grpc server runs on port :9000 and the song covers are statically served on port :9001.

```
go run .
```

#### Next.js - Twitcher/src directory.
Runs in port :3000 by default.

```
npm run start
```

Change port
```
npm run start -- -p 3001
```

## TODO

- Implement all Twitch alerts. Right now only alerts for new follows and subs are implemented.
- Select which moods and genres to scrape from ncs.io.
- Separately upload audio files.
- Add option to create and manipulate audio playlists.
- Add YouTube as a stream option.
- Add option to change the text font in the song overlay.
- Add flag to change default ports for the grpc and static files servers.