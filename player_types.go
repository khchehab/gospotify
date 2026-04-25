package gospotify

import (
	"encoding/json"
	"fmt"
)

type PlaybackObject struct {
	// Device is the device that is currently active.
	Device DeviceObject `json:"device"`
	// RepeatState is the repeat state. Possible values are "off", "track", "context".
	RepeatState string `json:"repeat_state"`
	// ShuffleState is whether shuffle is on or off.
	ShuffleState bool `json:"shuffle_state"`
	// Context is the context the playback. Can be null.
	Context *struct {
		// Type is the object type, e.g. "artist", "playlist", "album", "show".
		Type string `json:"type"`
		// Href is a link to the Web API endpoint providing full details of the track.
		Href string `json:"href"`
		// ExternalURLs for this context.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// URI is the Spotify URI for the context.
		URI string `json:"uri"`
	} `json:"context"`
	// Timestamp is the unix millisecond timestamp when the playback state was last changed (play, pause, skip, scrub, new song, etc.).
	Timestamp int `json:"timestamp"`
	// ProgressMs is the progress into the currently playing track or episode. Can be null.
	ProgressMs *int `json:"progress_ms"`
	// Playing is if something is currently playing, return true.
	Playing bool `json:"is_playing"`
	// Item is the currently playing track or episode. Can be null. TODO update documentation
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
	// CurrentlyPlayingType is the object type of the currently playing item.
	// Can be one of "track", "episode", "ad" or "unknown".
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// Actions is the list of actions allowed to update the user interface based on which playback actions are available within the current context.
	Actions struct {
		// InterruptingPlayback is for interrupting playback. Optional field.
		InterruptingPlayback bool `json:"interrupting_playback,omitempty"`
		// Pausing is for pausing. Optional field.
		Pausing bool `json:"pausing,omitempty"`
		// Resuming is for resuming. Optional field.
		Resuming bool `json:"resuming,omitempty"`
		// Seeking is for seeking playback location. Optional field.
		Seeking bool `json:"seeking,omitempty"`
		// SkippingNext is for skipping to the next context. Optional field.
		SkippingNext bool `json:"skipping_next,omitempty"`
		// SkippingPrevious is for skipping to the previous context. Optional field.
		SkippingPrevious bool `json:"skipping_prev,omitempty"`
		// TogglingRepeatContext is for toggling the repeat context flag. Optional field.
		TogglingRepeatContext bool `json:"toggling_repeat_context,omitempty"`
		// TogglingShuffle is for toggling shuffle flag. Optional field.
		TogglingShuffle bool `json:"toggling_shuffle,omitempty"`
		// TogglingRepeatTrack is for toggling the repeat track flag. Optional field.
		TogglingRepeatTrack bool `json:"toggling_repeat_track,omitempty"`
		// TransferringPlayback is for transferring playback between devices. Optional field.
		TransferringPlayback bool `json:"transferring_playback,omitempty"`
	} `json:"actions"`
}

func (p *PlaybackObject) UnmarshalJSON(data []byte) error {
	var raw struct {
		// Device is the device that is currently active.
		Device DeviceObject `json:"device"`
		// RepeatState is the repeat state. Possible values are "off", "track", "context".
		RepeatState string `json:"repeat_state"`
		// ShuffleState is whether shuffle is on or off.
		ShuffleState bool `json:"shuffle_state"`
		// Context is the context the playback. Can be null.
		Context *struct {
			// Type is the object type, e.g. "artist", "playlist", "album", "show".
			Type string `json:"type"`
			// Href is a link to the Web API endpoint providing full details of the track.
			Href string `json:"href"`
			// ExternalURLs for this context.
			ExternalURLs ExternalURLs `json:"external_urls"`
			// URI is the Spotify URI for the context.
			URI string `json:"uri"`
		} `json:"context"`
		// Timestamp is the unix millisecond timestamp when the playback state was last changed (play, pause, skip, scrub, new song, etc.).
		Timestamp int `json:"timestamp"`
		// ProgressMs is the progress into the currently playing track or episode. Can be null.
		ProgressMs *int `json:"progress_ms"`
		// Playing is if something is currently playing, return true.
		Playing bool `json:"is_playing"`
		// Item is the currently playing track or episode. Can be null. TODO update documentation
		Item json.RawMessage `json:"item"`
		// CurrentlyPlayingType is the object type of the currently playing item.
		// Can be one of "track", "episode", "ad" or "unknown".
		CurrentlyPlayingType string `json:"currently_playing_type"`
		// Actions is the list of actions allowed to update the user interface based on which playback actions are available within the current context.
		Actions struct {
			// InterruptingPlayback is for interrupting playback. Optional field.
			InterruptingPlayback bool `json:"interrupting_playback,omitempty"`
			// Pausing is for pausing. Optional field.
			Pausing bool `json:"pausing,omitempty"`
			// Resuming is for resuming. Optional field.
			Resuming bool `json:"resuming,omitempty"`
			// Seeking is for seeking playback location. Optional field.
			Seeking bool `json:"seeking,omitempty"`
			// SkippingNext is for skipping to the next context. Optional field.
			SkippingNext bool `json:"skipping_next,omitempty"`
			// SkippingPrevious is for skipping to the previous context. Optional field.
			SkippingPrevious bool `json:"skipping_prev,omitempty"`
			// TogglingRepeatContext is for toggling the repeat context flag. Optional field.
			TogglingRepeatContext bool `json:"toggling_repeat_context,omitempty"`
			// TogglingShuffle is for toggling shuffle flag. Optional field.
			TogglingShuffle bool `json:"toggling_shuffle,omitempty"`
			// TogglingRepeatTrack is for toggling the repeat track flag. Optional field.
			TogglingRepeatTrack bool `json:"toggling_repeat_track,omitempty"`
			// TransferringPlayback is for transferring playback between devices. Optional field.
			TransferringPlayback bool `json:"transferring_playback,omitempty"`
		} `json:"actions"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.Device = raw.Device
	p.RepeatState = raw.RepeatState
	p.ShuffleState = raw.ShuffleState
	p.Context = raw.Context
	p.Timestamp = raw.Timestamp
	p.ProgressMs = raw.ProgressMs
	p.Playing = raw.Playing
	p.CurrentlyPlayingType = raw.CurrentlyPlayingType
	p.Actions = raw.Actions

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

type DeviceObject struct {
	// ID is the device ID. This ID is unique and persistent to some extent.
	// However, this is not guaranteed, and any cached "device_id" should periodically be cleared out and refetched as necessary.
	ID *string `json:"id"`
	// Active if this device is the currently active device.
	Active bool `json:"is_active"`
	// PrivateSession if this device is currently in a private session.
	PrivateSession bool `json:"is_private_session"`
	// Restricted is whether controlling this device is restricted.
	// At present, if this is true, then no Web API commands will be accepted by this device.
	Restricted bool `json:"is_restricted"`
	// Name is a human-readable name for the device.
	Name string `json:"name"`
	// Type is the device type, such as "computer", "smartphone" or "speaker".
	Type string `json:"type"`
	// VolumePercent is the current volume in percent.
	VolumePercent *int `json:"volume_percent"`
	// SupportsVolume if this device can be used to set the volume.
	SupportsVolume bool `json:"supports_volume"`
}

type QueueItem struct {
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

func (q *QueueItem) UnmarshalJSON(data []byte) error {
	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
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
		if err := json.Unmarshal(data, &t); err != nil {
			return fmt.Errorf("failed to unmarshal track: %w", err)
		}
		q.Track = &t
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
		if err := json.Unmarshal(data, &e); err != nil {
			return fmt.Errorf("failed to unmarshal episode: %w", err)
		}
		q.Episode = &e
	default:
		return fmt.Errorf("unknown item type: %s", peek.Type)
	}

	return nil
}

type UserQueue struct {
	// CurrentlyPlaying is the currently playing track or episode. Can be null.
	CurrentlyPlaying *QueueItem `json:"currently_playing"`
	// Queue is the tracks or episodes in the queue. Can be empty.
	Queue []QueueItem `json:"queue"`
}

type PlayHistoryObject struct {
}

type TransferPlaybackRequest struct {
	// DeviceIDs is an array containing the ID of the device on which playback should be started/transferred.
	// Note: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return 400 Bad Request.
	DeviceIDs []string `json:"device_ids"`
	// Play if true, ensure playback happens on a new device. If false or not provided, keep the current playback state.
	Play bool `json:"play"`
}

type StartResumePlaybackRequest struct {
	// ContextURI is a Spotify URI of the context to play. Valid contexts are albums, artists & playlists.
	ContextURI *string `json:"context_uri,omitempty"`
	// URIs is an array of the Spotify track URIs to play.
	URIs []string `json:"uris,omitempty"`
	// Offset indicates from where in the context playback should start.
	// Only available when context_uri corresponds to an album or playlist object "position" is zero-based and can’t be negative.
	Offset map[string]any `json:"offset,omitempty"`
	// PositionMs is the position of the playback (in milliseconds).
	PositionMs *int `json:"position_ms,omitempty"`
}
