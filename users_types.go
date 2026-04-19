package gospotify

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

// followedArtistsResponse is the response for the followed artists' endpoint.
type followedArtistsResponse struct {
	// Artists is a paged set of followed by the user.
	Artists Cursor[ArtistObject] `json:"artists"`
}
