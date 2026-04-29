package gospotify

// SimplifiedAudiobookObject is a reduced representation of an audiobook, returned when audiobooks appear nested inside other objects (e.g. inside a chapter).
type SimplifiedAudiobookObject struct {
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
	ExternalURLs ExternalURLsObject `json:"external_urls"`
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
}

// AudiobookObject is the full representation of a Spotify audiobook, including its paginated chapters.
type AudiobookObject struct {
	SimplifiedAudiobookObject
	// Chapters is the chapters of the audiobook.
	Chapters Page[SimplifiedChapterObject] `json:"chapters"`
}
