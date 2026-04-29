package gospotify

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
	if snapshotID != "newsnap1" {
		t.Errorf("got snapshot_id %q, want 'newsnap1'", snapshotID)
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
	if snapshotID != "snap2" {
		t.Errorf("got snapshot_id %q, want 'snap2'", snapshotID)
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
	if snapshotID != "snap3" {
		t.Errorf("got snapshot_id %q, want 'snap3'", snapshotID)
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
	if err.Error() != "image data cannot be empty" {
		t.Errorf("got %q, want 'image data cannot be empty'", err.Error())
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

// ---- PlaylistTrackObject.UnmarshalJSON ----
//
// Note: UnmarshalJSON calls unmarshalTrackOrEpisode(data) where data is the
// outer PlaylistTrackObject JSON (not the nested item). The helper inspects
// the top-level "type" field of data to determine whether to unmarshal a track
// or episode. The "item" field is only used to decide whether to attempt
// unmarshal at all (nil raw.Item means absent field → skip).

func TestPlaylistTrackObject_UnmarshalJSON_TrackItem(t *testing.T) {
	// The outer JSON must carry a top-level "type":"track" so that
	// unmarshalTrackOrEpisode(data) finds the correct type.
	data := `{
		"added_at": "2021-01-01T00:00:00Z",
		"is_local": false,
		"type": "track",
		"id": "t1",
		"name": "My Track",
		"duration_ms": 180000,
		"item": {"type": "track", "id": "t1"}
	}`
	var p PlaylistTrackObject
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

func TestPlaylistTrackObject_UnmarshalJSON_EpisodeItem(t *testing.T) {
	data := `{
		"added_at": "2021-06-15T12:00:00Z",
		"is_local": false,
		"type": "episode",
		"id": "e1",
		"name": "My Episode",
		"duration_ms": 3600000,
		"item": {"type": "episode", "id": "e1"}
	}`
	var p PlaylistTrackObject
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

func TestPlaylistTrackObject_UnmarshalJSON_NoItem_BothNil(t *testing.T) {
	// When the "item" field is absent, raw.Item is nil and we return early
	// without attempting to unmarshal a track or episode.
	data := `{"added_at": "2021-01-01T00:00:00Z", "is_local": false}`
	var p PlaylistTrackObject
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

func TestPlaylistTrackObject_UnmarshalJSON_UnknownType_ReturnsError(t *testing.T) {
	data := `{
		"added_at": "2021-01-01T00:00:00Z",
		"is_local": false,
		"type": "weird_type",
		"item": {"type": "weird_type", "id": "x1"}
	}`
	var p PlaylistTrackObject
	err := json.Unmarshal([]byte(data), &p)
	if err == nil {
		t.Fatal("expected error for unknown item type")
	}
	if !strings.Contains(err.Error(), "unknown item type") {
		t.Errorf("expected error to contain 'unknown item type', got %q", err.Error())
	}
}

// ---- PlaylistObject fields via playlistBase embedding ----

func TestPlaylistObject_playlistBase_FieldsUnmarshalCorrectly(t *testing.T) {
	data := `{
		"id": "pl99",
		"name": "Embedded Playlist",
		"collaborative": true,
		"snapshot_id": "snap99",
		"public": true,
		"owner": {"id": "user1", "href": "https://api.spotify.com/v1/users/user1", "type": "user", "uri": "spotify:user:user1", "external_urls": {"spotify": "https://open.spotify.com/user/user1"}},
		"external_urls": {"spotify": "https://open.spotify.com/playlist/pl99"},
		"href": "https://api.spotify.com/v1/playlists/pl99",
		"images": [],
		"items": {"href": "", "total": 0, "limit": 20, "offset": 0, "items": []}
	}`
	var pl PlaylistObject
	if err := json.Unmarshal([]byte(data), &pl); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if pl.ID != "pl99" {
		t.Errorf("got ID %q, want pl99", pl.ID)
	}
	if pl.Name != "Embedded Playlist" {
		t.Errorf("got Name %q, want 'Embedded Playlist'", pl.Name)
	}
	if !pl.Collaborative {
		t.Error("expected Collaborative=true")
	}
	if pl.SnapshotID != "snap99" {
		t.Errorf("got SnapshotID %q, want snap99", pl.SnapshotID)
	}
	if pl.Public == nil || !*pl.Public {
		t.Error("expected Public=true")
	}
	if pl.Owner.ID != "user1" {
		t.Errorf("got Owner.ID %q, want user1", pl.Owner.ID)
	}
}

func TestSimplifiedPlaylistObject_playlistBase_FieldsUnmarshalCorrectly(t *testing.T) {
	data := `{
		"id": "spl1",
		"name": "Simplified Playlist",
		"collaborative": false,
		"snapshot_id": "ssnap1",
		"public": false,
		"owner": {"id": "user2", "href": "https://api.spotify.com/v1/users/user2", "type": "user", "uri": "spotify:user:user2", "external_urls": {"spotify": ""}},
		"external_urls": {"spotify": "https://open.spotify.com/playlist/spl1"},
		"href": "https://api.spotify.com/v1/playlists/spl1",
		"images": [],
		"tracks": {"href": "https://api.spotify.com/v1/playlists/spl1/tracks", "total": 5}
	}`
	var pl SimplifiedPlaylistObject
	if err := json.Unmarshal([]byte(data), &pl); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if pl.ID != "spl1" {
		t.Errorf("got ID %q, want spl1", pl.ID)
	}
	if pl.Name != "Simplified Playlist" {
		t.Errorf("got Name %q, want 'Simplified Playlist'", pl.Name)
	}
	if pl.Collaborative {
		t.Error("expected Collaborative=false")
	}
	if pl.Tracks == nil || pl.Tracks.Total != 5 {
		t.Errorf("expected Tracks.Total=5, got %v", pl.Tracks)
	}
}
