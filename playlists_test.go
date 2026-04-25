package gospotify

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlaylist_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetPlaylist(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
	if err.Error() != "playlist id cannot be empty" {
		t.Errorf("got %q, want 'playlist id cannot be empty'", err.Error())
	}
}

func TestGetPlaylist_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/playlists/pl1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"pl1","name":"My Playlist","type":"playlist","snapshot_id":"snap1","items":{"href":"","total":0,"limit":20,"offset":0,"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	playlist, err := client.GetPlaylist(context.Background(), "pl1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if playlist.ID != "pl1" {
		t.Errorf("got ID %q, want pl1", playlist.ID)
	}
	if playlist.Name != "My Playlist" {
		t.Errorf("got Name %q, want 'My Playlist'", playlist.Name)
	}
}

func TestGetPlaylist_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"status":404,"message":"Playlist not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetPlaylist(context.Background(), "bad-id")
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

func TestChangePlaylistDetails_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.ChangePlaylistDetails(context.Background(), "", PlaylistDetailRequest{})
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestChangePlaylistDetails_Success(t *testing.T) {
	var capturedBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/playlists/pl1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &capturedBody)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.ChangePlaylistDetails(context.Background(), "pl1", PlaylistDetailRequest{Name: "Updated Name"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedBody["name"] != "Updated Name" {
		t.Errorf("got body name %v, want 'Updated Name'", capturedBody["name"])
	}
}

func TestGetPlaylistItems_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetPlaylistItems(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestGetPlaylistItems_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/playlists/pl1/items" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":3,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	items, err := client.GetPlaylistItems(context.Background(), "pl1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if items.Total != 3 {
		t.Errorf("got Total %d, want 3", items.Total)
	}
}

func TestUpdatePlaylistItems_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.UpdatePlaylistItems(context.Background(), "", PlaylistItemsRequest{})
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestUpdatePlaylistItems_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"snapshot_id":"newsnap1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	snapshotID, err := client.UpdatePlaylistItems(context.Background(), "pl1", PlaylistItemsRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *snapshotID != "newsnap1" {
		t.Errorf("got snapshot_id %q, want 'newsnap1'", *snapshotID)
	}
}

func TestAddItemsToPlaylist_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.AddItemsToPlaylist(context.Background(), "", AddItemToPlaylistRequest{})
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestAddItemsToPlaylist_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/playlists/pl1/items" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"snapshot_id":"snap2"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	snapshotID, err := client.AddItemsToPlaylist(context.Background(), "pl1", AddItemToPlaylistRequest{URIs: []string{"spotify:track:abc"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *snapshotID != "snap2" {
		t.Errorf("got snapshot_id %q, want 'snap2'", *snapshotID)
	}
}

func TestRemovePlaylistItems_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.RemovePlaylistItems(context.Background(), "", RemovePlaylistItemsRequest{})
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestRemovePlaylistItems_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/playlists/pl1/items" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"snapshot_id":"snap3"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	snapshotID, err := client.RemovePlaylistItems(context.Background(), "pl1", RemovePlaylistItemsRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *snapshotID != "snap3" {
		t.Errorf("got snapshot_id %q, want 'snap3'", *snapshotID)
	}
}

func TestGetCurrentUserPlaylists_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/playlists" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":6,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	playlists, err := client.GetCurrentUserPlaylists(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if playlists.Total != 6 {
		t.Errorf("got Total %d, want 6", playlists.Total)
	}
}

func TestCreatePlaylist_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/me/playlists" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":"newpl1","name":"New Playlist","type":"playlist","snapshot_id":"s1","items":{"href":"","total":0,"limit":20,"offset":0,"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	playlist, err := client.CreatePlaylist(context.Background(), CreatePlaylistRequest{Name: "New Playlist"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if playlist.ID != "newpl1" {
		t.Errorf("got ID %q, want newpl1", playlist.ID)
	}
	if playlist.Name != "New Playlist" {
		t.Errorf("got Name %q, want 'New Playlist'", playlist.Name)
	}
}

func TestGetPlaylistCoverImage_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetPlaylistCoverImage(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestGetPlaylistCoverImage_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/playlists/pl1/images" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"url":"https://example.com/img.jpg","height":300,"width":300}]`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	images, err := client.GetPlaylistCoverImage(context.Background(), "pl1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(images) != 1 {
		t.Fatalf("got %d images, want 1", len(images))
	}
	if images[0].URL != "https://example.com/img.jpg" {
		t.Errorf("got URL %q, want 'https://example.com/img.jpg'", images[0].URL)
	}
}

func TestAddCustomPlaylistCoverImage_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.AddCustomPlaylistCoverImage(context.Background(), "", []byte("data"))
	if err == nil {
		t.Fatal("expected error for empty playlist id")
	}
}

func TestAddCustomPlaylistCoverImage_EmptyBody_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.AddCustomPlaylistCoverImage(context.Background(), "pl1", []byte{})
	if err == nil {
		t.Fatal("expected error for empty body")
	}
	if err.Error() != "body cannot be empty" {
		t.Errorf("got %q, want 'body cannot be empty'", err.Error())
	}
}

func TestAddCustomPlaylistCoverImage_Success(t *testing.T) {
	var capturedContentType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		capturedContentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.AddCustomPlaylistCoverImage(context.Background(), "pl1", []byte("base64imagedata"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedContentType != "image/jpeg" {
		t.Errorf("got Content-Type %q, want 'image/jpeg'", capturedContentType)
	}
}
