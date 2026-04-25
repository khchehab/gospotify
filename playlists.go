package gospotify

import (
	"context"
	"errors"
)

// GetPlaylist gets a playlist owned by a Spotify user.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithFields]: Filters for the query.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetPlaylist(ctx context.Context, playlistID string, opts ...QueryOption) (*PlaylistObject, error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}
	var playlist PlaylistObject
	if err := c.get(ctx, "/playlists/"+playlistID, &playlist, opts...); err != nil {
		return nil, err
	}
	return &playlist, nil
}

// ChangePlaylistDetails changes a playlist's name and public/private state (The user must, of course, own the playlist).
func (c *Client) ChangePlaylistDetails(ctx context.Context, playlistID string, body PlaylistDetailRequest) error {
	if playlistID == "" {
		return errors.New("playlist id cannot be empty")
	}
	if err := c.put(ctx, "/playlists/"+playlistID, body, "", nil); err != nil {
		return err
	}
	return nil
}

// GetPlaylistItems gets full details of the items of a playlist owned by a Spotify user.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithFields]: Filters for the query.
//   - [WithLimit]: The maximum number of items to return.
//   - [WithOffset]: The index of the first item to return.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetPlaylistItems(ctx context.Context, playlistID string, opts ...QueryOption) (*Page[PlaylistTrackObject], error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}
	var playlistItems Page[PlaylistTrackObject]
	if err := c.get(ctx, "/playlists/"+playlistID+"/items", &playlistItems, opts...); err != nil {
		return nil, err
	}
	return &playlistItems, nil
}

// UpdatePlaylistItems either reorder or replace items in a playlist depending on the request's parameters.
// To reorder items, include [range_start], [insert_before], [range_length] and [snapshot_id] in the request's body.
// To replace items, include uris as either a query parameter or in the request's body.
// Replacing items in a playlist will overwrite its existing items.
// This operation can be used for replacing or clearing items in a playlist.
//
// QueryOptions that can be used are:
//   - [WithURIs]: A list of Spotify URIs to set.
func (c *Client) UpdatePlaylistItems(ctx context.Context, playlistID string, body PlaylistItemsRequest, opts ...QueryOption) (*string, error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}

	var response playlistOperationResponse
	if err := c.put(ctx, "/playlists/"+playlistID+"/items", body, "", &response, opts...); err != nil {
		return nil, err
	}
	return &response.SnapshotID, nil
}

// AddItemsToPlaylist adds one or more items to a user's playlist.
//
// QueryOptions that can be used are:
//   - [WithPosition]: The position to insert the items, a zero-based index.
//   - [WithURIs]: A list of Spotify URIs to set.
func (c *Client) AddItemsToPlaylist(ctx context.Context, playlistID string, body AddItemToPlaylistRequest, opts ...QueryOption) (*string, error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}

	var response playlistOperationResponse
	if err := c.post(ctx, "/playlists/"+playlistID+"/items", body, &response, opts...); err != nil {
		return nil, err
	}
	return &response.SnapshotID, nil
}

// RemovePlaylistItems removes one or more items from a user's playlist.
func (c *Client) RemovePlaylistItems(ctx context.Context, playlistID string, body RemovePlaylistItemsRequest) (*string, error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}

	var response playlistOperationResponse
	if err := c.delete(ctx, "/playlists/"+playlistID+"/items", body, &response); err != nil {
		return nil, err
	}
	return &response.SnapshotID, nil
}

// GetCurrentUserPlaylists gets a list of the playlists owned or followed by the current Spotify user.
//
// QueryOptions that can be used are:
//   - [WithLimit]: The maximum number of items to return.
//   - [WithOffset]: The index of the first item to return.
func (c *Client) GetCurrentUserPlaylists(ctx context.Context, opts ...QueryOption) (*Page[SimplifiedPlaylistObject], error) {
	var playlist Page[SimplifiedPlaylistObject]
	if err := c.get(ctx, "/me/playlists", &playlist, opts...); err != nil {
		return nil, err
	}
	return &playlist, nil
}

// CreatePlaylist creates a playlist for the current Spotify user (The playlist will be empty until you add tracks).
// Each user is generally limited to a maximum of 11,000 playlists.
func (c *Client) CreatePlaylist(ctx context.Context, body CreatePlaylistRequest) (*PlaylistObject, error) {
	var playlist PlaylistObject
	if err := c.post(ctx, "/me/playlists", body, &playlist); err != nil {
		return nil, err
	}
	return &playlist, nil
}

// GetPlaylistCoverImage gets the current image associated with a specific playlist.
func (c *Client) GetPlaylistCoverImage(ctx context.Context, playlistID string) ([]ImageObject, error) {
	if playlistID == "" {
		return nil, errors.New("playlist id cannot be empty")
	}

	var images []ImageObject
	if err := c.get(ctx, "/playlists/"+playlistID+"/images", &images); err != nil {
		return nil, err
	}
	return images, nil
}

// AddCustomPlaylistCoverImage replaces the image used to represent a specific playlist.
func (c *Client) AddCustomPlaylistCoverImage(ctx context.Context, playlistID string, body []byte) error {
	if playlistID == "" {
		return errors.New("playlist id cannot be empty")
	}
	if len(body) == 0 {
		return errors.New("body cannot be empty")
	}

	return c.put(ctx, "/playlists/"+playlistID+"/images", body, "image/jpeg", nil)
}
