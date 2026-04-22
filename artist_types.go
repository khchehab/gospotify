package gospotify

type ArtistObject struct {
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

type ArtistAlbum struct {
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
	// Deprecated: AlbumGroup is a field that describes the relationship between the artist and the album
	AlbumGroup *string `json:"album_group"`
}
