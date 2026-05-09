package gospotify

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPlaybackState_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"is_playing":true,"repeat_state":"off","shuffle_state":false,"device":{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":50}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	state, err := client.GetPlaybackState(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !state.Playing {
		t.Error("expected Playing=true")
	}
}

func TestGetPlaybackState_204_ReturnsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	state, err := client.GetPlaybackState(context.Background())
	if err == nil {
		t.Fatal("expected error when playback is not available")
	}
	if err.Error() != "playback is not available or active" {
		t.Errorf("got %q, want 'playback is not available or active'", err.Error())
	}
	if state != nil {
		t.Error("expected nil state when playback is not available")
	}
}

func TestGetPlaybackState_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":{"status":401,"message":"Unauthorized"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetPlaybackState(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	var errResp *ErrorResponse
	if !errors.As(err, &errResp) {
		t.Fatalf("expected *ErrorResponse, got %T", err)
	}
	if errResp.ErrorObject.Status != 401 {
		t.Errorf("got status %d, want 401", errResp.ErrorObject.Status)
	}
}

func TestTransferPlayback_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.TransferPlayback(context.Background(), TransferPlaybackRequest{DeviceIDs: []string{"dev1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetAvailableDevices_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player/devices" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":80,"is_active":true,"is_private_session":false,"is_restricted":false}]`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	devices, err := client.GetAvailableDevices(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(devices) != 1 {
		t.Fatalf("got %d devices, want 1", len(devices))
	}
	if devices[0].ID == nil || *devices[0].ID != "dev1" {
		t.Errorf("got ID %v, want dev1", devices[0].ID)
	}
}

func TestGetCurrentPlayingTrack_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player/currently-playing" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"is_playing":true,"repeat_state":"off","shuffle_state":false,"device":{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":50}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	track, err := client.GetCurrentPlayingTrack(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !track.Playing {
		t.Error("expected Playing=true")
	}
}

func TestStartResumePlayback_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/play" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.StartResumePlayback(context.Background(), StartResumePlaybackRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStartResumePlayback_WithDeviceID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("device_id") != "dev1" {
			t.Errorf("expected device_id=dev1, got %q", r.URL.Query().Get("device_id"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.StartResumePlayback(context.Background(), StartResumePlaybackRequest{}, WithDeviceID("dev1"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPausePlayback_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/pause" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.PausePlayback(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSkipToNext_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/next" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.SkipToNext(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSkipToPrevious_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/previous" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.SkipToPrevious(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSeekToPosition_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/seek" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("position_ms") != "30000" {
			t.Errorf("expected position_ms=30000, got %q", r.URL.Query().Get("position_ms"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.SeekToPosition(context.Background(), 30000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSetRepeatMode_EmptyState_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.SetRepeatMode(context.Background(), RepeatState(""))
	if err == nil {
		t.Fatal("expected error for empty state")
	}
	if err.Error() != "state cannot be empty" {
		t.Errorf("got %q, want 'state cannot be empty'", err.Error())
	}
}

func TestSetRepeatMode_InvalidState_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.SetRepeatMode(context.Background(), RepeatState("invalid"))
	if err == nil {
		t.Fatal("expected error for invalid state")
	}
}

func TestSetRepeatMode_ValidStates(t *testing.T) {
	cases := []struct {
		name  string
		state RepeatState
	}{
		{"track", RepeatTrack},
		{"context", RepeatContext},
		{"off", RepeatOff},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Query().Get("state") != string(tc.state) {
					t.Errorf("expected state=%s, got %q", tc.state, r.URL.Query().Get("state"))
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			client := newTestClient(srv.URL)
			err := client.SetRepeatMode(context.Background(), tc.state)
			if err != nil {
				t.Fatalf("unexpected error for state %q: %v", tc.state, err)
			}
		})
	}
}

func TestSetPlaybackVolume_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/volume" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("volume_percent") != "75" {
			t.Errorf("expected volume_percent=75, got %q", r.URL.Query().Get("volume_percent"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.SetPlaybackVolume(context.Background(), 75)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTogglePlaybackShuffle_True(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/shuffle" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("state") != "true" {
			t.Errorf("expected state=true, got %q", r.URL.Query().Get("state"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.TogglePlaybackShuffle(context.Background(), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTogglePlaybackShuffle_False(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != "false" {
			t.Errorf("expected state=false, got %q", r.URL.Query().Get("state"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.TogglePlaybackShuffle(context.Background(), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetRecentlyPlayedTracks_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player/recently-played" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":5,"limit":20,"cursors":{},"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	history, err := client.GetRecentlyPlayedTracks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if history.Total != 5 {
		t.Errorf("got Total %d, want 5", history.Total)
	}
}

func TestGetRecentlyPlayedTracks_WithCursorOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("limit") != "10" {
			t.Errorf("expected limit=10, got %q", q.Get("limit"))
		}
		if q.Get("before") != "1609459200000" {
			t.Errorf("expected before=1609459200000, got %q", q.Get("before"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":0,"limit":10,"cursors":{},"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetRecentlyPlayedTracks(context.Background(), WithLimit(10), WithBeforeMs(1609459200000))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetUserQueue_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player/queue" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"currently_playing":null,"queue":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	queue, err := client.GetUserQueue(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if queue == nil {
		t.Fatal("expected non-nil queue")
	}
}

func TestAddItemToPlaybackQueue_Success(t *testing.T) {
	uri := "spotify:track:abc123"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/queue" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("uri") != uri {
			t.Errorf("expected uri=%s in query, got %q", uri, r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.AddItemToPlaybackQueue(context.Background(), uri)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAppendQueryParams_PositionMs(t *testing.T) {
	got := appendQueryParams("/me/player/seek", requiredQueryParam{key: "position_ms", value: 30000})
	want := "/me/player/seek?position_ms=30000"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_VolumePercent(t *testing.T) {
	got := appendQueryParams("/me/player/volume", requiredQueryParam{key: "volume_percent", value: 75})
	want := "/me/player/volume?volume_percent=75"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_StateString(t *testing.T) {
	got := appendQueryParams("/me/player/repeat", requiredQueryParam{key: "state", value: "track"})
	want := "/me/player/repeat?state=track"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_StateBool(t *testing.T) {
	got := appendQueryParams("/me/player/shuffle", requiredQueryParam{key: "state", value: true})
	want := "/me/player/shuffle?state=true"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	got = appendQueryParams("/me/player/shuffle", requiredQueryParam{key: "state", value: false})
	want = "/me/player/shuffle?state=false"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// ---- PlaybackObject.UnmarshalJSON ----

func TestPlaybackObject_UnmarshalJSON_TrackItem(t *testing.T) {
	data := `{
		"is_playing": true,
		"repeat_state": "off",
		"shuffle_state": false,
		"currently_playing_type": "track",
		"item": {"type": "track", "id": "t1", "name": "My Track", "duration_ms": 180000}
	}`
	var p PlaybackObject
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Track == nil {
		t.Fatal("expected Track to be non-nil")
	}
	if p.Episode != nil {
		t.Error("expected Episode to be nil")
	}
	if p.Track.ID != "t1" {
		t.Errorf("got Track.ID %q, want t1", p.Track.ID)
	}
}

func TestPlaybackObject_UnmarshalJSON_EpisodeItem(t *testing.T) {
	data := `{
		"is_playing": true,
		"repeat_state": "off",
		"shuffle_state": false,
		"currently_playing_type": "episode",
		"item": {"type": "episode", "id": "e1", "name": "My Episode", "duration_ms": 3600000}
	}`
	var p PlaybackObject
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Episode == nil {
		t.Fatal("expected Episode to be non-nil")
	}
	if p.Track != nil {
		t.Error("expected Track to be nil")
	}
	if p.Episode.ID != "e1" {
		t.Errorf("got Episode.ID %q, want e1", p.Episode.ID)
	}
}

func TestPlaybackObject_UnmarshalJSON_NoItem_BothNil(t *testing.T) {
	// item field is absent — raw.Item will be nil, so we return early with no error.
	data := `{
		"is_playing": false,
		"repeat_state": "off",
		"shuffle_state": false,
		"currently_playing_type": "unknown"
	}`
	var p PlaybackObject
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Track != nil {
		t.Error("expected Track to be nil when item is absent")
	}
	if p.Episode != nil {
		t.Error("expected Episode to be nil when item is absent")
	}
}

func TestPlaybackObject_UnmarshalJSON_UnknownItemType_ReturnsError(t *testing.T) {
	data := `{
		"is_playing": true,
		"item": {"type": "unknown_type", "id": "x1"}
	}`
	var p PlaybackObject
	err := json.Unmarshal([]byte(data), &p)
	if err == nil {
		t.Fatal("expected error for unknown item type")
	}
	if !strings.Contains(err.Error(), "unknown item type") {
		t.Errorf("expected error to contain 'unknown item type', got %q", err.Error())
	}
}

// ---- QueueItemObject.UnmarshalJSON ----

func TestQueueItemObject_UnmarshalJSON_Track(t *testing.T) {
	data := `{"type": "track", "id": "t2", "name": "Queue Track", "duration_ms": 200000}`
	var q QueueItemObject
	if err := json.Unmarshal([]byte(data), &q); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q.Track == nil {
		t.Fatal("expected Track to be non-nil")
	}
	if q.Episode != nil {
		t.Error("expected Episode to be nil")
	}
	if q.Track.ID != "t2" {
		t.Errorf("got Track.ID %q, want t2", q.Track.ID)
	}
}

func TestQueueItemObject_UnmarshalJSON_Episode(t *testing.T) {
	data := `{"type": "episode", "id": "e2", "name": "Queue Episode", "duration_ms": 1800000}`
	var q QueueItemObject
	if err := json.Unmarshal([]byte(data), &q); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q.Episode == nil {
		t.Fatal("expected Episode to be non-nil")
	}
	if q.Track != nil {
		t.Error("expected Track to be nil")
	}
	if q.Episode.ID != "e2" {
		t.Errorf("got Episode.ID %q, want e2", q.Episode.ID)
	}
}

func TestQueueItemObject_UnmarshalJSON_UnknownType_ReturnsError(t *testing.T) {
	data := `{"type": "ad", "id": "ad1"}`
	var q QueueItemObject
	err := json.Unmarshal([]byte(data), &q)
	if err == nil {
		t.Fatal("expected error for unknown item type")
	}
	if !strings.Contains(err.Error(), "unknown item type") {
		t.Errorf("expected error to contain 'unknown item type', got %q", err.Error())
	}
}

// ---- GetAlbum error wrapping ----

func TestGetAlbum_404Error_ContainsFunctionName(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetAlbum(context.Background(), "bad-id")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "GetAlbum") {
		t.Errorf("expected error to contain 'GetAlbum', got %q", err.Error())
	}
}

// ---- GetTrack error wrapping ----

func TestGetTrack_404Error_ContainsFunctionName(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetTrack(context.Background(), "bad-id")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "GetTrack") {
		t.Errorf("expected error to contain 'GetTrack', got %q", err.Error())
	}
}

// ---- GetPlaylist error wrapping ----

func TestGetPlaylist_404Error_ContainsFunctionName(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetPlaylist(context.Background(), "bad-id")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "GetPlaylist") {
		t.Errorf("expected error to contain 'GetPlaylist', got %q", err.Error())
	}
}

// ---- SearchForItem error wrapping ----

func TestSearchForItem_404Error_ContainsFunctionName(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.SearchForItem(context.Background(), "test", []ItemType{ItemTypeTrack})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "SearchForItem") {
		t.Errorf("expected error to contain 'SearchForItem', got %q", err.Error())
	}
}
