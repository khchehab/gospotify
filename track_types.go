package gospotify

// TrackObject is the object returned by the Spotify API for a track.
type TrackObject struct {
	// Album is the album on which the track appears.
	// The album object includes a link in href to full information about the album.
	Album AlbumObject `json:"album"`
	// Artists are the artists who performed the track.
	// Each artist object includes a link in href to more detailed information about the artist.
	Artists []SimplifiedArtistObject `json:"artists"`
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
