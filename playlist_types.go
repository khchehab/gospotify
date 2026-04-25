package gospotify

import (
	"encoding/json"
	"fmt"
)

type PlaylistObject struct {
	// Collaborative is true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// Description is the playlist description. Only returned for modified, verified playlists, otherwise null.
	Description *string `json:"description"`
	// ExternalURLs is the known external URLs for this playlist.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Href is a link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// ID is the Spotify ID for the playlist
	ID string `json:"id"`
	// Images for the playlist.
	Images []ImageObject `json:"images"`
	// Name is the name of the playlist.
	Name string `json:"name"`
	// Owner is the user who owns the playlist
	Owner struct {
		// ExternalURLs is the known external URLs for this user.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint for this user.
		Href string `json:"href"`
		// ID is the Spotify user ID for the user.
		ID string `json:"id"`
		// Type is the object type: "user".
		Type string `json:"type"`
		// URI is the Spotify URI for the user.
		URI string `json:"uri"`
		// DisplayName is the name displayed on the user's profile. null if not available.
		DisplayName *string `json:"display_name"`
	} `json:"owner"`
	// Public is the playlist's public/private status.
	Public *bool `json:"public"`
	// SnapshotID is the version identifier for the current playlist.
	SnapshotID string `json:"snapshot_id"`
	// Items is the items of the playlist.
	Items Page[PlaylistTrackObject] `json:"items"`
	// Type is the object type: "playlist".
	Type string `json:"type"`
	// URI is the Spotify URI for the playlist.
	URI string `json:"uri"`
}

type SimplifiedPlaylistObject struct {
	// Collaborative is true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// Description is the playlist description. Only returned for modified, verified playlists, otherwise null.
	Description *string `json:"description"`
	// ExternalURLs is the known external URLs for this playlist.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Href is a link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// ID is the Spotify ID for the playlist
	ID string `json:"id"`
	// Images for the playlist.
	Images []ImageObject `json:"images"`
	// Name is the name of the playlist.
	Name string `json:"name"`
	// Owner is the user who owns the playlist
	Owner struct {
		// ExternalURLs is the known external URLs for this user.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint for this user.
		Href string `json:"href"`
		// ID is the Spotify user ID for the user.
		ID string `json:"id"`
		// Type is the object type: "user".
		Type string `json:"type"`
		// URI is the Spotify URI for the user.
		URI string `json:"uri"`
		// DisplayName is the name displayed on the user's profile. null if not available.
		DisplayName *string `json:"display_name"`
	} `json:"owner"`
	// Public is the playlist's public/private status.
	Public *bool `json:"public"`
	// SnapshotID is the version identifier for the current playlist.
	SnapshotID string `json:"snapshot_id"`
	// Items is a collection containing a link [Href] to the Web API endpoint where full details of the playlist's items can be retrieved,
	// along with the total number of items in the playlist.
	// A track object may be null. This can happen if a track is no longer available.
	Items *struct {
		// Href is a link to the Web API endpoint where full details of the playlist's tracks can be retrieved.
		Href string `json:"href"`
		// Total is the number of tracks in the playlist.
		Total int `json:"total"`
	} `json:"items"`
	// Type is the object type: "playlist".
	Type string `json:"type"`
	// URI is the Spotify URI for the playlist.
	URI string `json:"uri"`
}

type PlaylistTrackObject struct {
	// AddedAt is the date and time the track or episode was added. Note: some very old playlists may return null in this field.
	AddedAt *string `json:"added_at"`
	// AddedBy is the Spotify user who added the track or episode. Note: some very old playlists may return null in this field.
	AddedBy *struct {
		// ExternalURLs is the known external URLs for this user.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint for this user.
		Href string `json:"href"`
		// ID is the Spotify user ID for the user.
		ID string `json:"id"`
		// Type is the object type: "user".
		Type string `json:"type"`
		// URI is the Spotify URI for the user.
		URI string `json:"uri"`
	} `json:"added_by"`
	// Local is whether this track or episode is a local file or not.
	Local bool `json:"is_local"`
	// Item is the information about the track or episode.
	Track *struct {
		// Album is the album on which the track appears.
		// The album object includes a link in href to full information about the album.
		Album struct {
			// AlbumType is the type of the album. Allowed values are "album", "single", or "compilation".
			AlbumType string `json:"album_type"`
			// TotalTracks is the number of tracks in the album.
			TotalTracks int `json:"total_tracks"`
			// Deprecated: AvailableMarkets is the markets in which the album is available: ISO 3166-1 alpha-2 country codes.
			AvailableMarkets []string `json:"available_markets"`
			// ExternalURLs is the known external URLs for this album.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// Href is a link to the Web API endpoint providing full details of the album.
			Href string `json:"href"`
			// ID is the Spotify ID for the album.
			ID string `json:"id"`
			// Images is the cover art for the album in various sizes, widest first.
			Images []ImageObject `json:"images"`
			// Name is the name of the album. In the case of an album takedown, the value may be an empty string.
			Name string `json:"name"`
			// ReleaseDate is the date the album was first released.
			ReleaseDate string `json:"release_date"`
			// ReleaseDatePrecision is the precision with which ReleaseDate value is known. Allowed values are "year", "month", or "day".
			ReleaseDatePrecision string `json:"release_date_precision"`
			// Restrictions are Included in the response when a content restriction is applied.
			Restrictions *Restrictions `json:"restrictions"`
			// Type is the object type. Allowed values is "album".
			Type string `json:"type"`
			// URI is the Spotify URI for the album.
			URI string `json:"uri"`
			// Artists are the artists of the album. Each artist object includes a link in href to more detailed information about the artist.
			Artists []struct { // called SimplifiedArtistObject
				// ExternalURLs is the known external URLs for this artist.
				ExternalURLs ExternalURLs `json:"external_urls"`
				// Href is a link to the Web API endpoint providing full details of the artist.
				Href string `json:"href"`
				// ID is the Spotify ID for the artist.
				ID string `json:"id"`
				// Name is the name of the artist.
				Name string `json:"name"`
				// Type is the object type.
				Type string `json:"type"`
				// URI is the Spotify URI for the artist.
				URI string `json:"uri"`
			} `json:"artists"`
		} `json:"album"`
		// Artists is the artists who performed the track.
		// Each artist object includes a link in href to more detailed information about the artist.
		Artists []struct {
			// ExternalURLs is the known external URLs for this artist.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// Href is a link to the Web API endpoint providing full details of the artist.
			Href string `json:"href"`
			// ID is the Spotify ID for the artist.
			ID string `json:"id"`
			// Name is the name of the artist.
			Name string `json:"name"`
			// Type is the object type.
			Type string `json:"type"`
			// URI is the Spotify URI for the artist.
			URI string `json:"uri"`
		} `json:"artists"`
		// Deprecated: AvailableMarkets is a list of the countries in which the track can be played.
		AvailableMarkets []string `json:"available_markets"`
		// DiscNumber is the disc number (usually 1 unless the album consists of more than one disc).
		DiscNumber int `json:"disc_number"`
		// DurationMs is the track length in milliseconds.
		DurationMs int `json:"duration_ms"`
		// Explicit is whether the track has explicit lyrics.
		Explicit bool `json:"explicit"`
		// ExternalIDs is the known external IDs for the track.
		ExternalIDs ExternalIDs `json:"external_ids"`
		// ExternalURLs is the known external URLs for this track.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint providing full details of the track.
		Href string `json:"href"`
		// ID is the Spotify ID for the track.
		ID string `json:"id"`
		// Playable is true if the track is playable in the given market, otherwise false.
		Playable bool `json:"is_playable"`
		// Deprecated: LinkedFrom is part of the response when Track Relinking is applied, and the requested track has been replaced with different track.
		LinkedFrom map[string]any `json:"linked_from"`
		// Restrictions are included in the response when a content restriction is applied.
		Restrictions *Restrictions `json:"restrictions"`
		// Name is the name of the track.
		Name string `json:"name"`
		// Deprecated: Popularity is the popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
		Popularity int `json:"popularity"`
		// Deprecated: PreviewURL is a link to a 30-second preview (MP3 format) of the track. Can be null
		PreviewURL *string `json:"preview_url"`
		// TrackNumber is the number of the track. If an album has several discs, the track number is the number on the specified disc.
		TrackNumber int `json:"track_number"`
		// Type is the object type: "track".
		Type string `json:"type"`
		// URI is the Spotify URI for the track.
		URI string `json:"uri"`
		// Local is whether the track is from a local file.
		Local bool `json:"is_local"`
	}
	Episode *struct {
		// Deprecated: AudioPreviewURL is a URL to a 30-second preview (MP3 format) of the episode. null if not available.
		AudioPreviewURL *string `json:"audio_preview_url"`
		// Description is a description of the episode. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
		Description string `json:"description"`
		// HTMLDescription is a description of the episode. This field may contain HTML tags.
		HTMLDescription string `json:"html_description"`
		// DurationMs is the episode length in milliseconds.
		DurationMs int `json:"duration_ms"`
		// Explicit is whether the episode has explicit content.
		Explicit bool `json:"explicit"`
		// ExternalURLs is the external URLs for this episode.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint providing full details of the episode.
		Href string `json:"href"`
		// ID is the Spotify ID for the episode.
		ID string `json:"id"`
		// Images is the cover art for the episode in various sizes, widest first.
		Images []ImageObject `json:"images"`
		// ExternallyHosted is true if the episode is hosted outside Spotify's CDN.
		ExternallyHosted bool `json:"is_externally_hosted"`
		// Playable is True if the episode is playable in the given market, otherwise false.
		Playable bool `json:"is_playable"`
		// Deprecated: Language is the language used in the episode, identified by an ISO 639 code.
		Language *string `json:"language"`
		// Languages is a list of the languages used in the episode, identified by their ISO 639-1 code.
		Languages []string `json:"languages"`
		// Name is the name of the episode.
		Name string `json:"name"`
		// ReleaseDate is the date the episode was first released.
		ReleaseDate string `json:"release_date"`
		// ReleaseDatePrecision is the precision with which release_date value is known. Allowed values are "year", "month", or "day".
		ReleaseDatePrecision string `json:"release_date_precision"`
		// ResumePoint is the user's most recent position in the episode.
		// Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
		ResumePoint struct {
			// FullyPlayed is whether the episode has been fully played by the user.
			FullyPlayed bool `json:"fully_played"`
			// ResumePositionMs is the user's most recent position in the episode in milliseconds.
			ResumePositionMs int `json:"resume_position_ms"`
		} `json:"resume_point"`
		// Type is the object type. Allowed values: "episode".
		Type string `json:"type"`
		// URI is the Spotify URI for the episode.
		URI string `json:"uri"`
		// Restrictions is included in the response when a content restriction is applied.
		Restrictions *Restrictions `json:"restrictions"`
		// Show is the show on which the episode belongs.
		Show struct {
			// Deprecated: AvailableMarkets is a list of the countries in which the show can be played, identified by their ISO 3166-1 alpha-2 code.
			AvailableMarkets []string `json:"available_markets"`
			// Copyrights is the copyright statements of the show.
			Copyrights []CopyrightObject `json:"copyrights"`
			// Description is a description of the show. HTML tags are stripped away from this field, use HTMLDescription field in case HTML tags are needed.
			Description string `json:"description"`
			// HTMLDescription is a description of the show. This field may contain HTML tags.
			HTMLDescription string `json:"html_description"`
			// Explicit is whether the show has explicit content.
			Explicit bool `json:"explicit"`
			// ExternalURLs is the external URLs for this show.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// Href is a link to the Web API endpoint providing full details of the show.
			Href string `json:"href"`
			// ID is the Spotify ID for the show.
			ID string `json:"id"`
			// Images is the cover art for the show in various sizes, widest first.
			Images []ImageObject `json:"images"`
			// ExternallyHosted is true if all shows episodes are hosted outside Spotify's CDN.
			ExternallyHosted *bool `json:"is_externally_hosted"`
			// Languages is a list of the languages used in the show, identified by their ISO 639 code.
			Languages []string `json:"languages"`
			// MediaType is the media type of the show.
			MediaType string `json:"media_type"`
			// Name is the name of the show.
			Name string `json:"name"`
			// Deprecated: Publisher is the publisher of the show.
			Publisher *string `json:"publisher"`
			// Type is the object type. Allowed values is "show".
			Type string `json:"type"`
			// URI is the Spotify URI for the show.
			URI string `json:"uri"`
			// TotalEpisodes is the number of episodes in this show.
			TotalEpisodes int `json:"total_episodes"`
		} `json:"show"`
	}
}

func (p *PlaylistTrackObject) UnmarshalJSON(data []byte) error {
	var raw struct {
		AddedAt *string `json:"added_at"`
		AddedBy *struct {
			ExternalURLs ExternalURLs `json:"external_urls"`
			Href         string       `json:"href"`
			ID           string       `json:"id"`
			Type         string       `json:"type"`
			URI          string       `json:"uri"`
		} `json:"added_by"`
		Local bool            `json:"is_local"`
		Item  json.RawMessage `json:"item"`
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
		var t struct {
			// Album is the album on which the track appears.
			// The album object includes a link in href to full information about the album.
			Album struct {
				// AlbumType is the type of the album. Allowed values are "album", "single", or "compilation".
				AlbumType string `json:"album_type"`
				// TotalTracks is the number of tracks in the album.
				TotalTracks int `json:"total_tracks"`
				// Deprecated: AvailableMarkets is the markets in which the album is available: ISO 3166-1 alpha-2 country codes.
				AvailableMarkets []string `json:"available_markets"`
				// ExternalURLs is the known external URLs for this album.
				ExternalURLs ExternalURLs `json:"external_urls"`
				// Href is a link to the Web API endpoint providing full details of the album.
				Href string `json:"href"`
				// ID is the Spotify ID for the album.
				ID string `json:"id"`
				// Images is the cover art for the album in various sizes, widest first.
				Images []ImageObject `json:"images"`
				// Name is the name of the album. In the case of an album takedown, the value may be an empty string.
				Name string `json:"name"`
				// ReleaseDate is the date the album was first released.
				ReleaseDate string `json:"release_date"`
				// ReleaseDatePrecision is the precision with which ReleaseDate value is known. Allowed values are "year", "month", or "day".
				ReleaseDatePrecision string `json:"release_date_precision"`
				// Restrictions are Included in the response when a content restriction is applied.
				Restrictions *Restrictions `json:"restrictions"`
				// Type is the object type. Allowed values is "album".
				Type string `json:"type"`
				// URI is the Spotify URI for the album.
				URI string `json:"uri"`
				// Artists are the artists of the album. Each artist object includes a link in href to more detailed information about the artist.
				Artists []struct { // called SimplifiedArtistObject
					// ExternalURLs is the known external URLs for this artist.
					ExternalURLs ExternalURLs `json:"external_urls"`
					// Href is a link to the Web API endpoint providing full details of the artist.
					Href string `json:"href"`
					// ID is the Spotify ID for the artist.
					ID string `json:"id"`
					// Name is the name of the artist.
					Name string `json:"name"`
					// Type is the object type.
					Type string `json:"type"`
					// URI is the Spotify URI for the artist.
					URI string `json:"uri"`
				} `json:"artists"`
			} `json:"album"`
			// Artists is the artists who performed the track.
			// Each artist object includes a link in href to more detailed information about the artist.
			Artists []struct {
				// ExternalURLs is the known external URLs for this artist.
				ExternalURLs ExternalURLs `json:"external_urls"`
				// Href is a link to the Web API endpoint providing full details of the artist.
				Href string `json:"href"`
				// ID is the Spotify ID for the artist.
				ID string `json:"id"`
				// Name is the name of the artist.
				Name string `json:"name"`
				// Type is the object type.
				Type string `json:"type"`
				// URI is the Spotify URI for the artist.
				URI string `json:"uri"`
			} `json:"artists"`
			// Deprecated: AvailableMarkets is a list of the countries in which the track can be played.
			AvailableMarkets []string `json:"available_markets"`
			// DiscNumber is the disc number (usually 1 unless the album consists of more than one disc).
			DiscNumber int `json:"disc_number"`
			// DurationMs is the track length in milliseconds.
			DurationMs int `json:"duration_ms"`
			// Explicit is whether the track has explicit lyrics.
			Explicit bool `json:"explicit"`
			// ExternalIDs is the known external IDs for the track.
			ExternalIDs ExternalIDs `json:"external_ids"`
			// ExternalURLs is the known external URLs for this track.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// Href is a link to the Web API endpoint providing full details of the track.
			Href string `json:"href"`
			// ID is the Spotify ID for the track.
			ID string `json:"id"`
			// Playable is true if the track is playable in the given market, otherwise false.
			Playable bool `json:"is_playable"`
			// Deprecated: LinkedFrom is part of the response when Track Relinking is applied, and the requested track has been replaced with different track.
			LinkedFrom map[string]any `json:"linked_from"`
			// Restrictions are included in the response when a content restriction is applied.
			Restrictions *Restrictions `json:"restrictions"`
			// Name is the name of the track.
			Name string `json:"name"`
			// Deprecated: Popularity is the popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
			Popularity int `json:"popularity"`
			// Deprecated: PreviewURL is a link to a 30-second preview (MP3 format) of the track. Can be null
			PreviewURL *string `json:"preview_url"`
			// TrackNumber is the number of the track. If an album has several discs, the track number is the number on the specified disc.
			TrackNumber int `json:"track_number"`
			// Type is the object type: "track".
			Type string `json:"type"`
			// URI is the Spotify URI for the track.
			URI string `json:"uri"`
			// Local is whether the track is from a local file.
			Local bool `json:"is_local"`
		}
		if err := json.Unmarshal(raw.Item, &t); err != nil {
			return fmt.Errorf("failed to unmarshal track: %w", err)
		}
		p.Track = &t
	case "episode":
		var e struct {
			// Deprecated: AudioPreviewURL is a URL to a 30-second preview (MP3 format) of the episode. null if not available.
			AudioPreviewURL *string `json:"audio_preview_url"`
			// Description is a description of the episode. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
			Description string `json:"description"`
			// HTMLDescription is a description of the episode. This field may contain HTML tags.
			HTMLDescription string `json:"html_description"`
			// DurationMs is the episode length in milliseconds.
			DurationMs int `json:"duration_ms"`
			// Explicit is whether the episode has explicit content.
			Explicit bool `json:"explicit"`
			// ExternalURLs is the external URLs for this episode.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// Href is a link to the Web API endpoint providing full details of the episode.
			Href string `json:"href"`
			// ID is the Spotify ID for the episode.
			ID string `json:"id"`
			// Images is the cover art for the episode in various sizes, widest first.
			Images []ImageObject `json:"images"`
			// ExternallyHosted is true if the episode is hosted outside Spotify's CDN.
			ExternallyHosted bool `json:"is_externally_hosted"`
			// Playable is True if the episode is playable in the given market, otherwise false.
			Playable bool `json:"is_playable"`
			// Deprecated: Language is the language used in the episode, identified by an ISO 639 code.
			Language *string `json:"language"`
			// Languages is a list of the languages used in the episode, identified by their ISO 639-1 code.
			Languages []string `json:"languages"`
			// Name is the name of the episode.
			Name string `json:"name"`
			// ReleaseDate is the date the episode was first released.
			ReleaseDate string `json:"release_date"`
			// ReleaseDatePrecision is the precision with which release_date value is known. Allowed values are "year", "month", or "day".
			ReleaseDatePrecision string `json:"release_date_precision"`
			// ResumePoint is the user's most recent position in the episode.
			// Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
			ResumePoint struct {
				// FullyPlayed is whether the episode has been fully played by the user.
				FullyPlayed bool `json:"fully_played"`
				// ResumePositionMs is the user's most recent position in the episode in milliseconds.
				ResumePositionMs int `json:"resume_position_ms"`
			} `json:"resume_point"`
			// Type is the object type. Allowed values: "episode".
			Type string `json:"type"`
			// URI is the Spotify URI for the episode.
			URI string `json:"uri"`
			// Restrictions is included in the response when a content restriction is applied.
			Restrictions *Restrictions `json:"restrictions"`
			// Show is the show on which the episode belongs.
			Show struct {
				// Deprecated: AvailableMarkets is a list of the countries in which the show can be played, identified by their ISO 3166-1 alpha-2 code.
				AvailableMarkets []string `json:"available_markets"`
				// Copyrights is the copyright statements of the show.
				Copyrights []CopyrightObject `json:"copyrights"`
				// Description is a description of the show. HTML tags are stripped away from this field, use HTMLDescription field in case HTML tags are needed.
				Description string `json:"description"`
				// HTMLDescription is a description of the show. This field may contain HTML tags.
				HTMLDescription string `json:"html_description"`
				// Explicit is whether the show has explicit content.
				Explicit bool `json:"explicit"`
				// ExternalURLs is the external URLs for this show.
				ExternalURLs ExternalURLs `json:"external_urls"`
				// Href is a link to the Web API endpoint providing full details of the show.
				Href string `json:"href"`
				// ID is the Spotify ID for the show.
				ID string `json:"id"`
				// Images is the cover art for the show in various sizes, widest first.
				Images []ImageObject `json:"images"`
				// ExternallyHosted is true if all shows episodes are hosted outside Spotify's CDN.
				ExternallyHosted *bool `json:"is_externally_hosted"`
				// Languages is a list of the languages used in the show, identified by their ISO 639 code.
				Languages []string `json:"languages"`
				// MediaType is the media type of the show.
				MediaType string `json:"media_type"`
				// Name is the name of the show.
				Name string `json:"name"`
				// Deprecated: Publisher is the publisher of the show.
				Publisher *string `json:"publisher"`
				// Type is the object type. Allowed values is "show".
				Type string `json:"type"`
				// URI is the Spotify URI for the show.
				URI string `json:"uri"`
				// TotalEpisodes is the number of episodes in this show.
				TotalEpisodes int `json:"total_episodes"`
			} `json:"show"`
		}
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
