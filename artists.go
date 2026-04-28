package gospotify

import (
	"context"
)

// GetArtist gets Spotify catalog information for a single artist identified by their unique Spotify ID.
func (c *Client) GetArtist(ctx context.Context, id string) (*ArtistObject, error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var artist ArtistObject
	if err := c.get(ctx, "/artists/"+id, &artist); err != nil {
		return nil, err
	}
	return &artist, nil
}

// GetArtistAlbums gets Spotify catalog information about an artist's albums.
//
// QueryOptions that can be used are:
// * [WithIncludeGroups]: A list of keywords that will be used to filter the response.
// * [WithMarket]: An ISO 3166-1 alpha-2 country code.
// * [WithLimit]: The maximum number of items to return.
// * [WithOffset]: The index of the first item to return.
func (c *Client) GetArtistAlbums(ctx context.Context, id string, opts ...QueryOption) (*Page[ArtistDiscographyAlbumObject], error) {
	if err := requireNonEmpty("id", id); err != nil {
		return nil, err
	}
	var artistAlbums Page[ArtistDiscographyAlbumObject]
	if err := c.get(ctx, "/artists/"+id+"/albums", &artistAlbums, opts...); err != nil {
		return nil, err
	}
	return &artistAlbums, nil
}
