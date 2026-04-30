package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentUserProfile_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"user1","display_name":"Test User","type":"user","email":"test@example.com"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if profile.ID != "user1" {
		t.Errorf("got ID %q, want user1", profile.ID)
	}
	if profile.DisplayName == nil || *profile.DisplayName != "Test User" {
		t.Errorf("got DisplayName %v, want 'Test User'", profile.DisplayName)
	}
}

func TestGetCurrentUserProfile_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":{"status":401,"message":"Unauthorized"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetCurrentUserProfile(context.Background())
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

func TestGetUserTopArtists_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/top/artists" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":10,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	artists, err := client.GetUserTopArtists(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if artists.Total != 10 {
		t.Errorf("got Total %d, want 10", artists.Total)
	}
}

func TestGetUserTopArtists_WithTimeRange(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("time_range") != "long_term" {
			t.Errorf("expected time_range=long_term, got %q", r.URL.Query().Get("time_range"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":0,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithTimeRange(LongTerm))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetUserTopTracks_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/top/tracks" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":20,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	tracks, err := client.GetUserTopTracks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tracks.Total != 20 {
		t.Errorf("got Total %d, want 20", tracks.Total)
	}
}

func TestGetUserTopTracks_WithOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("time_range") != "short_term" {
			t.Errorf("expected time_range=short_term, got %q", q.Get("time_range"))
		}
		if q.Get("limit") != "5" {
			t.Errorf("expected limit=5, got %q", q.Get("limit"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":0,"limit":5,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetUserTopTracks(context.Background(), WithTimeRange(ShortTerm), WithLimit(5))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetFollowedArtists_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/following" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("type") != "artist" {
			t.Errorf("expected type=artist, got %q", r.URL.Query().Get("type"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"artists":{"href":"","total":7,"limit":20,"cursors":{},"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	artists, err := client.GetFollowedArtists(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if artists.Total != 7 {
		t.Errorf("got Total %d, want 7", artists.Total)
	}
}

func TestGetFollowedArtists_WithAfter(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("after") != "lastArtistID" {
			t.Errorf("expected after=lastArtistID, got %q", r.URL.Query().Get("after"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"artists":{"href":"","total":0,"limit":20,"cursors":{},"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background(), WithAfter("lastArtistID"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
