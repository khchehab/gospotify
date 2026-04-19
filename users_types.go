package gospotify

import (
	"net/url"
	"strconv"
)

// topItemsParameters is the parameters for the Get a User's Top Items endpoint.
type topItemsParameters struct {
	// timeRange is over what time frame the affinities are computed.
	timeRange TimeRange
	// limit is the maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	limit int
	// offset is the index of the first item to return. Default: 0 (the first item).
	// Use with limit to get the next set of items.
	offset int
}

// toQuery converts the parameters to a query string.
func (p topItemsParameters) toQuery() string {
	query := url.Values{}
	query.Set("time_range", string(p.timeRange))
	query.Set("limit", strconv.Itoa(p.limit))
	query.Set("offset", strconv.Itoa(p.offset))
	return query.Encode()
}

// TopItemsOption is an option for the Get a User's Top Items endpoint.
type TopItemsOption func(*topItemsParameters)

// WithTimeRange sets the time range for the affinities.
func WithTimeRange(timeRange TimeRange) TopItemsOption {
	return func(p *topItemsParameters) {
		p.timeRange = timeRange
	}
}

// WithLimit sets the maximum number of items to return.
func WithLimit(limit int) TopItemsOption {
	return func(p *topItemsParameters) {
		if limit < 1 {
			limit = 1
		}
		if limit > 50 {
			limit = 50
		}
		p.limit = limit
	}
}

// WithOffset sets the index of the first item to return.
func WithOffset(offset int) TopItemsOption {
	return func(p *topItemsParameters) {
		if offset < 0 {
			offset = 0
		}
		p.offset = offset
	}
}

// UserProfile is the user's profile.
type UserProfile struct {
	// Deprecated: Country is the country of the user, as set in the user's account profile.
	Country *string `json:"country"`
	// DisplayName is the name displayed on the user's profile. null if not available.
	DisplayName *string `json:"display_name"`
	// Deprecated: Email is the user's email address, as entered by the user when creating their account.
	Email *string `json:"email"`
	// Deprecated: ExplicitContent is the user's explicit content settings.
	ExplicitContent *struct {
		// FilterEnabled when true, indicates that explicit content should not be played.
		FilterEnabled bool `json:"filter_enabled"`
		// FilterLocked when true, indicates that the explicit content setting is locked and can't be changed by the user.
		FilterLocked bool `json:"filter_locked"`
	} `json:"explicit_content"`
	// ExternalURLs is the known external URLs for this user.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Deprecated: Followers is the information about the followers of the user.
	Followers FollowersObject `json:"followers"`
	// Href is a link to the Web API endpoint for this user.
	Href string `json:"href"`
	// ID is the Spotify user ID for the user.
	ID string `json:"id"`
	// Images is the user's profile image.
	Images []ImageObject `json:"images"`
	// Deprecated: Product is the user's Spotify subscription level: "premium", "free", etc. (The subscription level "open" can be considered the same as "free".).
	Product *string `json:"product"`
	// Type is the object type: "user".
	Type string `json:"type"`
	// URI is the Spotify URI for the user.
	URI string `json:"uri"`
}
