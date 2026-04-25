package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlaybackState_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/player" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"is_playing":true,"repeat_state":"off","shuffle_state":false,"device":{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":50}}`))
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

func TestGetPlaybackState_204_ReturnsNilPlayback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	state, err := client.GetPlaybackState(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 204 means no active playback — state will be zero-value (not nil since it's a value type)
	if state.Playing {
		t.Error("expected Playing=false for 204 response")
	}
}

func TestGetPlaybackState_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":{"status":401,"message":"Unauthorized"}}`))
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
		w.Write([]byte(`[{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":80,"is_active":true,"is_private_session":false,"is_restricted":false}]`))
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
		w.Write([]byte(`{"is_playing":true,"repeat_state":"off","shuffle_state":false,"device":{"id":"dev1","name":"My Speaker","type":"Speaker","volume_percent":50}}`))
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
	err := client.SetRepeatMode(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty state")
	}
	if err.Error() != "state cannot be empty" {
		t.Errorf("got %q, want 'state cannot be empty'", err.Error())
	}
}

func TestSetRepeatMode_InvalidState_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.SetRepeatMode(context.Background(), "invalid")
	if err == nil {
		t.Fatal("expected error for invalid state")
	}
}

func TestSetRepeatMode_ValidStates(t *testing.T) {
	states := []string{"track", "context", "off"}
	for _, state := range states {
		t.Run(state, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Query().Get("state") != state {
					t.Errorf("expected state=%s, got %q", state, r.URL.Query().Get("state"))
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			client := newTestClient(srv.URL)
			err := client.SetRepeatMode(context.Background(), state)
			if err != nil {
				t.Fatalf("unexpected error for state %q: %v", state, err)
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
		w.Write([]byte(`{"href":"","total":5,"limit":20,"cursors":{},"items":[]}`))
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
		w.Write([]byte(`{"href":"","total":0,"limit":10,"cursors":{},"items":[]}`))
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
		w.Write([]byte(`{"currently_playing":null,"queue":[]}`))
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/me/player/queue" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("uri") != "spotify:track:abc123" {
			t.Errorf("expected uri=spotify:track:abc123, got %q", r.URL.Query().Get("uri"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.AddItemToPlaybackQueue(context.Background(), "spotify:track:abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConcatenatePosition(t *testing.T) {
	got := concatenatePosition("/me/player/seek", 30000)
	want := "/me/player/seek?position_ms=30000"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConcatenateVolumePercent(t *testing.T) {
	got := concatenateVolumePercent("/me/player/volume", 75)
	want := "/me/player/volume?volume_percent=75"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConcatenateState(t *testing.T) {
	got := concatenateState("/me/player/repeat", "track")
	want := "/me/player/repeat?state=track"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConcatenateStateB(t *testing.T) {
	got := concatenateStateB("/me/player/shuffle", true)
	want := "/me/player/shuffle?state=true"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	got = concatenateStateB("/me/player/shuffle", false)
	want = "/me/player/shuffle?state=false"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConcatenateURI(t *testing.T) {
	got := concatenateURI("/me/player/queue", "spotify:track:abc")
	want := "/me/player/queue?uri=spotify:track:abc"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
