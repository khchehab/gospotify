package main

import (
	"context"
	"fmt"
	"gospotify"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := os.Getenv("SPOTIFY_REDIRECT_URL")
	scopes := []string{"user-library-read", "user-follow-read", "playlist-read-private"}

	ts, err := gospotify.AuthorizationCode(clientID, clientSecret, redirectURL, scopes, 8080)
	if err != nil {
		fmt.Println("error getting token:", err)
	}

	c := gospotify.NewClient(ts)
	items, err := c.CheckUserSavedItems(context.Background(), gospotify.WithURIs("spotify:track:7a3LWj5xSFhFRYmztS8wgK", "spotify:album:4aawyAB9vmqN3uQ7FjRGTy", "spotify:artist:2takcwOaAZWiXQijPHIx7B"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(items)
	fmt.Println("done!")
}
