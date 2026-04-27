package gospotify

import (
	"context"
)

// GetTrack gets Spotify catalog information for a single track identified by its unique Spotify ID.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetTrack(ctx context.Context, id string, opts ...QueryOption) (*TrackObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var track TrackObject
	if err := c.get(ctx, "/tracks/"+id, &track, opts...); err != nil {
		return nil, err
	}
	return &track, nil
}

// GetUserSavedTracks gets a list of the songs saved in the current Spotify user's 'Your Music' library.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserSavedTracks(ctx context.Context, opts ...QueryOption) (*Page[SavedTrackObject], error) {
	var savedTracks Page[SavedTrackObject]
	if err := c.get(ctx, "/me/tracks", &savedTracks, opts...); err != nil {
		return nil, err
	}
	return &savedTracks, nil
}
