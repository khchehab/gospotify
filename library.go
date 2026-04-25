package gospotify

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// SaveItemsToLibrary adds one or more items to the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
func (c *Client) SaveItemsToLibrary(ctx context.Context, uris []string) error {
	if len(uris) == 0 {
		return errors.New("uris cannot be empty")
	}
	if err := c.put(ctx, concatenateURIs("/me/library", uris), nil, "", nil); err != nil {
		return err
	}
	return nil
}

// RemoveItemsFromLibrary removes one or more items from the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
func (c *Client) RemoveItemsFromLibrary(ctx context.Context, uris []string) error {
	if len(uris) == 0 {
		return errors.New("uris cannot be empty")
	}
	if err := c.delete(ctx, concatenateURIs("/me/library", uris)); err != nil {
		return err
	}
	return nil
}

// CheckUserSavedItems checks if one or more items are already saved in the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, artists, users, and playlists.
func (c *Client) CheckUserSavedItems(ctx context.Context, uris []string) ([]bool, error) {
	if len(uris) == 0 {
		return nil, errors.New("uris cannot be empty")
	}
	var savedItems []bool
	if err := c.get(ctx, concatenateURIs("/me/library/contains", uris), &savedItems); err != nil {
		return nil, err
	}
	return savedItems, nil
}

// concatenateURIs concatenates the given URIs (query parameter) into a single string with the endpoint.
func concatenateURIs(endpoint string, uris []string) string {
	return fmt.Sprintf("%s?uris=%s", endpoint, url.QueryEscape(strings.Join(uris, ",")))
}
