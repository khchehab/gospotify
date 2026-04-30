package gospotify

import (
	"time"
)

const (
	// spotifyBaseUrl is the base URL for the Spotify Web API.
	spotifyBaseUrl = "https://api.spotify.com"
	// spotifyAPIVersion is the API version path segment appended to the base URL.
	spotifyAPIVersion = "/v1"

	// authStateLength is the length of the random CSRF state string generated for each OAuth flow.
	authStateLength = 16
	// authWaitTimeout is how long the authorization code flow will wait for the user to complete
	// the browser-based login before giving up.
	authWaitTimeout = 2 * time.Minute
	// authCodeExchangeTimeout is the HTTP timeout for the token exchange request after the
	// authorization code is received.
	authCodeExchangeTimeout = 10 * time.Second
	// authServerShutdownTimeout is how long the local callback server is given to shut down
	// gracefully after the authorization flow completes.
	authServerShutdownTimeout = 5 * time.Second
)

const (
	// ShortTerm is approximately last 4 weeks.
	ShortTerm TimeRange = "short_term"
	// MediumTerm is approximately last 6 months.
	MediumTerm TimeRange = "medium_term"
	// LongTerm is calculated from ~1 year of data and including all new data as it becomes available.
	LongTerm TimeRange = "long_term"
)

const (
	// RepeatOff disables repeat — playback stops at the end of the current context.
	RepeatOff RepeatState = "off"
	// RepeatTrack repeats the current track indefinitely.
	RepeatTrack RepeatState = "track"
	// RepeatContext repeats the current context (album, playlist, etc.) indefinitely.
	RepeatContext RepeatState = "context"
)

const (
	ItemTypeAlbum     ItemType = "album"
	ItemTypeArtist    ItemType = "artist"
	ItemTypePlaylist  ItemType = "playlist"
	ItemTypeTrack     ItemType = "track"
	ItemTypeShow      ItemType = "show"
	ItemTypeEpisode   ItemType = "episode"
	ItemTypeAudiobook ItemType = "audiobook"
)

// itemTypeMap is the set of valid ItemType values, used by [ItemType.Valid].
var itemTypeMap = map[ItemType]struct{}{
	ItemTypeAlbum:     {},
	ItemTypeArtist:    {},
	ItemTypePlaylist:  {},
	ItemTypeTrack:     {},
	ItemTypeShow:      {},
	ItemTypeEpisode:   {},
	ItemTypeAudiobook: {},
}

const (
	// Images

	ScopeUgcImageUpload Scope = "ugc-image-upload"

	// Spotify Connect

	ScopeUserReadPlaybackState    Scope = "user-read-playback-state"
	ScopeUserModifyPlaybackState  Scope = "user-modify-playback-state"
	ScopeUserReadCurrentlyPlaying Scope = "user-read-currently-playing"

	// Playback

	ScopeAppRemoteControl Scope = "app-remote-control"
	ScopeStreaming        Scope = "streaming"

	// Playlists

	ScopePlaylistReadPrivate       Scope = "playlist-read-private"
	ScopePlaylistReadCollaborative Scope = "playlist-read-collaborative"
	ScopePlaylistModifyPrivate     Scope = "playlist-modify-private"
	ScopePlaylistModifyPublic      Scope = "playlist-modify-public"

	// Follow

	ScopeUserFollowModify Scope = "user-follow-modify"
	ScopeUserFollowRead   Scope = "user-follow-read"

	// Listening History

	ScopeUserReadPlaybackPosition Scope = "user-read-playback-position"
	ScopeUserTopRead              Scope = "user-top-read"
	ScopeUserReadRecentlyPlayed   Scope = "user-read-recently-played"

	// Library

	ScopeUserLibraryModify Scope = "user-library-modify"
	ScopeUserLibraryRead   Scope = "user-library-read"

	// Users

	ScopeUserReadEmail    Scope = "user-read-email"
	ScopeUserReadPrivate  Scope = "user-read-private"
	ScopeUserPersonalized Scope = "user-personalized"

	// Open Access

	ScopeUserSoaLink           Scope = "user-soa-link"
	ScopeUserSoaUnlink         Scope = "user-soa-unlink"
	ScopeSoaManageEntitlements Scope = "soa-manage-entitlements"
	ScopeSoaManagePartner      Scope = "soa-manage-partner"
	ScopeSoaCreatePartner      Scope = "soa-create-partner"
)
