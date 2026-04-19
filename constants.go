package gospotify

const (
	SpotifyBaseUrl    = "https://api.spotify.com"
	SpotifyAPIVersion = "/v1"

	////////////
	// Scopes //
	////////////

	// Images

	ScopeUgcImageUpload = "ugc-image-upload"

	// Spotify Connect

	ScopeUserReadPlaybackState    = "user-read-playback-state"
	ScopeUserModifyPlaybackState  = "user-modify-playback-state"
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"

	// Playback

	ScopeAppRemoteControl = "app-remote-control"
	ScopeStreaming        = "streaming"

	// Playlists

	ScopePlaylistReadPrivate       = "playlist-read-private"
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	ScopePlaylistModifyPrivate     = "playlist-modify-private"
	ScopePlaylistModifyPublic      = "playlist-modify-public"

	// Follow

	ScopeUserFollowModify = "user-follow-modify"
	ScopeUserFollowRead   = "user-follow-read"

	// Listening History

	ScopeUserReadPlaybackPosition = "user-read-playback-position"
	ScopeUserTopRead              = "user-top-read"
	ScopeUserReadRecentlyPlayed   = "user-read-recently-played"

	// Library

	ScopeUserLibraryModify = "user-library-modify"
	ScopeUserLibraryRead   = "user-library-read"

	// Users

	ScopeUserReadEmail    = "user-read-email"
	ScopeUserReadPrivate  = "user-read-private"
	ScopeUserPersonalized = "user-personalized"

	// Open Access

	ScopeUserSoaLink           = "user-soa-link"
	ScopeUserSoaUnlink         = "user-soa-unlink"
	ScopeSoaManageEntitlements = "soa-manage-entitlements"
	ScopeSoaManagePartner      = "soa-manage-partner"
	ScopeSoaCreatePartner      = "soa-create-partner"
)
