package gospotify

type SimplifiedShowObject struct {
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
	ExternalURLs ExternalURLsObject `json:"external_urls"`
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
}

type ShowObject struct {
	SimplifiedShowObject
	// Episodes is the episodes of the show.
	Episodes Page[SimplifiedEpisodeObject] `json:"episodes"`
}

type SavedShowObject struct {
	// AddedAt is the date and time the show was saved.
	AddedAt string `json:"added_at"`
	// Show is the information about the show.
	Show SimplifiedShowObject `json:"show"`
}
