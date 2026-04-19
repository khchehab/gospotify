package gospotify

// ArtistObject is the object returned by the Spotify API for an artist.
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
	// Deprecated: ßPopularity is the popularity of the artist. The value will be between 0 and 100, with 100 being the most popular.
	// The artist's popularity is calculated from the popularity of all the artist's tracks.
	Popularity int `json:"popularity"`
	// Type is the object type.
	Type string `json:"type"`
	// URI is the Spotify URI for the artist.
	URI string `json:"uri"`
}

// SimplifiedArtistObject is the simplified object returned by the Spotify API for an artist.
type SimplifiedArtistObject struct {
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
}
