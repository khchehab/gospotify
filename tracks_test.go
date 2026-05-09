package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTrack_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetTrack(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetTrack_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tracks/track1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"track1","name":"Test Track","type":"track","duration_ms":210000}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	track, err := client.GetTrack(context.Background(), "track1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if track.ID != "track1" {
		t.Errorf("got ID %q, want track1", track.ID)
	}
	if track.Name != "Test Track" {
		t.Errorf("got Name %q, want 'Test Track'", track.Name)
	}
	if track.DurationMs != 210000 {
		t.Errorf("got DurationMs %d, want 210000", track.DurationMs)
	}
}

func TestGetTrack_WithMarket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("market") != "GB" {
			t.Errorf("expected market=GB, got %q", r.URL.Query().Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"track1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetTrack(context.Background(), "track1", WithMarket("GB"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetTrack_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"Track not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetTrack(context.Background(), "bad-id")
	if err == nil {
		t.Fatal("expected error")
	}
	var errResp *ErrorResponse
	if !errors.As(err, &errResp) {
		t.Fatalf("expected *ErrorResponse, got %T", err)
	}
	if errResp.ErrorObject.Status != 404 {
		t.Errorf("got status %d, want 404", errResp.ErrorObject.Status)
	}
}

func TestGetUserSavedTracks_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/tracks" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":5,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	tracks, err := client.GetUserSavedTracks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tracks.Total != 5 {
		t.Errorf("got Total %d, want 5", tracks.Total)
	}
}

func TestGetUserSavedTracks_WithOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("limit") != "10" {
			t.Errorf("expected limit=10, got %q", q.Get("limit"))
		}
		if q.Get("market") != "US" {
			t.Errorf("expected market=US, got %q", q.Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":0,"limit":10,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetUserSavedTracks(context.Background(), WithLimit(10), WithMarket("US"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
