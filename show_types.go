package gospotify

type ShowObject struct {
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
	// Episodes is the episodes of the show.
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
}

type SavedShowObject struct {
	// AddedAt is the date and time the show was saved.
	AddedAt string `json:"added_at"`
	// Show is the information about the show.
	Show Page[struct {
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
	}] `json:"show"`
}

type ShowEpisode struct {
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
}
