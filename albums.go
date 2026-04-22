package gospotify

import "context"

// GetAlbum gets Spotify catalog information for a single album.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
func (c *Client) GetAlbum(ctx context.Context, id string, opts ...QueryOption) (*AlbumObject, error) {
	var album AlbumObject
	if err := c.get(ctx, "/albums/"+id, &album, opts...); err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAlbumTracks gets Spotify catalog information about an album’s tracks. Optional parameters can be used to limit the number of tracks returned.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetAlbumTracks(ctx context.Context, id string, opts ...QueryOption) (*Page[AlbumTrack], error) {
	var albumTracks Page[AlbumTrack]
	if err := c.get(ctx, "/albums/"+id+"/tracks", &albumTracks, opts...); err != nil {
		return nil, err
	}
	return &albumTracks, nil
}

// GetUserSavedAlbums gets a list of the albums saved in the current Spotify user's 'Your Music' library.
//
// QueryOptions that can be used are:
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetUserSavedAlbums(ctx context.Context, opts ...QueryOption) (*Page[SavedAlbumObject], error) {
	var savedAlbums Page[SavedAlbumObject]
	if err := c.get(ctx, "/me/albums", &savedAlbums, opts...); err != nil {
		return nil, err
	}
	return &savedAlbums, nil
}
