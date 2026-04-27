package gospotify

type SearchResult struct {
	// Tracks is a page of the track result.
	Tracks Page[struct {
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
	}] `json:"tracks"`
	// Artists is a page of the artist result.
	Artists Page[struct {
		// ExternalURLs is the known external URLs for this artist.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Deprecated: Followers is the information about the followers of the artist.
		Followers FollowersObject `json:"followers"`
		// Deprecated: Genres is a list of the genres the artist is associated with. If not yet classified, the array is empty.
		Genres []string `json:"genres"`
		// Href is a link to the Web API endpoint providing full details of the artist.
		Href string `json:"href"`
		// ID is the Spotify ID for the artist.
		ID string `json:"id"`
		// Images is the images of the artist in various sizes, widest first.
		Images []ImageObject `json:"images"`
		// Name is the name of the artist.
		Name string `json:"name"`
		// Deprecated: Popularity is the popularity of the artist. The value will be between 0 and 100, with 100 being the most popular.
		// The artist's popularity is calculated from the popularity of all the artist's tracks.
		Popularity int `json:"popularity"`
		// Type is the object type.
		Type string `json:"type"`
		// URI is the Spotify URI for the artist.
		URI string `json:"uri"`
	}] `json:"artists"`
	// Albums is a page of the album result.
	Albums Page[struct {
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
	}] `json:"albums"`
	// Playlists is a page of the playlist result.
	Playlists Page[struct {
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
	}] `json:"playlists"`
	// Shows is a page of the show result.
	Shows Page[struct {
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
	}] `json:"shows"`
	// Episodes is a page of the episode result.
	Episodes Page[struct {
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
	}] `json:"episodes"`
	// Audiobooks is a page of the audiobook result.
	Audiobooks Page[struct {
		// Authors is the author(s) of the audiobook.
		Authors []AuthorObject `json:"authors"`
		// Deprecated: AvailableMarkets is the list of the countries in which the audiobook can be played, identified by their ISO 3166-1 alpha-2 code.
		AvailableMarkets []string `json:"available_markets"`
		// Copyrights is the copyright statements of the audiobook.
		Copyrights []CopyrightObject `json:"copyrights"`
		// Description is a description of the audiobook. HTML tags are stripped away from this field, use HTMLDescription field in case HTML tags are needed.
		Description string `json:"description"`
		// HTMLDescription is a description of the audiobook. This field may contain HTML tags.
		HTMLDescription string `json:"html_description"`
		// Edition is the edition of the audiobook.
		Edition string `json:"edition"`
		// Explicit is whether the audiobook has explicit content.
		Explicit bool `json:"explicit"`
		// ExternalURLs is the external URLs for this audiobook.
		ExternalURLs ExternalURLs `json:"external_urls"`
		// Href is a link to the Web API endpoint providing full details of the audiobook.
		Href string `json:"href"`
		// ID is the Spotify ID for the audiobook.
		ID string `json:"id"`
		// Images is the cover art for the audiobook in various sizes, widest first.
		Images []ImageObject `json:"images"`
		// Languages is a list of the languages used in the audiobook, identified by their ISO 639 code.
		Languages []string `json:"languages"`
		// MediaType is the media type of the audiobook.
		MediaType string `json:"media_type"`
		// Name is the name of the audiobook.
		Name string `json:"name"`
		// Narrators is the narrator(s) for the audiobook.
		Narrators []NarratorObject `json:"narrators"`
		// Deprecated: Publisher is the publisher of the audiobook.
		Publisher *string `json:"publisher"`
		// Type is the object type. Allowed values is "audiobook".
		Type string `json:"type"`
		// URI is the Spotify URI for the audiobook.
		URI string `json:"uri"`
		// TotalChapters is the number of chapters in this audiobook.
		TotalChapters int `json:"total_chapters"`
	}] `json:"audiobooks"`
}
