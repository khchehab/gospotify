package gospotify

import (
	"encoding/json"
	"fmt"
)

// unmarshalTrackOrEpisode peeks at the "type" field of data and unmarshals the payload into
// either a TrackObject or an EpisodeObject. Exactly one of the two return pointers will be
// non-nil on success. Returns an error for unknown types or malformed JSON.
func unmarshalTrackOrEpisode(data []byte) (*TrackObject, *EpisodeObject, error) {
	var peek struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return nil, nil, err
	}

	switch peek.Type {
	case "track":
		var t TrackObject
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal track: %w", err)
		}
		return &t, nil, nil
	case "episode":
		var e EpisodeObject
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal episode: %w", err)
		}
		return nil, &e, nil
	default:
		return nil, nil, fmt.Errorf("unknown item type: %s", peek.Type)
	}
}
