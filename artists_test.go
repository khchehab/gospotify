package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArtist_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetArtist(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetArtist_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/artists/artist1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"artist1","name":"Test Artist","type":"artist"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	artist, err := client.GetArtist(context.Background(), "artist1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if artist.ID != "artist1" {
		t.Errorf("got ID %q, want artist1", artist.ID)
	}
	if artist.Name != "Test Artist" {
		t.Errorf("got Name %q, want 'Test Artist'", artist.Name)
	}
}

func TestGetArtist_APIError_ReturnsErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":{"status":404,"message":"Artist not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetArtist(context.Background(), "bad-id")
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

func TestGetArtistAlbums_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetArtistAlbums(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
}

func TestGetArtistAlbums_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/artists/artist1/albums" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":3,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	albums, err := client.GetArtistAlbums(context.Background(), "artist1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if albums.Total != 3 {
		t.Errorf("got Total %d, want 3", albums.Total)
	}
}

func TestGetArtistAlbums_WithIncludeGroups(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("include_groups") != "album,single" {
			t.Errorf("expected include_groups=album,single, got %q", r.URL.Query().Get("include_groups"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"href":"","total":0,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetArtistAlbums(context.Background(), "artist1", WithIncludeGroups("album", "single"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
