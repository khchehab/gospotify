package main

import (
	"context"
	"gospotify"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Hello World!")

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := os.Getenv("SPOTIFY_REDIRECT_URL")
	scopes := []gospotify.Scope{gospotify.ScopeUserLibraryRead, gospotify.ScopeUserFollowRead, gospotify.ScopePlaylistReadPrivate}

	a := gospotify.NewAuthenticatorWithServerPort(clientID, clientSecret, redirectURL, 8080, logger)
	ts, err := a.AuthorizationCode(scopes)
	if err != nil {
		logger.Error("error getting token", "error", err)
		return
	}

	c := gospotify.NewClient(ts, logger)
	_, err = c.CheckUserSavedItems(context.Background(), []string{"spotify:track:7a3LWj5xSFhFRYmztS8wgK", "spotify:album:4aawyAB9vmqN3uQ7FjRGTy", "spotify:artist:2takcwOaAZWiXQijPHIx7B"})
	if err != nil {
		logger.Error("failed to check user saved items", "error", err)
		return
	}

	logger.Info("done!")
}
