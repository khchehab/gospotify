package gospotify

type ChapterObject struct {
	// Deprecated: AudioPreviewURL is a URL to a 30-second preview (MP3 format) of the chapter. null if not available.
	AudioPreviewURL *string `json:"audio_preview_url"`
	// Deprecated: AvailableMarkets is a list of the countries in which the chapter can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// ChapterNumber is the number of the chapter.
	ChapterNumber int `json:"chapter_number"`
	// Description is a description of the chapter. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// HTMLDescription is a description of the chapter. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// DurationMs is the chapter length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Explicit is whether the chapter has explicit content.
	Explicit bool `json:"explicit"`
	// ExternalURLs is the external URLs for this chapter.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Href is a link to the Web API endpoint providing full details of the chapter.
	Href string `json:"href"`
	// ID is the Spotify ID for the chapter.
	ID string `json:"id"`
	// Images is the cover art for the chapter in various sizes, widest first.
	Images []ImageObject `json:"images"`
	// Playable is True if the chapter is playable in the given market, otherwise false.
	Playable bool `json:"is_playable"`
	// Languages is a list of the languages used in the chapter, identified by their ISO 639-1 code.
	Languages []string `json:"languages"`
	// Name is the name of the chapter.
	Name string `json:"name"`
	// ReleaseDate is the date the chapter was first released.
	ReleaseDate string `json:"release_date"`
	// ReleaseDatePrecision is the precision with which release_date value is known. Allowed values are "year", "month", or "day".
	ReleaseDatePrecision string `json:"release_date_precision"`
	// ResumePoint is the user's most recent position in the chapter.
	// Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
	ResumePoint struct {
		// FullyPlayed is whether the episode has been fully played by the user.
		FullyPlayed bool `json:"fully_played"`
		// ResumePositionMs is the user's most recent position in the episode in milliseconds.
		ResumePositionMs int `json:"resume_position_ms"`
	} `json:"resume_point"`
	// Type is the object type. Allowed values: "episode".
	Type string `json:"type"`
	// URI is the Spotify URI for the chapter.
	URI string `json:"uri"`
	// Restrictions is included in the response when a content restriction is applied.
	Restrictions *Restrictions `json:"restrictions"`
	Audiobook    struct {
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
	} `json:"audiobook"`
}
