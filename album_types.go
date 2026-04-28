package gospotify

type SimplifiedAlbumObject struct {
	// AlbumType is the type of the album. Allowed values are "album", "single", or "compilation".
	AlbumType string `json:"album_type"`
	// TotalTracks is the number of tracks in the album.
	TotalTracks int `json:"total_tracks"`
	// Deprecated: AvailableMarkets is the markets in which the album is available: ISO 3166-1 alpha-2 country codes.
	AvailableMarkets []string `json:"available_markets"`
	// ExternalURLs is the known external URLs for this album.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
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
	Restrictions *RestrictionsObject `json:"restrictions"`
	// Type is the object type. Allowed values is "album".
	Type string `json:"type"`
	// URI is the Spotify URI for the album.
	URI string `json:"uri"`
	// Artists are the artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	Artists []SimplifiedArtistObject `json:"artists"`
}

type AlbumObject struct {
	SimplifiedAlbumObject
	Tracks Page[SimplifiedTrackObject] `json:"tracks"`
	// Copyrights is the copyright statements of the album.
	Copyrights []CopyrightObject `json:"copyrights"`
	// ExternalIDs is the known external IDs for the track.
	ExternalIDs ExternalIDsObject `json:"external_ids"`
	// Deprecated: Genres the array is always empty.
	Genres []string `json:"genres"`
	// Deprecated: Label is the label associated with the album.
	Label *string `json:"label"`
	// Deprecated: Popularity is the popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
	Popularity int `json:"popularity"`
}

type ArtistDiscographyAlbumObject struct {
	SimplifiedAlbumObject
	// Deprecated: AlbumGroup is a field that describes the relationship between the artist and the album
	AlbumGroup *string `json:"album_group"`
}

type SavedAlbumObject struct {
	// AddedAt is the date and time the album was saved.
	AddedAt string `json:"added_at"`
	// Track is the information about the track.
	Album AlbumObject `json:"album"`
}
