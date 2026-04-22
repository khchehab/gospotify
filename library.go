package gospotify

import "context"

// SaveItemsToLibrary adds one or more items to the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
//
// QueryOptions that can be used are:
// * [WithURIs]: A (required) list of Spotify URIs.
func (c *Client) SaveItemsToLibrary(ctx context.Context, opts ...QueryOption) error {
	if err := c.put(ctx, "/me/library", opts...); err != nil {
		return err
	}
	return nil
}

// RemoveItemsFromLibrary removes one or more items from the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
//
// QueryOptions that can be used are:
// * [WithURIs]: A (required) list of Spotify URIs.
func (c *Client) RemoveItemsFromLibrary(ctx context.Context, opts ...QueryOption) error {
	if err := c.delete(ctx, "/me/library", opts...); err != nil {
		return err
	}
	return nil
}

// CheckUserSavedItems checks if one or more items are already saved in the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, artists, users, and playlists.
//
// QueryOptions that can be used are:
// * [WithURIs]: A (required) list of Spotify URIs.
func (c *Client) CheckUserSavedItems(ctx context.Context, opts ...QueryOption) ([]bool, error) {
	var savedItems []bool
	if err := c.get(ctx, "/me/library/contains", &savedItems, opts...); err != nil {
		return nil, err
	}
	return savedItems, nil
}
