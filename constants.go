package gospotify

import (
	"time"
)

const (
	spotifyBaseUrl    = "https://api.spotify.com"
	spotifyAPIVersion = "/v1"

	authStateLength           = 16
	authWaitTimeout           = 2 * time.Minute
	authCodeExchangeTimeout   = 10 * time.Second
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
	RepeatOff     RepeatState = "off"
	RepeatTrack   RepeatState = "track"
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
