package gospotify

type AlbumObject struct {
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
	// ReleaseDate is the date the album was first released. It follows the format of "yyyy-mm".
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
	Tracks Page[struct {
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
	// Copyrights is the copyright statements of the album.
	Copyrights []CopyrightObject `json:"copyrights"`
	// ExternalIDs is the known external IDs for the track.
	ExternalIDs ExternalIDs `json:"external_ids"`
	// Deprecated: Genres the array is always empty.
	Genres []string `json:"genres"`
	// Deprecated: Label is the label associated with the album.
	Label *string `json:"label"`
	// Deprecated: Popularity is the popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
	Popularity int `json:"popularity"`
}

type SavedAlbumObject struct {
	// AddedAt is the date and time the album was saved.
	AddedAt string `json:"added_at"`
	// Track is the information about the track.
	Album AlbumObject `json:"album"`
}
type AlbumTrack struct {
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
