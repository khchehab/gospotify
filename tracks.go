package gospotify

import (
	"context"
	"errors"
	"fmt"
)

// GetTrack gets Spotify catalog information for a single track identified by its unique Spotify ID.
func (c *Client) GetTrack(ctx context.Context, id string, opts ...QueryOption) (*TrackObject, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var track TrackObject
	if err := c.get(ctx, fmt.Sprintf("/tracks/%s", id), &track, opts...); err != nil {
		return nil, err
	}
	return &track, nil
}

// GetUserSavedTracks gets a list of the songs saved in the current Spotify user's 'Your Music' library.
func (c *Client) GetUserSavedTracks(ctx context.Context, opts ...QueryOption) (*Page[SavedTrackObject], error) {
	var savedTracks Page[SavedTrackObject]
	if err := c.get(ctx, "/me/tracks", &savedTracks, opts...); err != nil {
		return nil, err
	}
	return &savedTracks, nil
}
