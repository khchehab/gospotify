package gospotify

import (
	"context"
	"fmt"
)

// GetEpisode gets Spotify catalog information for a single episode identified by its unique Spotify ID.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetEpisode(ctx context.Context, id string, opts ...QueryOption) (*EpisodeObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var episode EpisodeObject
	if err := c.get(ctx, "/episodes/"+id, &episode, opts...); err != nil {
		return nil, fmt.Errorf("GetEpisode %q: %w", id, err)
	}
	return &episode, nil
}

// GetUserSavedEpisodes gets a list of the episodes saved in the current Spotify user's library.
//
// QueryOptions that can be used are:
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserSavedEpisodes(ctx context.Context, opts ...QueryOption) (*Page[SavedEpisodeObject], error) {
	var savedEpisodes Page[SavedEpisodeObject]
	if err := c.get(ctx, "/me/episodes", &savedEpisodes, opts...); err != nil {
		return nil, fmt.Errorf("GetUserSavedEpisodes: %w", err)
	}
	return &savedEpisodes, nil
}
