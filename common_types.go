package gospotify

// TimeRange is the time frame the affinities are computed.
type TimeRange string

const (
	// ShortTerm is approximately last 4 weeks.
	ShortTerm TimeRange = "short_term"
	// MediumTerm is approximately last 6 months.
	MediumTerm TimeRange = "medium_term"
	// LongTerm is calculated from ~1 year of data and including all new data as it becomes available.
	LongTerm TimeRange = "long_term"
)

// Page is a page of data.
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

// ExternalURLs is the known external URLs for this object.
type ExternalURLs struct {
	// Spotify is the Spotify URL for the object.
	Spotify string `json:"spotify"`
}

// ExternalIDs is the known external IDs for the track.
type ExternalIDs struct {
	// ISRC is the International Standard Recording Code.
	ISRC string `json:"isrc"`
	// EAN is the International Article Number.
	EAN string `json:"ean"`
	// UP Cis the Universal Product Code.
	UPC string `json:"upc"`
}

// Restrictions is the object returned by the Spotify API for any restrictions.
type Restrictions struct {
	// Reason is the reason for the restriction. Supported values: "market", "product", or "explicit".
	Reason string `json:"reason"`
}

// ImageObject is the user's profile image.
type ImageObject struct {
	// URL is the source URL of the image.
	URL string `json:"url"`
	// Height is the image height in pixels.
	Height *int `json:"height"`
	// Width is the image width in pixels.
	Width *int `json:"width"`
}

// FollowersObject is the information about the followers of the user.
type FollowersObject struct {
	// Href will always be set to null, as the Web API does not support it at the moment.
	Href *string `json:"href"`
	// Total is the total number of followers.
	Total int `json:"total"`
}

// CursorsObject is the cursors used to find the next set of items.
type CursorsObject struct {
	// After is the cursor to use as a key to find the next page of items.
	After *string `json:"after"`
	// Before is the cursor to use as key to find the previous page of items.
	Before *string `json:"before"`
}
