package gospotify

// Scope represents a scope data type.
type Scope string

// TimeRange is the time frame the affinities are computed.
type TimeRange string

// Valid checks if the time range is valid based on a pre-defined allowed values.
func (t TimeRange) Valid() bool {
	switch t {
	case LongTerm, MediumTerm, ShortTerm:
		return true
	}
	return false
}

// RepeatState represents a repeat state data type.
type RepeatState string

// Valid checks if the repeat state is valid based on a pre-defined allowed values.
func (s RepeatState) Valid() bool {
	switch s {
	case RepeatOff, RepeatTrack, RepeatContext:
		return true
	}
	return false
}

// ItemType represents an item type data type.
type ItemType string

// Valid checks if the item type is valid based on a pre-defined allowed values.
func (t ItemType) Valid() bool {
	_, ok := itemTypeMap[t]
	return ok
}

// Page is a generic paginated response returned by Spotify list endpoints.
type Page[T any] struct {
	// Href is a link to the Web API endpoint returning the full result of the request.
	Href string `json:"href"`
	// Offset is the maximum number of items in the response (as set in the query or by default).
	Offset int `json:"offset"`
	// Limit is the maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// Next is the URL to the next page of items (null if none).
	Next *string `json:"next"`
	// Previous is the URL to the previous page of items (null if none).
	Previous *string `json:"previous"`
	// Total is the total number of items available to return.
	Total int `json:"total"`
	// Items is the list of paged data.
	Items []T `json:"items"`
}

// Cursor is a generic cursor-based paginated response used by Spotify endpoints that support cursor pagination.
type Cursor[T any] struct {
	// Href is a link to the Web API endpoint returning the full result of the request.
	Href string `json:"href"`
	// Limit is the maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// Next is the URL to the next page of items (null if none).
	Next *string `json:"next"`
	// Cursors is the cursors used to find the next set of items.
	Cursors CursorsObject `json:"cursors"`
	// Total is the total number of items available to return.
	Total int `json:"total"`
	// Items is the list of paged data.
	Items []T `json:"items"`
}

// CursorsObject holds the before and after cursor values used for cursor-based pagination.
type CursorsObject struct {
	// After is the cursor to use as a key to find the next page of items.
	After *string `json:"after"`
	// Before is the cursor to use as a key to find the previous page of items.
	Before *string `json:"before"`
}

// ExplicitContent holds the user's explicit content filter settings.
type ExplicitContent struct {
	// FilterEnabled when true, indicates that explicit content should not be played.
	FilterEnabled bool `json:"filter_enabled"`
	// FilterLocked when true, indicates that the explicit content setting is locked and can't be changed by the user.
	FilterLocked bool `json:"filter_locked"`
}

// ExternalURLsObject holds external URLs associated with a Spotify object.
type ExternalURLsObject struct {
	// Spotify is the Spotify URL for the object.
	Spotify string `json:"spotify"`
}

// FollowersObject holds follower information for a Spotify artist or user.
type FollowersObject struct {
	// Href will always be set to null, as the Web API does not support it at the moment.
	Href *string `json:"href"`
	// Total is the total number of followers.
	Total int `json:"total"`
}

// ImageObject represents a Spotify cover image in a specific size.
type ImageObject struct {
	// URL is the source URL of the image.
	URL string `json:"url"`
	// Height is the image height in pixels.
	Height *int `json:"height"`
	// Width is the image width in pixels.
	Width *int `json:"width"`
}

// RestrictionsObject describes why content is restricted from playback.
type RestrictionsObject struct {
	// Reason is the reason for the restriction. Supported values: "market", "product", or "explicit".
	Reason string `json:"reason"`
}

// ExternalIDsObject holds known external identifiers for a track.
type ExternalIDsObject struct {
	// ISRC is the International Standard Recording Code.
	ISRC string `json:"isrc"`
	// EAN is the International Article Number.
	EAN string `json:"ean"`
	// UPC is the Universal Product Code.
	UPC string `json:"upc"`
}

// CopyrightObject holds a copyright statement for an album, audiobook, or show.
type CopyrightObject struct {
	// Text is the copyright text for this content.
	Text string `json:"text"`
	// Type is the type of copyright: C = the copyright, P = the sound recording (performance) copyright.
	Type string `json:"type"`
}

// AuthorObject holds the name of an audiobook author.
type AuthorObject struct {
	// Name is the name of the author.
	Name string `json:"name"`
}

// NarratorObject holds the name of an audiobook narrator.
type NarratorObject struct {
	// Name is the name of the Narrator.
	Name string `json:"name"`
}

// ContextObject describes the playback context (e.g. an album, artist, or playlist) for the currently playing item.
type ContextObject struct {
	// Type is the object type, e.g. "artist", "playlist", "album", "show".
	Type string `json:"type"`
	// Href is a link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// ExternalURLs for this context.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
	// URI is the Spotify URI for the context.
	URI string `json:"uri"`
}

// ResumePointObject holds the user's most recent listening position within an episode or chapter.
type ResumePointObject struct {
	// FullyPlayed is whether the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`
	// ResumePositionMs is the user's most recent position in the episode in milliseconds.
	ResumePositionMs int `json:"resume_position_ms"`
}
