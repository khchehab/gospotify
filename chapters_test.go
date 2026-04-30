package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChapter_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetChapter(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetChapter_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chapters/ch1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"ch1","name":"Chapter 1","type":"chapter","duration_ms":3600000,"chapter_number":1}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	chapter, err := client.GetChapter(context.Background(), "ch1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chapter.ID != "ch1" {
		t.Errorf("got ID %q, want ch1", chapter.ID)
	}
	if chapter.Name != "Chapter 1" {
		t.Errorf("got Name %q, want 'Chapter 1'", chapter.Name)
	}
}

func TestGetChapter_WithMarket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("market") != "CA" {
			t.Errorf("expected market=CA, got %q", r.URL.Query().Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"ch1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetChapter(context.Background(), "ch1", WithMarket("CA"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetChapter_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"status":404,"message":"Chapter not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetChapter(context.Background(), "bad-id")
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
