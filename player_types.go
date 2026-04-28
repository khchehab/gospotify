package gospotify

import (
	"encoding/json"
	"fmt"
)

type ActionsObject struct {
	// InterruptingPlayback is for interrupting playback. Optional field.
	InterruptingPlayback bool `json:"interrupting_playback,omitempty"`
	// Pausing is for pausing. Optional field.
	Pausing bool `json:"pausing,omitempty"`
	// Resuming is for resuming. Optional field.
	Resuming bool `json:"resuming,omitempty"`
	// Seeking is for seeking playback location. Optional field.
	Seeking bool `json:"seeking,omitempty"`
	// SkippingNext is for skipping to the next context. Optional field.
	SkippingNext bool `json:"skipping_next,omitempty"`
	// SkippingPrevious is for skipping to the previous context. Optional field.
	SkippingPrevious bool `json:"skipping_prev,omitempty"`
	// TogglingRepeatContext is for toggling the repeat context flag. Optional field.
	TogglingRepeatContext bool `json:"toggling_repeat_context,omitempty"`
	// TogglingShuffle is for toggling shuffle flag. Optional field.
	TogglingShuffle bool `json:"toggling_shuffle,omitempty"`
	// TogglingRepeatTrack is for toggling the repeat track flag. Optional field.
	TogglingRepeatTrack bool `json:"toggling_repeat_track,omitempty"`
	// TransferringPlayback is for transferring playback between devices. Optional field.
	TransferringPlayback bool `json:"transferring_playback,omitempty"`
}

type PlaybackObject struct {
	// Device is the device that is currently active.
	Device DeviceObject `json:"device"`
	// RepeatState is the repeat state. Possible values are "off", "track", "context".
	RepeatState string `json:"repeat_state"`
	// ShuffleState is whether shuffle is on or off.
	ShuffleState bool `json:"shuffle_state"`
	// Context is the context the playback. Can be null.
	Context *ContextObject `json:"context"`
	// Timestamp is the unix millisecond timestamp when the playback state was last changed (play, pause, skip, scrub, new song, etc.).
	Timestamp int `json:"timestamp"`
	// ProgressMs is the progress into the currently playing track or episode. Can be null.
	ProgressMs *int `json:"progress_ms"`
	// Playing is if something is currently playing, return true.
	Playing bool `json:"is_playing"`
	Track   *TrackObject
	Episode *EpisodeObject
	// CurrentlyPlayingType is the object type of the currently playing item.
	// Can be one of "track", "episode", "ad" or "unknown".
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// Actions is the list of actions allowed to update the user interface based on which playback actions are available within the current context.
	Actions ActionsObject `json:"actions"`
}

func (p *PlaybackObject) UnmarshalJSON(data []byte) error {
	var raw struct {
		// Device is the device that is currently active.
		Device DeviceObject `json:"device"`
		// RepeatState is the repeat state. Possible values are "off", "track", "context".
		RepeatState string `json:"repeat_state"`
		// ShuffleState is whether shuffle is on or off.
		ShuffleState bool `json:"shuffle_state"`
		// Context is the context the playback. Can be null.
		Context *ContextObject `json:"context"`
		// Timestamp is the unix millisecond timestamp when the playback state was last changed (play, pause, skip, scrub, new song, etc.).
		Timestamp int `json:"timestamp"`
		// ProgressMs is the progress into the currently playing track or episode. Can be null.
		ProgressMs *int `json:"progress_ms"`
		// Playing is if something is currently playing, return true.
		Playing bool `json:"is_playing"`
		// Item is the currently playing track or episode. Can be null. TODO update documentation
		Item json.RawMessage `json:"item"`
		// CurrentlyPlayingType is the object type of the currently playing item.
		// Can be one of "track", "episode", "ad" or "unknown".
		CurrentlyPlayingType string `json:"currently_playing_type"`
		// Actions is the list of actions allowed to update the user interface based on which playback actions are available within the current context.
		Actions ActionsObject `json:"actions"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.Device = raw.Device
	p.RepeatState = raw.RepeatState
	p.ShuffleState = raw.ShuffleState
	p.Context = raw.Context
	p.Timestamp = raw.Timestamp
	p.ProgressMs = raw.ProgressMs
	p.Playing = raw.Playing
	p.CurrentlyPlayingType = raw.CurrentlyPlayingType
	p.Actions = raw.Actions

	if raw.Item == nil {
		return nil
	}

	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw.Item, &peek); err != nil {
		return err
	}

	switch peek.Type {
	case "track":
		var t TrackObject
		if err := json.Unmarshal(raw.Item, &t); err != nil {
			return fmt.Errorf("failed to unmarshal track: %w", err)
		}
		p.Track = &t
	case "episode":
		var e EpisodeObject
		if err := json.Unmarshal(raw.Item, &e); err != nil {
			return fmt.Errorf("failed to unmarshal episode: %w", err)
		}
		p.Episode = &e
	default:
		return fmt.Errorf("unknown item type: %s", peek.Type)
	}

	return nil
}

type DeviceObject struct {
	// ID is the device ID. This ID is unique and persistent to some extent.
	// However, this is not guaranteed, and any cached "device_id" should periodically be cleared out and refetched as necessary.
	ID *string `json:"id"`
	// Active if this device is the currently active device.
	Active bool `json:"is_active"`
	// PrivateSession if this device is currently in a private session.
	PrivateSession bool `json:"is_private_session"`
	// Restricted is whether controlling this device is restricted.
	// At present, if this is true, then no Web API commands will be accepted by this device.
	Restricted bool `json:"is_restricted"`
	// Name is a human-readable name for the device.
	Name string `json:"name"`
	// Type is the device type, such as "computer", "smartphone" or "speaker".
	Type string `json:"type"`
	// VolumePercent is the current volume in percent.
	VolumePercent *int `json:"volume_percent"`
	// SupportsVolume if this device can be used to set the volume.
	SupportsVolume bool `json:"supports_volume"`
}

type QueueItemObject struct {
	Track   *TrackObject
	Episode *EpisodeObject
}

func (q *QueueItemObject) UnmarshalJSON(data []byte) error {
	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return err
	}

	switch peek.Type {
	case "track":
		var t TrackObject
		if err := json.Unmarshal(data, &t); err != nil {
			return fmt.Errorf("failed to unmarshal track: %w", err)
		}
		q.Track = &t
	case "episode":
		var e EpisodeObject
		if err := json.Unmarshal(data, &e); err != nil {
			return fmt.Errorf("failed to unmarshal episode: %w", err)
		}
		q.Episode = &e
	default:
		return fmt.Errorf("unknown item type: %s", peek.Type)
	}

	return nil
}

type UserQueue struct {
	// CurrentlyPlaying is the currently playing track or episode. Can be null.
	CurrentlyPlaying *QueueItemObject `json:"currently_playing"`
	// Queue is the tracks or episodes in the queue. Can be empty.
	Queue []QueueItemObject `json:"queue"`
}

type PlayHistoryObject struct {
	// Track is the track the user listened to.
	Track TrackObject `json:"track"`
	// PlayedAt is the date and time the track was played.
	PlayedAt string `json:"played_at"`
	// Context is the context the track was played from.
	Context ContextObject `json:"context"`
}

type TransferPlaybackRequest struct {
	// DeviceIDs is an array containing the ID of the device on which playback should be started/transferred.
	// Note: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return 400 Bad Request.
	DeviceIDs []string `json:"device_ids"`
	// Play if true, ensure playback happens on a new device. If false or not provided, keep the current playback state.
	Play bool `json:"play"`
}

type StartResumePlaybackRequest struct {
	// ContextURI is a Spotify URI of the context to play. Valid contexts are albums, artists & playlists.
	ContextURI *string `json:"context_uri,omitempty"`
	// URIs is an array of the Spotify track URIs to play.
	URIs []string `json:"uris,omitempty"`
	// Offset indicates from where in the context playback should start.
	// Only available when context_uri corresponds to an album or playlist object "position" is zero-based and can’t be negative.
	Offset map[string]any `json:"offset,omitempty"`
	// PositionMs is the position of the playback (in milliseconds).
	PositionMs *int `json:"position_ms,omitempty"`
}
