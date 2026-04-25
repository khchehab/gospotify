package gospotify

import (
	"context"
	"errors"
	"fmt"
)

// GetPlaybackState gets information about the user’s current playback state, including track or episode, progress, and active device.
// If the playback is not available or not active, the playback object returned will be nil, along with a nil error.
// TODO this should be improved to return an error instead of `nil, nil`.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetPlaybackState(ctx context.Context, opts ...QueryOption) (*PlaybackObject, error) {
	var playback PlaybackObject
	if err := c.get(ctx, "/me/player", &playback, opts...); err != nil {
		return nil, err
	}
	return &playback, nil
}

// TransferPlayback transfers playback to a new device and optionally begin playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
func (c *Client) TransferPlayback(ctx context.Context, body TransferPlaybackRequest) error {
	return c.put(ctx, "/me/player", body, "", nil)
}

// GetAvailableDevices gets information about a user’s available Spotify Connect devices.
// Some device models are not supported and will not be listed in the API response.
func (c *Client) GetAvailableDevices(ctx context.Context) ([]DeviceObject, error) {
	var devices []DeviceObject
	if err := c.get(ctx, "/me/player/devices", &devices); err != nil {
		return nil, err
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
		return nil, err
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
	return c.put(ctx, "/me/player/play", body, "", nil, opts...)
}

// PausePlayback pauses playback on the user's account.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) PausePlayback(ctx context.Context, opts ...QueryOption) error {
	return c.put(ctx, "/me/player/pause", nil, "", nil, opts...)
}

// SkipToNext skips to next track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SkipToNext(ctx context.Context, opts ...QueryOption) error {
	return c.post(ctx, "/me/player/next", nil, nil, opts...)
}

// SkipToPrevious skips to previous track in the user’s queue.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SkipToPrevious(ctx context.Context, opts ...QueryOption) error {
	return c.post(ctx, "/me/player/previous", nil, nil, opts...)
}

// SeekToPosition seeks to the given position in the user’s currently playing track.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SeekToPosition(ctx context.Context, positionMs int, opts ...QueryOption) error {
	return c.put(ctx, concatenatePosition("/me/player/seek", positionMs), nil, "", nil, opts...)
}

// SetRepeatMode sets the repeat mode for the user's playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SetRepeatMode(ctx context.Context, state string, opts ...QueryOption) error {
	if state == "" {
		return errors.New("state cannot be empty")
	}
	if state != "track" && state != "context" && state != "off" {
		return fmt.Errorf("invalid state: %s", state)
	}
	return c.put(ctx, concatenateState("/me/player/repeat", state), nil, "", nil, opts...)
}

// SetPlaybackVolume sets the volume for the user’s current playback device.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) SetPlaybackVolume(ctx context.Context, volumePercent int, opts ...QueryOption) error {
	return c.put(ctx, concatenateVolumePercent("/me/player/volume", volumePercent), nil, "", nil, opts...)
}

// TogglePlaybackShuffle toggles shuffle on or off for user’s playback.
// This API only works for users who have Spotify Premium.
// The order of execution is not guaranteed when you use this API with other Player API endpoints.
//
// QueryOptions that can be used are:
//   - [WithDeviceID]: The id of the device this command is targeting.
//     If not supplied, the user's currently active device is the target.
func (c *Client) TogglePlaybackShuffle(ctx context.Context, state bool, opts ...QueryOption) error {
	return c.put(ctx, concatenateStateB("/me/player/shuffle", state), nil, "", nil, opts...)
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
		return nil, err
	}
	return &playHistory, nil
}

// GetUserQueue gets the list of objects that make up the user's queue.
func (c *Client) GetUserQueue(ctx context.Context) (*UserQueue, error) {
	var userQueue UserQueue
	if err := c.get(ctx, "/me/player/queue", &userQueue); err != nil {
		return nil, err
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
	return c.post(ctx, concatenateURI("/me/player/queue", uri), nil, nil, opts...)
}

// concatenatePosition concatenates the position to the endpoint.
func concatenatePosition(endpoint string, positionMs int) string {
	return fmt.Sprintf("%s?position_ms=%d", endpoint, positionMs)
}

// concatenateVolumePercent concatenates the volume percent to the endpoint.
func concatenateVolumePercent(endpoint string, volumePercent int) string {
	return fmt.Sprintf("%s?volume_percent=%d", endpoint, volumePercent)
}

// concatenateState concatenates the state to the endpoint.
func concatenateState(endpoint, state string) string {
	return fmt.Sprintf("%s?state=%s", endpoint, state)
}

// concatenateStateB concatenates the state (boolean) to the endpoint.
func concatenateStateB(endpoint string, state bool) string {
	return fmt.Sprintf("%s?state=%t", endpoint, state)
}

// concatenateURI concatenates the URI to the endpoint.
func concatenateURI(endpoint, uri string) string {
	return fmt.Sprintf("%s?uri=%s", endpoint, uri)
}
