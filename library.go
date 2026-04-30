package gospotify

import (
	"context"
	"fmt"
)

// SaveItemsToLibrary adds one or more items to the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
func (c *Client) SaveItemsToLibrary(ctx context.Context, uris []string) error {
	if err := requireNonEmptyArray("uris", uris); err != nil {
		return err
	}

	endpoint := appendQueryParams("/me/library", requiredQueryParam{
		key:   "uris",
		value: uris,
	})

	if err := c.put(ctx, endpoint, nil, "", nil); err != nil {
		return fmt.Errorf("SaveItemsToLibrary [%q]: %w", uris, err)
	}
	return nil
}

// RemoveItemsFromLibrary removes one or more items from the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, users, and playlists.
func (c *Client) RemoveItemsFromLibrary(ctx context.Context, uris []string) error {
	if err := requireNonEmptyArray("uris", uris); err != nil {
		return err
	}

	endpoint := appendQueryParams("/me/library", requiredQueryParam{
		key:   "uris",
		value: uris,
	})

	if err := c.delete(ctx, endpoint, nil, nil); err != nil {
		return fmt.Errorf("RemoveItemsFromLibrary [%q]: %w", uris, err)
	}
	return nil
}

// CheckUserSavedItems checks if one or more items are already saved in the current user's library.
// Accepts Spotify URIs for tracks, albums, episodes, shows, audiobooks, artists, users, and playlists.
func (c *Client) CheckUserSavedItems(ctx context.Context, uris []string) ([]bool, error) {
	if err := requireNonEmptyArray("uris", uris); err != nil {
		return nil, err
	}

	endpoint := appendQueryParams("/me/library/contains", requiredQueryParam{
		key:   "uris",
		value: uris,
	})

	var savedItems []bool
	if err := c.get(ctx, endpoint, &savedItems); err != nil {
		return nil, fmt.Errorf("CheckUserSavedItems [%q]: %w", uris, err)
	}
	return savedItems, nil
}
