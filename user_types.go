package gospotify

type CurrentUserProfile struct {
	// Deprecated: Country is the country of the user, as set in the user's account profile.
	Country *string `json:"country"`
	// DisplayName is the name displayed on the user's profile. null if not available.
	DisplayName *string `json:"display_name"`
	// Deprecated: Email is the user's email address, as entered by the user when creating their account.
	Email *string `json:"email"`
	// Deprecated: ExplicitContent is the user's explicit content settings.
	ExplicitContent *ExplicitContent `json:"explicit_content"`
	// ExternalURLs is the known external URLs for this user.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Deprecated: Followers is the information about the followers of the user.
	Followers FollowersObject `json:"followers"`
	// Href is a link to the Web API endpoint for this user.
	Href string `json:"href"`
	// ID is the Spotify user ID for the user.
	ID string `json:"id"`
	// Images is the user's profile image.
	Images []ImageObject `json:"images"`
	// Deprecated: Product is the user's Spotify subscription level: "premium", "free", etc. (The subscription level "open" can be considered the same as "free".).
	Product *string `json:"product"`
	// Type is the object type: "user".
	Type string `json:"type"`
	// URI is the Spotify URI for the user.
	URI string `json:"uri"`
}

type UserTopArtists struct {
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
}

type UserTopTracks struct {
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

type FollowedArtists struct {
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
}
