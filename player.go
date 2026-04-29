package gospotify

import (
	"context"
	"fmt"
)

// GetPlaybackState gets information about the user’s current playback state, including track or episode, progress, and active device.
// If the playback is not available or not active, [ErrNoActivePlayback] is returned.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetPlaybackState(ctx context.Context, opts ...QueryOption) (*PlaybackObject, error) {
	var playback *PlaybackObject
	if err := c.get(ctx, "/me/player", &playback, opts...); err != nil {
		return nil, fmt.Errorf("GetPlaybackState: %w", err)
	}
	if playback == nil {
		return nil, ErrNoActivePlayback
	}
	return playback, nil
}

// TransferPlayback transfers playback to a new device and optionally begin playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
func (c *Client) TransferPlayback(ctx context.Context, body TransferPlaybackRequest) error {
	if err := c.put(ctx, "/me/player", body, "", nil); err != nil {
		return fmt.Errorf("TransferPlayback: %w", err)
	}
	return nil
}

// GetAvailableDevices gets information about a user’s available Spotify Connect devices.
// Some device models are not supported and will not be listed in the API response.
func (c *Client) GetAvailableDevices(ctx context.Context) ([]DeviceObject, error) {
	var devices []DeviceObject
	if err := c.get(ctx, "/me/player/devices", &devices); err != nil {
		return nil, fmt.Errorf("GetAvailableDevices: %w", err)
	}
	return devices, nil
}

// GetCurrentPlayingTrack gets the object currently being played on the user's Spotify account.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetCurrentPlayingTrack(ctx context.Context, opts ...QueryOption) (*PlaybackObject, error) {
	var playback PlaybackObject
	if err := c.get(ctx, "/me/player/currently-playing", &playback, opts...); err != nil {
		return nil, fmt.Errorf("GetCurrentPlayingTrack: %w", err)
	}
	return &playback, nil
}

// StartResumePlayback starts a new context or resumes current playback on the user's active device.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) StartResumePlayback(ctx context.Context, body StartResumePlaybackRequest, opts ...QueryOption) error {
	if err := c.put(ctx, "/me/player/play", body, "", nil, opts...); err != nil {
		return fmt.Errorf("StartResumePlayback: %w", err)
	}
	return nil
}

// PausePlayback pauses playback on the user's account.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) PausePlayback(ctx context.Context, opts ...QueryOption) error {
	if err := c.put(ctx, "/me/player/pause", nil, "", nil, opts...); err != nil {
		return fmt.Errorf("PausePlayback: %w", err)
	}
	return nil
}

// SkipToNext skips to next track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SkipToNext(ctx context.Context, opts ...QueryOption) error {
	if err := c.post(ctx, "/me/player/next", nil, nil, opts...); err != nil {
		return fmt.Errorf("SkipToNext: %w", err)
	}
	return nil
}

// SkipToPrevious skips to previous track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SkipToPrevious(ctx context.Context, opts ...QueryOption) error {
	if err := c.post(ctx, "/me/player/previous", nil, nil, opts...); err != nil {
		return fmt.Errorf("SkipToPrevious: %w", err)
	}
	return nil
}

// SeekToPosition seeks to the given position in the user’s currently playing track.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SeekToPosition(ctx context.Context, positionMs int, opts ...QueryOption) error {
	if err := c.put(ctx, appendQueryParams("/me/player/seek", requiredQueryParam{
		key:   "position_ms",
		value: positionMs,
	}), nil, "", nil, opts...); err != nil {
		return fmt.Errorf("SeekToPosition: %w", err)
	}
	return nil
}

// SetRepeatMode sets the repeat mode for the user's playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SetRepeatMode(ctx context.Context, state RepeatState, opts ...QueryOption) error {
	if err := requireNonEmpty("state", string(state)); err != nil {
		return err
	}
	if !state.Valid() {
		return fmt.Errorf("invalid state: %s", state)
	}
	if err := c.put(ctx, appendQueryParams("/me/player/repeat", requiredQueryParam{
		key:   "state",
		value: string(state),
	}), nil, "", nil, opts...); err != nil {
		return fmt.Errorf("SetRepeatMode %q: %w", state, err)
	}
	return nil
}

// SetPlaybackVolume sets the volume for the user’s current playback device.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SetPlaybackVolume(ctx context.Context, volumePercent int, opts ...QueryOption) error {
	if err := c.put(ctx, appendQueryParams("/me/player/volume", requiredQueryParam{
		key:   "volume_percent",
		value: volumePercent,
	}), nil, "", nil, opts...); err != nil {
		return fmt.Errorf("SetPlaybackVolume %v: %w", volumePercent, err)
	}
	return nil
}

// TogglePlaybackShuffle toggles shuffle on or off for user’s playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) TogglePlaybackShuffle(ctx context.Context, state bool, opts ...QueryOption) error {
	if err := c.put(ctx, appendQueryParams("/me/player/shuffle", requiredQueryParam{
		key:   "state",
		value: state,
	}), nil, "", nil, opts...); err != nil {
		return fmt.Errorf("TogglePlaybackShuffle %t: %w", state, err)
	}
	return nil
}

// GetRecentlyPlayedTracks gets tracks from the current user's recently played tracks.
// Note: Currently doesn't support podcast episodes.
//
// QueryOptions that can be used are:
//   - [WithLimit]: The maximum number of items to return.
//   - [WithAfterMs]: Returns all items after (but not including) this cursor position. If after is specified, before must not be specified.
//   - [WithBeforeMs]: Returns all items before (but not including) this cursor position. If before is specified, after must not be specified.
func (c *Client) GetRecentlyPlayedTracks(ctx context.Context, opts ...QueryOption) (*Cursor[PlayHistoryObject], error) {
	var playHistory Cursor[PlayHistoryObject]
	if err := c.get(ctx, "/me/player/recently-played", &playHistory, opts...); err != nil {
		return nil, fmt.Errorf("GetRecentlyPlayedTracks: %w", err)
	}
	return &playHistory, nil
}

// GetUserQueue gets the list of objects that make up the user's queue.
func (c *Client) GetUserQueue(ctx context.Context) (*UserQueue, error) {
	var userQueue UserQueue
	if err := c.get(ctx, "/me/player/queue", &userQueue); err != nil {
		return nil, fmt.Errorf("GetUserQueue: %w", err)
	}
	return &userQueue, nil
}

// AddItemToPlaybackQueue adds an item to be played next in the user's current playback queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) AddItemToPlaybackQueue(ctx context.Context, uri string, opts ...QueryOption) error {
	if err := c.post(ctx, appendQueryParams("/me/player/queue", requiredQueryParam{
		key:   "uri",
		value: uri,
	}), nil, nil, opts...); err != nil {
		return fmt.Errorf("AddItemToPlaybackQueue %q: %w", uri, err)
	}
	return nil
}
