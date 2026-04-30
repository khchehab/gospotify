package gospotify

// CurrentUserProfile contains detailed profile information about the currently authenticated user.
type CurrentUserProfile struct {
	// Deprecated: Country is the country of the user, as set in the user's account profile.
	Country *string `json:"country"`
	// DisplayName is the name displayed on the user's profile. null if not available.
	DisplayName *string `json:"display_name"`
	// Deprecated: Email is the user's email address, as entered by the user when creating their account.
	Email *string `json:"email"`
	// Deprecated: ExplicitContent is the user's explicit content settings.
	ExplicitContent *ExplicitContent `json:"explicit_content"`
	// ExternalURLs is the known external URLs for this user.
	ExternalURLs ExternalURLsObject `json:"external_urls"`
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
