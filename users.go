package gospotify

import (
	"context"
	"fmt"
)

// GetCurrentUserProfile gets detailed profile information about the current user (including the current user's username).
func (c *Client) GetCurrentUserProfile(ctx context.Context) (*UserProfile, error) {
	var userProfile UserProfile
	if err := c.get(ctx, "/me", &userProfile); err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// GetUserTopArtists gets the current user's top artists based on calculated affinity.
func (c *Client) GetUserTopArtists(ctx context.Context, opts ...QueryOption) (*Page[ArtistObject], error) {
	p := applyQueryParameters(opts...)

	var topArtists Page[ArtistObject]
	if err := c.get(ctx, fmt.Sprintf("/me/top/artists?%s", p.toQuery()), &topArtists); err != nil {
		return nil, err
	}
	return &topArtists, nil
}

// GetUserTopTracks gets the current user's top tracks based on calculated affinity.
func (c *Client) GetUserTopTracks(ctx context.Context, opts ...QueryOption) (*Page[TrackObject], error) {
	p := applyQueryParameters(opts...)

	var topTracks Page[TrackObject]
	if err := c.get(ctx, fmt.Sprintf("/me/top/tracks?%s", p.toQuery()), &topTracks); err != nil {
		return nil, err
	}
	return &topTracks, nil
}

// GetFollowedArtists get the current user's followed artists.
func (c *Client) GetFollowedArtists(ctx context.Context, opts ...QueryOption) (*FollowedArtists, error) {
	p := applyQueryParameters(opts...)
	endpoint := "/me/following?type=artist"
	if q := p.toQuery(); q != "" {
		endpoint += "&" + q
	}

	var response followedArtistsResponse
	if err := c.get(ctx, endpoint, &response); err != nil {
		return nil, err
	}
	return &response.Artists, nil
}
