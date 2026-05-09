package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEpisode_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetEpisode(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetEpisode_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/episodes/ep1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"ep1","name":"Episode 1","type":"episode","duration_ms":1800000}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	episode, err := client.GetEpisode(context.Background(), "ep1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if episode.ID != "ep1" {
		t.Errorf("got ID %q, want ep1", episode.ID)
	}
	if episode.Name != "Episode 1" {
		t.Errorf("got Name %q, want 'Episode 1'", episode.Name)
	}
}

func TestGetEpisode_WithMarket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("market") != "US" {
			t.Errorf("expected market=US, got %q", r.URL.Query().Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"ep1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetEpisode(context.Background(), "ep1", WithMarket("US"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetEpisode_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"Episode not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetEpisode(context.Background(), "bad-id")
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

func TestGetUserSavedEpisodes_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/episodes" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":3,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	episodes, err := client.GetUserSavedEpisodes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if episodes.Total != 3 {
		t.Errorf("got Total %d, want 3", episodes.Total)
	}
}

func TestGetUserSavedEpisodes_WithOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("limit") != "5" {
			t.Errorf("expected limit=5, got %q", q.Get("limit"))
		}
		if q.Get("offset") != "2" {
			t.Errorf("expected offset=2, got %q", q.Get("offset"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":0,"limit":5,"offset":2,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetUserSavedEpisodes(context.Background(), WithLimit(5), WithOffset(2))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
