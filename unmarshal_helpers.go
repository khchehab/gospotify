package gospotify

import (
	"encoding/json"
	"fmt"
)

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
