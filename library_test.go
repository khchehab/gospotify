package gospotify

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveItemsToLibrary_EmptyURIs_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.SaveItemsToLibrary(context.Background(), []string{})
	if err == nil {
		t.Fatal("expected error for empty uris")
	}
	if err.Error() != "uris cannot be empty" {
		t.Errorf("got %q, want 'uris cannot be empty'", err.Error())
	}
}

func TestSaveItemsToLibrary_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/me/library" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("uris") != "spotify:track:abc,spotify:track:def" {
			t.Errorf("unexpected uris: %q", r.URL.Query().Get("uris"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.SaveItemsToLibrary(context.Background(), []string{"spotify:track:abc", "spotify:track:def"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRemoveItemsFromLibrary_EmptyURIs_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	err := client.RemoveItemsFromLibrary(context.Background(), []string{})
	if err == nil {
		t.Fatal("expected error for empty uris")
	}
	if err.Error() != "uris cannot be empty" {
		t.Errorf("got %q, want 'uris cannot be empty'", err.Error())
	}
}

func TestRemoveItemsFromLibrary_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/me/library" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("uris") != "spotify:track:abc" {
			t.Errorf("unexpected uris: %q", r.URL.Query().Get("uris"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	err := client.RemoveItemsFromLibrary(context.Background(), []string{"spotify:track:abc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheckUserSavedItems_EmptyURIs_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.CheckUserSavedItems(context.Background(), []string{})
	if err == nil {
		t.Fatal("expected error for empty uris")
	}
	if err.Error() != "uris cannot be empty" {
		t.Errorf("got %q, want 'uris cannot be empty'", err.Error())
	}
}

func TestCheckUserSavedItems_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/library/contains" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("uris") != "spotify:track:abc,spotify:track:def" {
			t.Errorf("unexpected uris: %q", r.URL.Query().Get("uris"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[true,false]`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	saved, err := client.CheckUserSavedItems(context.Background(), []string{"spotify:track:abc", "spotify:track:def"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(saved) != 2 {
		t.Fatalf("got %d results, want 2", len(saved))
	}
	if !saved[0] {
		t.Error("expected saved[0]=true")
	}
	if saved[1] {
		t.Error("expected saved[1]=false")
	}
}

func TestAppendQueryParams_URIs(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		uris     []string
		want     string
	}{
		{
			name:     "single URI",
			endpoint: "/me/library",
			uris:     []string{"spotify:track:abc"},
			want:     "/me/library?uris=spotify%3Atrack%3Aabc",
		},
		{
			name:     "multiple URIs",
			endpoint: "/me/library/contains",
			uris:     []string{"spotify:track:abc", "spotify:track:def"},
			want:     "/me/library/contains?uris=spotify%3Atrack%3Aabc%2Cspotify%3Atrack%3Adef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendQueryParams(tt.endpoint, requiredQueryParam{key: "uris", value: tt.uris})
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
