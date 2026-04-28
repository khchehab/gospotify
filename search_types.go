package gospotify

type SearchResult struct {
	// Tracks is a page of the track result.
	Tracks Page[TrackObject] `json:"tracks"`
	// Artists is a page of the artist result.
	Artists Page[ArtistObject] `json:"artists"`
	// Albums is a page of the album result.
	Albums Page[SimplifiedAlbumObject] `json:"albums"`
	// Playlists is a page of the playlist result.
	Playlists Page[SimplifiedPlaylistObject] `json:"playlists"`
	// Shows is a page of the show result.
	Shows Page[SimplifiedShowObject] `json:"shows"`
	// Episodes is a page of the episode result.
	Episodes Page[SimplifiedEpisodeObject] `json:"episodes"`
	// Audiobooks is a page of the audiobook result.
	Audiobooks Page[SimplifiedAudiobookObject] `json:"audiobooks"`
}
