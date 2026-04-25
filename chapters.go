package gospotify

import (
	"context"
	"errors"
)

// GetChapter gets Spotify catalog information for a single audiobook chapter.
// Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetChapter(ctx context.Context, id string, opts ...QueryOption) (*ChapterObject, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	var chapter ChapterObject
	if err := c.get(ctx, "/chapters/"+id, &chapter, opts...); err != nil {
		return nil, err
	}
	return &chapter, nil
}
