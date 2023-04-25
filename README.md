# Twitcher

Twitcher is a CLI based 24/7 streaming program. It can play video, sound and overlay layouts on top of the video. It is at a very early stage as it only fits my current needs. However, I have plans to make a front-end web gui to properly manage the stream.


The current features are:
- Scrape ncs.io to download copyright free music to play in the stream.
- Randomly create playlists to play during stream with the available songs.
- Overlay that contains the song's name and cover image and the author's name. This overlay changes automatically when a new song plays.
- OAuth2 implemented to access the Twitch api and receive alerts from Twitch using websockets.


TODO:
- Implement all Twitch alerts. Right now only alerts for new follows and subs are implemented.
- Change hardcoded things that are suited for my specific needs. Such as the position of the song's overlay and alerts. 
- Add YouTube as a stream option.
- Create a web GUI.

Creating a web gui will require some code re-writing.
The first thing that needs to be implemented is a simple web gui that would just start and stop the stream, show some information about the stream and then build more features in the gui to have more control over the stream, the music scraper and CRUD related operations.

TODO for when the web gui is ready:
- Select which moods and genres to scrape from ncs.io.
- Separately upload external songs or audio files.
- Add option to create and manipulate audio and video playlists.
- Preview stream.
