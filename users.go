package gospotify

import (
	"context"
)

// GetCurrentUserProfile gets detailed profile information about the current user (including the current user's username).
func (c *Client) GetCurrentUserProfile(ctx context.Context) (*CurrentUserProfile, error) {
	var userProfile CurrentUserProfile
	if err := c.get(ctx, "/me", &userProfile); err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// GetUserTopArtists gets the current user's top artists based on calculated affinity.
//
// QueryOptions that can be used are:
// * [WithTimeRange]: Specify the time frame the affinities are computed.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserTopArtists(ctx context.Context, opts ...QueryOption) (*Page[UserTopArtists], error) {
	var topArtists Page[UserTopArtists]
	if err := c.get(ctx, "/me/top/artists", &topArtists, opts...); err != nil {
		return nil, err
	}
	return &topArtists, nil
}

// GetUserTopTracks gets the current user's top tracks based on calculated affinity.
//
// QueryOptions that can be used are:
// * [WithTimeRange]: Specify the time frame the affinities are computed.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserTopTracks(ctx context.Context, opts ...QueryOption) (*Page[UserTopTracks], error) {
	var topTracks Page[UserTopTracks]
	if err := c.get(ctx, "/me/top/tracks", &topTracks, opts...); err != nil {
		return nil, err
	}
	return &topTracks, nil
}

// GetFollowedArtists get the current user's followed artists.
//
// QueryOptions that can be used are:
// * [WithAfter]: The last artist ID retrieved from the previous request.
// * [WithLimit]: The maximum number of items to return.
func (c *Client) GetFollowedArtists(ctx context.Context, opts ...QueryOption) (*Cursor[FollowedArtists], error) {
	var response struct {
		Artists Cursor[FollowedArtists] `json:"artists"`
	}
	if err := c.get(ctx, "/me/following?type=artist", &response, opts...); err != nil {
		return nil, err
	}
	return &response.Artists, nil
}
