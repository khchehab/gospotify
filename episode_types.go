package gospotify

type SimplifiedEpisodeObject struct {
	// Deprecated: AudioPreviewURL is a URL to a 30-second preview (MP3 format) of the episode. null if not available.
	AudioPreviewURL *string `json:"audio_preview_url"`
	// Description is a description of the episode. HTML tags are stripped away from this field, use HTMLDescription field in case HTML tags are needed.
	Description string `json:"description"`
	// HTMLDescription is a description of the episode. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// DurationMs is the episode length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Explicit is whether the episode has explicit content.
	Explicit bool `json:"explicit"`
	// ExternalURLs is the external URLs for this episode.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
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
	ResumePoint ResumePointObject `json:"resume_point"`
	// Type is the object type. Allowed values: "episode".
	Type string `json:"type"`
	// URI is the Spotify URI for the episode.
	URI string `json:"uri"`
	// Restrictions is included in the response when a content restriction is applied.
	Restrictions *RestrictionsObject `json:"restrictions"`
}

type EpisodeObject struct {
	SimplifiedEpisodeObject
	// Show is the show on which the episode belongs.
	Show SimplifiedShowObject `json:"show"`
}

type SavedEpisodeObject struct {
	// AddedAt is the date and time the episode was saved.
	AddedAt string `json:"added_at"`
	// Episode is the information about the episode.
	Episode EpisodeObject `json:"episode"`
}
