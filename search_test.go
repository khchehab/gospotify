package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchForItem_EmptyQuery_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.SearchForItem(context.Background(), "", []ItemType{ItemTypeTrack})
	if err == nil {
		t.Fatal("expected error for empty query")
	}
	if err.Error() != "query cannot be empty" {
		t.Errorf("got %q, want 'query cannot be empty'", err.Error())
	}
}

func TestSearchForItem_EmptyTypes_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.SearchForItem(context.Background(), "test", []ItemType{})
	if err == nil {
		t.Fatal("expected error for empty types")
	}
	if err.Error() != "types cannot be empty" {
		t.Errorf("got %q, want 'types cannot be empty'", err.Error())
	}
}

func TestSearchForItem_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("q") != "test query" {
			t.Errorf("expected q='test query', got %q", r.URL.Query().Get("q"))
		}
		if r.URL.Query().Get("types") != "track,artist" {
			t.Errorf("expected types=track,artist, got %q", r.URL.Query().Get("types"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"tracks":{"href":"","total":1,"limit":20,"offset":0,"items":[]},
			"artists":{"href":"","total":2,"limit":20,"offset":0,"items":[]}
		}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	result, err := client.SearchForItem(context.Background(), "test query", []ItemType{ItemTypeTrack, ItemTypeArtist})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Tracks.Total != 1 {
		t.Errorf("got Tracks.Total %d, want 1", result.Tracks.Total)
	}
	if result.Artists.Total != 2 {
		t.Errorf("got Artists.Total %d, want 2", result.Artists.Total)
	}
}

func TestSearchForItem_SpecialCharsInQuery_AreEncoded(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q != "hello world" {
			t.Errorf("expected decoded q='hello world', got %q", q)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"tracks":{"href":"","total":0,"limit":20,"offset":0,"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.SearchForItem(context.Background(), "hello world", []ItemType{ItemTypeTrack})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSearchForItem_WithOptions(t *testing.T) {
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
		w.Write([]byte(`{"tracks":{"href":"","total":0,"limit":10,"offset":0,"items":[]}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.SearchForItem(context.Background(), "beatles", []ItemType{ItemTypeTrack}, WithLimit(10), WithMarket("US"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSearchForItem_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":{"status":401,"message":"No token provided"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.SearchForItem(context.Background(), "test", []ItemType{ItemTypeTrack})
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

func TestAppendQueryParams_Search(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		q        string
		types    []string
		want     string
	}{
		{
			name:     "single type",
			endpoint: "/search",
			q:        "beatles",
			types:    []string{"track"},
			want:     "/search?q=beatles&types=track",
		},
		{
			name:     "multiple types",
			endpoint: "/search",
			q:        "adele",
			types:    []string{"track", "artist"},
			want:     "/search?q=adele&types=track%2Cartist",
		},
		{
			name:     "query with space",
			endpoint: "/search",
			q:        "hello world",
			types:    []string{"track"},
			want:     "/search?q=hello+world&types=track",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendQueryParams(tt.endpoint,
				requiredQueryParam{key: "q", value: tt.q},
				requiredQueryParam{key: "types", value: tt.types},
			)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
