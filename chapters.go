package gospotify

import (
	"context"
	"fmt"
)

// GetChapter gets Spotify catalog information for a single audiobook chapter.
// Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetChapter(ctx context.Context, id string, opts ...QueryOption) (*ChapterObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var chapter ChapterObject
	if err := c.get(ctx, "/chapters/"+id, &chapter, opts...); err != nil {
		return nil, fmt.Errorf("GetChapter %q: %w", id, err)
	}
	return &chapter, nil
}
