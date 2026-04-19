package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// ---- shared test helpers ----

type capturedRequest struct {
	path     string
	rawQuery string
	query    url.Values
}

func newUsersTestServer(t *testing.T, statusCode int, body string) (*httptest.Server, *capturedRequest) {
	t.Helper()
	captured := &capturedRequest{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured.path = r.URL.Path
		captured.rawQuery = r.URL.RawQuery
		captured.query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(body))
	}))
	return srv, captured
}

func newUsersClient(serverURL string) *Client {
	return &Client{
		baseURL:    serverURL,
		httpClient: &http.Client{},
	}
}

func assertQueryParam(t *testing.T, query url.Values, key, want string) {
	t.Helper()
	got := query.Get(key)
	if got != want {
		t.Errorf("query param %q: got %q, want %q", key, got, want)
	}
}

func assertQueryParamPresent(t *testing.T, query url.Values, key string) {
	t.Helper()
	if _, ok := query[key]; !ok {
		t.Errorf("query param %q not present", key)
	}
}

// ---- GetCurrentUserProfile ----

func TestGetCurrentUserProfile_HappyPath(t *testing.T) {
	body := `{"id":"user123","display_name":"Test User","type":"user","uri":"spotify:user:user123","href":"https://api.spotify.com/v1/users/user123"}`
	srv, req := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if profile == nil {
		t.Fatal("expected non-nil profile")
	}
	if req.path != "/me" {
		t.Errorf("got path %q, want /me", req.path)
	}
	if req.rawQuery != "" {
		t.Errorf("expected no query string, got %q", req.rawQuery)
	}
	if profile.ID != "user123" {
		t.Errorf("got ID %q, want user123", profile.ID)
	}
	if profile.DisplayName == nil || *profile.DisplayName != "Test User" {
		t.Errorf("got DisplayName %v, want 'Test User'", profile.DisplayName)
	}
}

func TestGetCurrentUserProfile_Non200_ReturnsError(t *testing.T) {
	body := `{"error":{"status":401,"message":"Unauthorized"}}`
	srv, _ := newUsersTestServer(t, http.StatusUnauthorized, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(context.Background())

	if profile != nil {
		t.Error("expected nil profile on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetCurrentUserProfile_500_ReturnsError(t *testing.T) {
	body := `{"error":{"status":500,"message":"Internal server error"}}`
	srv, _ := newUsersTestServer(t, http.StatusInternalServerError, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(context.Background())

	if profile != nil {
		t.Error("expected nil profile on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetCurrentUserProfile_InvalidJSON_ReturnsError(t *testing.T) {
	srv, _ := newUsersTestServer(t, http.StatusOK, `{invalid json`)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(context.Background())

	if profile != nil {
		t.Error("expected nil profile on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetCurrentUserProfile_CancelledContext(t *testing.T) {
	srv, _ := newUsersTestServer(t, http.StatusOK, `{}`)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	client := newUsersClient(srv.URL)
	profile, err := client.GetCurrentUserProfile(ctx)

	if profile != nil {
		t.Error("expected nil profile on cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

// ---- GetUserTopArtists ----

const topArtistsBody = `{"href":"https://api.spotify.com/v1/me/top/artists","limit":20,"offset":0,"total":2,"items":[{"id":"a1","name":"Artist One"}]}`

func TestGetUserTopArtists_HappyPath_NoOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if req.path != "/me/top/artists" {
		t.Errorf("got path %q, want /me/top/artists", req.path)
	}
	if result.Total != 2 {
		t.Errorf("got Total %d, want 2", result.Total)
	}
	if len(result.Items) != 1 || result.Items[0].Name != "Artist One" {
		t.Errorf("unexpected items: %+v", result.Items)
	}
}

func TestGetUserTopArtists_WithLimit(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithLimit(10))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "limit", "10")
}

func TestGetUserTopArtists_WithOffset(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithOffset(5))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "offset", "5")
}

func TestGetUserTopArtists_WithTimeRange(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithTimeRange(LongTerm))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "time_range", "long_term")
}

func TestGetUserTopArtists_MultipleOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithLimit(10), WithOffset(20))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "limit", "10")
	assertQueryParam(t, req.query, "offset", "20")
}

func TestGetUserTopArtists_Non200_ReturnsError(t *testing.T) {
	body := `{"error":{"status":403,"message":"Forbidden"}}`
	srv, _ := newUsersTestServer(t, http.StatusForbidden, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopArtists(context.Background())

	if result != nil {
		t.Error("expected nil result on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetUserTopArtists_PaginationFields(t *testing.T) {
	nextURL := "https://api.spotify.com/v1/me/top/artists?offset=20"
	body := `{"href":"...","limit":20,"offset":0,"total":50,"next":"` + nextURL + `","previous":null,"items":[]}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Next == nil || *result.Next != nextURL {
		t.Errorf("got Next %v, want %q", result.Next, nextURL)
	}
	if result.Previous != nil {
		t.Errorf("expected nil Previous, got %v", result.Previous)
	}
}

func TestGetUserTopArtists_EmptyItems(t *testing.T) {
	body := `{"href":"...","limit":20,"offset":0,"total":0,"items":[]}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 0 {
		t.Errorf("got Total %d, want 0", result.Total)
	}
	if result.Items == nil {
		t.Error("expected non-nil Items slice for empty list")
	}
	if len(result.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(result.Items))
	}
}

// ---- GetUserTopTracks ----

const topTracksBody = `{"href":"https://api.spotify.com/v1/me/top/tracks","limit":20,"offset":0,"total":1,"items":[{"id":"t1","name":"Track One"}]}`

func TestGetUserTopTracks_HappyPath_NoOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topTracksBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopTracks(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if req.path != "/me/top/tracks" {
		t.Errorf("got path %q, want /me/top/tracks", req.path)
	}
	if result.Total != 1 {
		t.Errorf("got Total %d, want 1", result.Total)
	}
	if len(result.Items) != 1 || result.Items[0].Name != "Track One" {
		t.Errorf("unexpected items: %+v", result.Items)
	}
}

func TestGetUserTopTracks_WithLimit(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topTracksBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopTracks(context.Background(), WithLimit(10))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "limit", "10")
}

func TestGetUserTopTracks_WithOffset(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topTracksBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopTracks(context.Background(), WithOffset(5))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "offset", "5")
}

func TestGetUserTopTracks_MultipleOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topTracksBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopTracks(context.Background(), WithLimit(10), WithOffset(20))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "limit", "10")
	assertQueryParam(t, req.query, "offset", "20")
}

func TestGetUserTopTracks_Non200_ReturnsError(t *testing.T) {
	body := `{"error":{"status":403,"message":"Forbidden"}}`
	srv, _ := newUsersTestServer(t, http.StatusForbidden, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopTracks(context.Background())

	if result != nil {
		t.Error("expected nil result on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetUserTopTracks_PaginationFields(t *testing.T) {
	prevURL := "https://api.spotify.com/v1/me/top/tracks?offset=0"
	nextURL := "https://api.spotify.com/v1/me/top/tracks?offset=40"
	body := `{"href":"...","limit":20,"offset":20,"total":50,"next":"` + nextURL + `","previous":"` + prevURL + `","items":[]}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopTracks(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Next == nil || *result.Next != nextURL {
		t.Errorf("got Next %v, want %q", result.Next, nextURL)
	}
	if result.Previous == nil || *result.Previous != prevURL {
		t.Errorf("got Previous %v, want %q", result.Previous, prevURL)
	}
}

func TestGetUserTopTracks_EmptyItems(t *testing.T) {
	body := `{"href":"...","limit":20,"offset":0,"total":0,"items":[]}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetUserTopTracks(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil Items slice for empty list")
	}
	if len(result.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(result.Items))
	}
}

// ---- GetFollowedArtists ----

const followedArtistsBody = `{"artists":{"href":"https://api.spotify.com/v1/me/following","limit":20,"total":1,"cursors":{"after":"abc"},"items":[{"id":"a1","name":"Artist One"}]}}`

func TestGetFollowedArtists_HappyPath_NoOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetFollowedArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if req.path != "/me/following" {
		t.Errorf("got path %q, want /me/following", req.path)
	}
	assertQueryParam(t, req.query, "type", "artist")
	if result.Total != 1 {
		t.Errorf("got Total %d, want 1", result.Total)
	}
	if len(result.Items) != 1 || result.Items[0].Name != "Artist One" {
		t.Errorf("unexpected items: %+v", result.Items)
	}
	if result.Cursors.After == nil || *result.Cursors.After != "abc" {
		t.Errorf("got Cursors.After %v, want 'abc'", result.Cursors.After)
	}
}

func TestGetFollowedArtists_NoOptions_NoExtraAmpersand(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.rawQuery != "type=artist" {
		t.Errorf("got raw query %q, want exactly 'type=artist'", req.rawQuery)
	}
	if strings.Contains(req.rawQuery, "&&") || strings.HasSuffix(req.rawQuery, "&") {
		t.Errorf("malformed query string: %q", req.rawQuery)
	}
}

func TestGetFollowedArtists_WithLimit(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background(), WithLimit(5))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "type", "artist")
	assertQueryParam(t, req.query, "limit", "5")
}

func TestGetFollowedArtists_WithAfter(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background(), WithAfter("cursor_xyz"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "type", "artist")
	assertQueryParam(t, req.query, "after", "cursor_xyz")
}

func TestGetFollowedArtists_MultipleOptions(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background(), WithLimit(10), WithAfter("cursor_abc"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "type", "artist")
	assertQueryParam(t, req.query, "limit", "10")
	assertQueryParam(t, req.query, "after", "cursor_abc")
}

func TestGetFollowedArtists_Non200_ReturnsError(t *testing.T) {
	body := `{"error":{"status":401,"message":"Unauthorized"}}`
	srv, _ := newUsersTestServer(t, http.StatusUnauthorized, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetFollowedArtists(context.Background())

	if result != nil {
		t.Error("expected nil result on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestGetFollowedArtists_ResponseUnwrapping(t *testing.T) {
	// The method must return &response.Artists, not the raw followedArtistsResponse.
	// Verify by checking that Items maps to artists.items in JSON, not a top-level key.
	body := `{"artists":{"href":"...","limit":20,"total":2,"cursors":{},"items":[{"id":"b1","name":"Band One"},{"id":"b2","name":"Band Two"}]}}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetFollowedArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("got Total %d, want 2", result.Total)
	}
	if len(result.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(result.Items))
	}
	if result.Items[0].Name != "Band One" || result.Items[1].Name != "Band Two" {
		t.Errorf("unexpected item names: %q %q", result.Items[0].Name, result.Items[1].Name)
	}
}

func TestGetFollowedArtists_EmptyItems(t *testing.T) {
	body := `{"artists":{"href":"...","limit":20,"total":0,"cursors":{},"items":[]}}`
	srv, _ := newUsersTestServer(t, http.StatusOK, body)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	result, err := client.GetFollowedArtists(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil Items slice for empty list")
	}
	if len(result.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(result.Items))
	}
}

func TestGetUserTopArtists_WithMarket(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopArtists(context.Background(), WithMarket("US"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "market", "US")
}

func TestGetUserTopTracks_WithMarket(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, topTracksBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetUserTopTracks(context.Background(), WithMarket("GB"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "market", "GB")
}

func TestGetFollowedArtists_WithMarket(t *testing.T) {
	srv, req := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	client := newUsersClient(srv.URL)
	_, err := client.GetFollowedArtists(context.Background(), WithMarket("DE"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertQueryParam(t, req.query, "type", "artist")
	assertQueryParam(t, req.query, "market", "DE")
	if strings.Count(req.rawQuery, "?") != 0 {
		t.Errorf("raw query must not contain '?', got: %q", req.rawQuery)
	}
}

func TestGetFollowedArtists_CancelledContext(t *testing.T) {
	srv, _ := newUsersTestServer(t, http.StatusOK, followedArtistsBody)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	client := newUsersClient(srv.URL)
	result, err := client.GetFollowedArtists(ctx)

	if result != nil {
		t.Error("expected nil result on cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}
