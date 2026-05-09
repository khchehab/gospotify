package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetShow_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetShow(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetShow_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shows/show1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"show1","name":"Test Show","type":"show","total_episodes":50}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	show, err := client.GetShow(context.Background(), "show1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if show.ID != "show1" {
		t.Errorf("got ID %q, want show1", show.ID)
	}
	if show.Name != "Test Show" {
		t.Errorf("got Name %q, want 'Test Show'", show.Name)
	}
	if show.TotalEpisodes != 50 {
		t.Errorf("got TotalEpisodes %d, want 50", show.TotalEpisodes)
	}
}

func TestGetShow_WithMarket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("market") != "DE" {
			t.Errorf("expected market=DE, got %q", r.URL.Query().Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"show1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetShow(context.Background(), "show1", WithMarket("DE"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetShow_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"Show not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetShow(context.Background(), "bad-id")
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

func TestGetShowEpisodes_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetShowEpisodes(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
}

func TestGetShowEpisodes_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shows/show1/episodes" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":50,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	episodes, err := client.GetShowEpisodes(context.Background(), "show1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if episodes.Total != 50 {
		t.Errorf("got Total %d, want 50", episodes.Total)
	}
}

func TestGetUserSavedShows_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/shows" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":4,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	shows, err := client.GetUserSavedShows(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if shows.Total != 4 {
		t.Errorf("got Total %d, want 4", shows.Total)
	}
}
