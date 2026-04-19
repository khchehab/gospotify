package gospotify

import (
	"context"
	"errors"
	"fmt"
)

// GetAlbum gets Spotify catalog information for a single album.
func (c *Client) GetAlbum(ctx context.Context, id string, opts ...QueryOption) (*AlbumObject, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var album AlbumObject
	if err := c.get(ctx, fmt.Sprintf("/albums/%s", id), &album, opts...); err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAlbumTracks gets Spotify catalog information about an album’s tracks.
func (c *Client) GetAlbumTracks(ctx context.Context, id string, opts ...QueryOption) (*Page[SimplifiedTrackObject], error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var tracks Page[SimplifiedTrackObject]
	if err := c.get(ctx, fmt.Sprintf("/albums/%s/tracks", id), &tracks, opts...); err != nil {
		return nil, err
	}
	return &tracks, nil
}

// GetUserSavedAlbums gets a list of the albums saved in the current Spotify user's 'Your Music' library.
func (c *Client) GetUserSavedAlbums(ctx context.Context, opts ...QueryOption) (*Page[SavedAlbumObject], error) {
	var savedAlbums Page[SavedAlbumObject]
	if err := c.get(ctx, "/me/albums", &savedAlbums, opts...); err != nil {
		return nil, err
	}
	return &savedAlbums, nil
}
