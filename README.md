# Twitcher

Twitcher is a 24/7 Twitch broadcasting software with a web GUI. It can play video, sound and overlay layouts on top of the video. It is at a very early stage as it only fits my current needs. However, I have plans to add more features.

## Tech used:

- I decided to use Go and SQLite for the backend and Next.js (TypeScript) for the frontend. However, I really wanted to learn and experiment with grpc and protocol buffers, which resulted in using Next.js' server side rendering to communicate with the Go backend through grpc using protobufs due to the lack of grpc browser support. On the other hand, web browsers do support webRTC which is used in this case to showcase a live preview of the livestream.
- FFMPEG is used to handle the video, audio and image overlays for the encoding of the livestream and live preview.
- The Go backend communicates with Twitch's API to authenticate the user using OAuth 2.0 and to subscribe and then listen to WebSocket events coming from Twitch such as new follows or new subscribers.

## The current features are:

- Web GUI.
- 24/7 livestream.
- Live preview of livestream.
- Start preview without starting a livestream.
- Scrape ncs.io to download copyright free music to play in the livestream.
- Automatic creation of a playlist when starting preview or livestream.
- Automatic additon of songs to the playlist that is about to end.
- Move songs around in the playlist queue using drag and drop.
- Overlay that contains the song's name and cover image and the author's name. This overlay changes automatically when a new song plays.
- OAuth2 implemented to access Twitch api and receive alerts from Twitch that can be displayed on the livestream.

## TODO:

- Change hardcoded things that are suited for my specific needs. Such as the position of the song's overlay and alerts.
- Implement all Twitch alerts. Right now only alerts for new follows and subs are implemented.
- Select which moods and genres to scrape from ncs.io.
- Separately upload audio files.
- Add option to add video sources.
- Add option to create and manipulate audio playlists.
- Add YouTube as a stream option.
