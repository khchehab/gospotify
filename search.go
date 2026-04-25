package gospotify

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
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
func (c *Client) SearchForItem(ctx context.Context, q string, types []string, opts ...QueryOption) (*SearchResult, error) {
	// TODO validate the parameters: q, types, or at least types with the allowed types.
	if q == "" {
		return nil, errors.New("query cannot be empty")
	}
	if len(types) == 0 {
		return nil, errors.New("types cannot be empty")
	}
	var searchResult SearchResult
	if err := c.get(ctx, concatenateSearch("/search", q, types), &searchResult, opts...); err != nil {
		return nil, err
	}
	return &searchResult, nil
}

// concatenateSearch concatenates the endpoint, the query and the types into a single URL.
func concatenateSearch(endpoint, q string, types []string) string {
	return fmt.Sprintf("%s?q=%s&types=%s", endpoint, url.QueryEscape(q), url.QueryEscape(strings.Join(types, ",")))
}
