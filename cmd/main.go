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

	ts, err := gospotify.AuthorizationCode(clientID, clientSecret, redirectURL, []string{"user-read-private", "user-top-read"}, 8080)
	if err != nil {
		fmt.Println("error getting token:", err)
	}

	c := gospotify.NewClient(ts)
	topTracks, err := c.GetUserTopTracks(context.Background(), gospotify.WithTimeRange(gospotify.LongTerm))
	if err != nil {
		fmt.Println("error getting top tracks:", err)
		return
	}
	fmt.Println(topTracks)
}
