package gospotify

import (
	"context"
	"fmt"
)

// SearchForItem gets Spotify catalog information about albums, artists, playlists, tracks, shows, episodes or audiobooks that match a keyword string.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithLimit]: The maximum number of items to return.
//   - [WithOffset]: The index of the first item to return.
//   - [WithIncludeExternal]: If include_external=audio is specified, it signals that the client can play externally hosted audio content and marks the content as playable in the response.
//     By default, externally hosted audio content is marked as unplayable in the response.
func (c *Client) SearchForItem(ctx context.Context, q string, types []ItemType, opts ...QueryOption) (*SearchResult, error) {
	if err := requireNonEmpty("query", q); err != nil {
		return nil, err
	}
	if err := requireNonEmptyArray("types", types); err != nil {
		return nil, err
	}

	// validate the types values and convert them to an array of strings
	sTypes := make([]string, len(types))
	for i, t := range types {
		if !t.Valid() {
			return nil, fmt.Errorf("invalid item type: %s", t)
		}
		sTypes[i] = string(t)
	}

	endpoint := appendQueryParams("/search", requiredQueryParam{
		key:   "q",
		value: q,
	}, requiredQueryParam{
		key:   "types",
		value: sTypes,
	})

	var searchResult SearchResult
	if err := c.get(ctx, endpoint, &searchResult, opts...); err != nil {
		return nil, err
	}
	return &searchResult, nil
}
