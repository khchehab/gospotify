# gospotify

A Go client library for the Spotify Web API with built-in OAuth2 authentication and automatic token refresh.

This wrapper does not implement deprecated endpoints, but deprecated fields will still be taken into consideration.

## TODO

Below is a TODO list of endpoints to implement.

- [ ] Implement endpoints
  - [x] Albums
  - [x] Artists
  - [x] Audiobooks
  - [x] Chapters
  - [x] Episodes
  - [x] Library
  - [ ] Player
  - [x] Playlists
  - [x] Search
  - [x] Shows
  - [x] Tracks
  - [x] Users
- [ ] Test the endpoints in some way, whether manually or automated, though preferable automated.
- [ ] Refactor code to abstract common functionalities (especially in Client struct).
- [ ] Find a way to reduce the parameters or make them pluggable.
