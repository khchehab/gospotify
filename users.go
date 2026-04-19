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
func (c *Client) GetUserTopArtists(ctx context.Context, opts ...TopItemsOption) (*Page[ArtistObject], error) {
	p := applyTopItems(opts...)

	var topArtists Page[ArtistObject]
	if err := c.get(ctx, fmt.Sprintf("/me/top/artists?%s", p.toQuery()), &topArtists); err != nil {
		return nil, err
	}
	return &topArtists, nil
}

// GetUserTopTracks gets the current user's top tracks based on calculated affinity.
func (c *Client) GetUserTopTracks(ctx context.Context, opts ...TopItemsOption) (*Page[TrackObject], error) {
	p := applyTopItems(opts...)

	var topTracks Page[TrackObject]
	if err := c.get(ctx, fmt.Sprintf("/me/top/tracks?%s", p.toQuery()), &topTracks); err != nil {
		return nil, err
	}
	return &topTracks, nil
}

// applyTopItems applies the options to the topItemsParameters.
func applyTopItems(opts ...TopItemsOption) topItemsParameters {
	p := topItemsParameters{
		timeRange: MediumTerm,
		limit:     20,
		offset:    0,
	}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}
