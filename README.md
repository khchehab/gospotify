# gospotify

[![Go Reference](https://pkg.go.dev/badge/github.com/khchehab/gospotify.svg)](https://pkg.go.dev/github.com/khchehab/gospotify)
[![Go Report Card](https://goreportcard.com/badge/github.com/khchehab/gospotify)](https://goreportcard.com/report/github.com/khchehab/gospotify)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/khchehab/gospotify/main)](https://go.dev/doc/install)

A Go client library for the Spotify Web API with built-in OAuth2 authentication and automatic token refresh.

## Requirements

Go 1.26.0 or later.

## Installation

```sh
go get github.com/khchehab/gospotify
```

## Authentication

gospotify supports three OAuth2 flows. Choose based on your application's needs:

| Flow                         | Use when                                                                                                                    |
|------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| Authorization Code           | Server-side apps that can keep a client secret, and need access to user-specific data                                       |
| Authorization Code with PKCE | Public clients (CLIs, desktop apps, SPAs) that cannot securely store a client secret, and need access to user-specific data |
| Client Credentials           | Machine-to-machine access to public catalog data; no user context or user-specific scopes                                   |

All three flows return an `oauth2.TokenSource` that automatically refreshes the access token when it expires.

### Authorization Code

Starts a local HTTP server on port 8080 (configurable) to receive the redirect callback. The user's default browser is
opened automatically to complete login.

```go
auth := gospotify.NewAuthenticator(clientID, clientSecret, "http://localhost:8080/callback", nil)

ts, err := auth.AuthorizationCode([]gospotify.Scope{
  gospotify.ScopeUserReadPrivate,
  gospotify.ScopeUserLibraryRead,
})
```

Use `NewAuthenticatorWithServerPort` to bind the callback server to a different port.

### Authorization Code with PKCE

Same browser-redirect flow as above, but without a client secret. The `clientSecret` field is unused and can be left
empty.

```go
auth := gospotify.NewAuthenticator(clientID, "", "http://localhost:8080/callback", nil)

ts, err := auth.AuthorizationCodeWithPKCE([]gospotify.Scope{
  gospotify.ScopeUserReadPrivate,
})
```

### Client Credentials

No browser interaction. Returns a token source immediately.

```go
auth := gospotify.NewAuthenticator(clientID, clientSecret, "", nil)

ts, err := auth.ClientCredentials(nil)
```

## Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/khchehab/gospotify"
)

func main() {
	auth := gospotify.NewAuthenticator(
		"YOUR_CLIENT_ID",
		"YOUR_CLIENT_SECRET",
		"http://localhost:8080/callback",
		nil,
	)

	ts, err := auth.AuthorizationCode([]gospotify.Scope{
		gospotify.ScopeUserReadPrivate,
		gospotify.ScopeUserTopRead,
	})
	if err != nil {
		log.Fatal(err)
	}

	client := gospotify.NewClient(ts, nil)

	profile, err := client.GetCurrentUserProfile(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Logged in as: %s\n", profile.DisplayName)

	topTracks, err := client.GetUserTopTracks(context.Background(), gospotify.WithLimit(5))
	if err != nil {
		log.Fatal(err)
	}

	for _, track := range topTracks.Items {
		fmt.Printf("- %s\n", track.Name)
	}
}
```

Pass a `*slog.Logger` as the last argument to `NewAuthenticator` or `NewClient` to enable debug logging. If `nil` is
passed, all log output is discarded.

## Covered Endpoints

| Category   | Endpoints                                                                                                                                                                                                                                                                |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Albums     | Get Album, Get Album Tracks, Get User Saved Albums                                                                                                                                                                                                                       |
| Artists    | Get Artist, Get Artist Albums                                                                                                                                                                                                                                            |
| Audiobooks | Get Audiobook, Get Audiobook Chapters, Get User Saved Audiobooks                                                                                                                                                                                                         |
| Chapters   | Get Chapter                                                                                                                                                                                                                                                              |
| Episodes   | Get Episode, Get User Saved Episodes                                                                                                                                                                                                                                     |
| Library    | Save Items, Remove Items, Check Saved Items                                                                                                                                                                                                                              |
| Player     | Get Playback State, Transfer Playback, Get Available Devices, Get Currently Playing Track, Start/Resume/Pause Playback, Skip Next/Previous, Seek to Position, Set Repeat Mode, Set Volume, Toggle Shuffle, Get Recently Played Tracks, Get User Queue, Add Item to Queue |
| Playlists  | Get Playlist, Change Playlist Details, Get Playlist Items, Update Playlist Items, Add Items, Remove Items, Get Current User Playlists, Create Playlist, Get/Set Playlist Cover Image                                                                                     |
| Search     | Search for Items                                                                                                                                                                                                                                                         |
| Shows      | Get Show, Get Show Episodes, Get User Saved Shows                                                                                                                                                                                                                        |
| Tracks     | Get Track, Get User Saved Tracks                                                                                                                                                                                                                                         |
| Users      | Get Current User Profile, Get User Top Artists, Get User Top Tracks, Get Followed Artists                                                                                                                                                                                |

## Error Handling

API errors are returned as `*ErrorResponse`, which carries the HTTP status code and Spotify's error message:

```go
track, err := client.GetTrack(ctx, "invalid-id")

if err != nil {
  var apiErr *gospotify.ErrorResponse
  
  if errors.As(err, &apiErr) {
    fmt.Printf("Spotify error %d: %s\n", apiErr.ErrorObject.Status, apiErr.ErrorObject.Message)
  }
}
```

`ErrNoActivePlayback` is returned by `GetPlaybackState` when no playback session is currently active.

## License

MIT
