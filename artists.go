package gospotify

import (
	"context"
	"errors"
	"fmt"
)

// GetArtist gets Spotify catalog information for a single artist identified by their unique Spotify ID.
func (c *Client) GetArtist(ctx context.Context, id string) (*ArtistObject, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var artist ArtistObject
	if err := c.get(ctx, fmt.Sprintf("/artists/%s", id), &artist); err != nil {
		return nil, err
	}
	return &artist, nil
}

// GetArtistAlbums gets Spotify catalog information about an artist's albums.
func (c *Client) GetArtistAlbums(ctx context.Context, id string, opts ...QueryOption) (*Page[SimplifiedAlbumObject], error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var albums Page[SimplifiedAlbumObject]
	if err := c.get(ctx, fmt.Sprintf("/artists/%s/albums", id), &albums, opts...); err != nil {
		return nil, err
	}
	return &albums, nil
}
