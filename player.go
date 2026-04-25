package gospotify

import "context"

// GetPlaybackState gets information about the user’s current playback state, including track or episode, progress, and active device.
// If the playback is not available or not active, the playback object returned will be nil, along with a nil error.
// TODO this should be improved to return an error instead of `nil, nil`.
//
// QueryOptions that can be used are:
//   - [WithMarket]: An ISO 3166-1 alpha-2 country code.
//   - [WithAdditionalTypes]: The list of item types that your client supports besides the default track type.
func (c *Client) GetPlaybackState(ctx context.Context, opts ...QueryOption) (*PlaybackObject, error) {
	var playback *PlaybackObject
	if err := c.get(ctx, "/me/player", &playback, opts...); err != nil {
		return nil, err
	}
	return playback, nil
}
