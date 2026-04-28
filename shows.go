package gospotify

import (
	"context"
)

// GetShow gets Spotify catalog information for a single show identified by its unique Spotify ID.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetShow(ctx context.Context, id string, opts ...QueryOption) (*ShowObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var show ShowObject
	if err := c.get(ctx, "/shows/"+id, &show, opts...); err != nil {
		return nil, err
	}
	return &show, nil
}

// GetShowEpisodes gets Spotify catalog information about a show’s episodes.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetShowEpisodes(ctx context.Context, id string, opts ...QueryOption) (*Page[SimplifiedEpisodeObject], error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var showEpisodes Page[SimplifiedEpisodeObject]
	if err := c.get(ctx, "/shows/"+id+"/episodes", &showEpisodes, opts...); err != nil {
		return nil, err
	}
	return &showEpisodes, nil
}

// GetUserSavedShows gets a list of shows saved in the current Spotify user's library.
//
// QueryOptions that can be used are:
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserSavedShows(ctx context.Context, opts ...QueryOption) (*Page[SavedShowObject], error) {
	var savedShows Page[SavedShowObject]
	if err := c.get(ctx, "/me/shows/", &savedShows, opts...); err != nil {
		return nil, err
	}
	return &savedShows, nil
}
