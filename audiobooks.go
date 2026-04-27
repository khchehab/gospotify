package gospotify

import (
	"context"
)

// GetAudiobook gets Spotify catalog information for a single audiobook.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetAudiobook(ctx context.Context, id string, opts ...QueryOption) (*AudiobookObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var audiobook AudiobookObject
	if err := c.get(ctx, "/audiobooks/"+id, &audiobook, opts...); err != nil {
		return nil, err
	}
	return &audiobook, nil
}

// GetAudiobookChapters gets Spotify catalog information about an audiobook's chapters.
// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetAudiobookChapters(ctx context.Context, id string, opts ...QueryOption) (*Page[AudiobookChapter], error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var audiobookChapters Page[AudiobookChapter]
	if err := c.get(ctx, "/audiobooks/"+id+"/chapters", &audiobookChapters, opts...); err != nil {
		return nil, err
	}
	return &audiobookChapters, nil
}

// GetUserSavedAudiobooks gets a list of the audiobooks saved in the current Spotify user's 'Your Music' library.
//
// QueryOptions that can be used are:
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserSavedAudiobooks(ctx context.Context, opts ...QueryOption) (*Page[SavedAudiobookObject], error) {
	var savedAudiobooks Page[SavedAudiobookObject]
	if err := c.get(ctx, "/me/audiobooks/", &savedAudiobooks, opts...); err != nil {
		return nil, err
	}
	return &savedAudiobooks, nil
}
