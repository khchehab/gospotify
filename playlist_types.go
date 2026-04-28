package gospotify

import (
	"encoding/json"
	"fmt"
)

type PlaylistItemsRefObject struct {
	// Href is a link to the Web API endpoint where full details of the playlist's tracks can be retrieved.
	Href string `json:"href"`
	// Total is the number of tracks in the playlist.
	Total int `json:"total"`
}

type PlaylistUserObject struct {
	// ExternalURLsObject is the known external URLs for this user.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
	// Href is a link to the Web API endpoint for this user.
	Href string `json:"href"`
	// ID is the Spotify user ID for the user.
	ID string `json:"id"`
	// Type is the object type: "user".
	Type string `json:"type"`
	// URI is the Spotify URI for the user.
	URI string `json:"uri"`
}

type PlaylistOwnerObject struct {
	PlaylistUserObject
	// DisplayName is the name displayed on the user's profile. null if not available.
	DisplayName *string `json:"display_name"`
}

type SimplifiedPlaylistObject struct {
	// Collaborative is true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// Description is the playlist description. Only returned for modified, verified playlists, otherwise null.
	Description *string `json:"description"`
	// ExternalURLs is the known external URLs for this playlist.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
	// Href is a link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// ID is the Spotify ID for the playlist
	ID string `json:"id"`
	// Images for the playlist.
	Images []ImageObject `json:"images"`
	// Name is the name of the playlist.
	Name string `json:"name"`
	// Owner is the user who owns the playlist
	Owner PlaylistOwnerObject `json:"owner"`
	// Public is the playlist's public/private status.
	Public *bool `json:"public"`
	// SnapshotID is the version identifier for the current playlist.
	SnapshotID string `json:"snapshot_id"`
	// Items is a collection containing a link [Href] to the Web API endpoint where full details of the playlist's items can be retrieved,
	// along with the total number of items in the playlist.
	// A track object may be null. This can happen if a track is no longer available.
	Items *PlaylistItemsRefObject `json:"items"`
	// Deprecated: Tracks is a collection containing a link [Href] to the Web API endpoint where full details of the playlist's items can be retrieved,
	// along with the total number of items in the playlist.
	// A track object may be null. This can happen if a track is no longer available.
	// Use [Items] instead.
	Tracks *PlaylistItemsRefObject `json:"tracks"`
	// Type is the object type: "playlist".
	Type string `json:"type"`
	// URI is the Spotify URI for the playlist.
	URI string `json:"uri"`
}

type PlaylistObject struct {
	// Collaborative is true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// Description is the playlist description. Only returned for modified, verified playlists, otherwise null.
	Description *string `json:"description"`
	// ExternalURLs is the known external URLs for this playlist.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
	// Href is a link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// ID is the Spotify ID for the playlist
	ID string `json:"id"`
	// Images for the playlist.
	Images []ImageObject `json:"images"`
	// Name is the name of the playlist.
	Name string `json:"name"`
	// Owner is the user who owns the playlist
	Owner PlaylistOwnerObject `json:"owner"`
	// Public is the playlist's public/private status.
	Public *bool `json:"public"`
	// SnapshotID is the version identifier for the current playlist.
	SnapshotID string `json:"snapshot_id"`
	// Items is the items of the playlist.
	Items Page[PlaylistTrackObject] `json:"items"`
	// Deprecated: Tracks is the tracks of the playlist. Use [Items] instead.
	Tracks *Page[PlaylistTrackObject] `json:"tracks"`
	// Type is the object type: "playlist".
	Type string `json:"type"`
	// URI is the Spotify URI for the playlist.
	URI string `json:"uri"`
}

type PlaylistTrackObject struct {
	// AddedAt is the date and time the track or episode was added. Note: some very old playlists may return null in this field.
	AddedAt *string `json:"added_at"`
	// AddedBy is the Spotify user who added the track or episode. Note: some very old playlists may return null in this field.
	AddedBy *PlaylistUserObject `json:"added_by"`
	// Local is whether this track or episode is a local file or not.
	Local bool `json:"is_local"`
	// Item is the information about the track or episode. TODO update documentation
	Track   *TrackObject
	Episode *EpisodeObject
}

func (p *PlaylistTrackObject) UnmarshalJSON(data []byte) error {
	var raw struct {
		AddedAt *string             `json:"added_at"`
		AddedBy *PlaylistUserObject `json:"added_by"`
		Local   bool                `json:"is_local"`
		Item    json.RawMessage     `json:"item"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.AddedAt = raw.AddedAt
	p.AddedBy = raw.AddedBy
	p.Local = raw.Local

	if raw.Item == nil {
		return nil
	}

	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw.Item, &peek); err != nil {
		return err
	}

	switch peek.Type {
	case "track":
		var t TrackObject
		if err := json.Unmarshal(raw.Item, &t); err != nil {
			return fmt.Errorf("failed to unmarshal track: %w", err)
		}
		p.Track = &t
	case "episode":
		var e EpisodeObject
		if err := json.Unmarshal(raw.Item, &e); err != nil {
			return fmt.Errorf("failed to unmarshal episode: %w", err)
		}
		p.Episode = &e
	default:
		return fmt.Errorf("unknown item type: %s", peek.Type)
	}

	return nil
}

type PlaylistDetailRequest struct {
	// Name is the new name for the playlist.
	Name string `json:"name"`
	// Public is the playlist's public/private status.
	Public *bool `json:"public"`
	// Collaborative if true, the playlist will become collaborative and other users will be able to modify the playlist in their Spotify client.
	// You can only set collaborative to true on non-public playlists.
	Collaborative bool `json:"collaborative"`
	// Description is the value for playlist description as displayed in Spotify Clients and in the Web API.
	Description string `json:"description"`
}

type PlaylistItemsRequest struct {
	// URIs is the list of Spotify URIs to set, can be track or episode URIs.
	URIs []string `json:"uris,omitempty"`
	// RangeStart is the position of the first item to be reordered.
	RangeStart *int `json:"range_start,omitempty"`
	// InsertBefore is the position where the items should be inserted.
	InsertBefore *int `json:"insert_before,omitempty"`
	// RangeLength is the amount of items to be reordered.
	RangeLength *int `json:"range_length,omitempty"`
	// SnapshotID is the playlist's snapshot ID against which you want to make the changes.
	SnapshotID *string `json:"snapshot_id,omitempty"`
}

type AddItemToPlaylistRequest struct {
	// URIs is an array of the Spotify URIs to add.
	URIs []string `json:"uris,omitempty"`
	// Position is the position to insert the items, a zero-based index.
	Position *int `json:"position,omitempty"`
}

type RemovePlaylistItemsRequest struct {
	// Items is an array of objects containing Spotify URIs of the tracks or episodes to remove.
	Items []struct {
		// URI is the Spotify URI.
		URI string `json:"uri"`
	} `json:"items"`
	// SnapshotID is the playlist's snapshot ID against which you want to make the changes.
	SnapshotID string `json:"snapshot_id"`
}

type CreatePlaylistRequest struct {
	// Name is the name for the new playlist.
	// This name does not need to be unique; a user may have several playlists with the same name.
	Name string `json:"name"`
	// Public is the playlist's public/private status. It defaults to true.
	Public *bool `json:"public,omitempty"`
	// Collaborative is true if the playlist will be collaborative. It defaults to false.
	Collaborative *bool `json:"collaborative,omitempty"`
	// Description is the value for playlist description as displayed in Spotify Clients and in the Web API.
	Description *string `json:"description,omitempty"`
}

type playlistOperationResponse struct {
	// SnapshotID is a snapshot ID for the playlist.
	SnapshotID string `json:"snapshot_id"`
}
