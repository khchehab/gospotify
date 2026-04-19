package gospotify

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

type testPayload struct {
	Name string `json:"name"`
}

func newTestClient(serverURL string) *Client {
	return &Client{
		baseURL:    serverURL,
		httpClient: &http.Client{},
	}
}

func TestGet_200_UnmarshalsResponseBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name":"test-track"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(context.Background(), "", &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "test-track" {
		t.Errorf("got Name %q, want test-track", result.Name)
	}
}

func TestGet_Non200_ReturnsErrorResponse(t *testing.T) {
	body := `{"error":{"status":401,"message":"No token provided"}}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(body))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(context.Background(), "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}

	var errResp *ErrorResponse
	if !errors.As(err, &errResp) {
		t.Fatalf("expected *ErrorResponse, got %T: %v", err, err)
	}
	if errResp.ErrorObject.Status != 401 {
		t.Errorf("got status %d, want 401", errResp.ErrorObject.Status)
	}
	if errResp.ErrorObject.Message != "No token provided" {
		t.Errorf("got message %q, want 'No token provided'", errResp.ErrorObject.Message)
	}
	if err.Error() != "[401] No token provided" {
		t.Errorf("got Error() %q, want '[401] No token provided'", err.Error())
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}

func TestGet_Non200_InvalidJSON_ReturnsUnmarshalError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is not json`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(context.Background(), "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Errorf("expected *json.SyntaxError, got %T: %v", err, err)
	}
	var errResp *ErrorResponse
	if errors.As(err, &errResp) {
		t.Error("error should not be an *ErrorResponse")
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}

func TestGet_200_InvalidJSON_ReturnsUnmarshalError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{bad json`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(context.Background(), "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Errorf("expected *json.SyntaxError, got %T: %v", err, err)
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}

func TestGet_TransportError_ReturnsError(t *testing.T) {
	// Start a server then close it immediately so the port is unreachable.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	closedURL := srv.URL
	srv.Close()

	client := newTestClient(closedURL)
	var result testPayload
	err := client.get(context.Background(), "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}
	var urlErr *url.Error
	if !errors.As(err, &urlErr) {
		t.Errorf("expected *url.Error, got %T: %v", err, err)
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}

func TestGet_CancelledContext_AlreadyCancelled(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name":"should-not-reach"}`))
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before the call

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(ctx, "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled in error chain, got: %v", err)
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}

func TestGet_OptsPassedThrough(t *testing.T) {
	var capturedQuery string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(context.Background(), "/some/endpoint", &result, WithLimit(25), WithOffset(10))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	q, err2 := url.ParseQuery(capturedQuery)
	if err2 != nil {
		t.Fatalf("could not parse query %q: %v", capturedQuery, err2)
	}
	if q.Get("limit") != "25" {
		t.Errorf("limit: got %q, want 25", q.Get("limit"))
	}
	if q.Get("offset") != "10" {
		t.Errorf("offset: got %q, want 10", q.Get("offset"))
	}
}

// ---- buildURL ----

func TestBuildURL_NoOpts(t *testing.T) {
	client := newTestClient("https://api.spotify.com/v1")
	got := client.buildURL("/tracks/1")
	want := "https://api.spotify.com/v1/tracks/1"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBuildURL_WithOpts_NoExistingQuery(t *testing.T) {
	client := newTestClient("https://api.spotify.com/v1")
	got := client.buildURL("/me/top/artists", WithLimit(10))
	parsed, err := url.Parse(got)
	if err != nil {
		t.Fatalf("invalid URL: %v", err)
	}
	if parsed.Query().Get("limit") != "10" {
		t.Errorf("limit: got %q, want 10", parsed.Query().Get("limit"))
	}
	// must have exactly one '?' in the URL
	if strings.Count(got, "?") != 1 {
		t.Errorf("expected exactly one '?' in URL, got: %q", got)
	}
}

func TestBuildURL_WithOpts_ExistingQuery(t *testing.T) {
	client := newTestClient("https://api.spotify.com/v1")
	got := client.buildURL("/me/following?type=artist", WithLimit(5))
	parsed, err := url.Parse(got)
	if err != nil {
		t.Fatalf("invalid URL: %v", err)
	}
	q := parsed.Query()
	if q.Get("type") != "artist" {
		t.Errorf("type: got %q, want artist", q.Get("type"))
	}
	if q.Get("limit") != "5" {
		t.Errorf("limit: got %q, want 5", q.Get("limit"))
	}
	if strings.Count(got, "?") != 1 {
		t.Errorf("expected exactly one '?' in URL, got: %q", got)
	}
}

func TestBuildURL_NoOpts_EndpointWithExistingQuery_Unchanged(t *testing.T) {
	client := newTestClient("https://api.spotify.com/v1")
	got := client.buildURL("/me/following?type=artist")
	want := "https://api.spotify.com/v1/me/following?type=artist"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBuildURL_NoOpts_NoQueryAppended(t *testing.T) {
	client := newTestClient("https://api.spotify.com/v1")
	got := client.buildURL("/me")
	if strings.Contains(got, "?") {
		t.Errorf("expected no '?' in URL for no-opts call, got: %q", got)
	}
}

func TestGet_CancelledContext_MidFlight(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name":"too-late"}`))
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	client := newTestClient(srv.URL)
	var result testPayload
	err := client.get(ctx, "", &result)

	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled in error chain, got: %v", err)
	}
	if result.Name != "" {
		t.Errorf("result should be zero-value, got Name=%q", result.Name)
	}
}
