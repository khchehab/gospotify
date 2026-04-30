package gospotify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAudiobook_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetAudiobook(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
	if err.Error() != "id cannot be empty" {
		t.Errorf("got %q, want 'id cannot be empty'", err.Error())
	}
}

func TestGetAudiobook_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/audiobooks/ab1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"ab1","name":"Test Audiobook","type":"audiobook","total_chapters":12}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	book, err := client.GetAudiobook(context.Background(), "ab1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if book.ID != "ab1" {
		t.Errorf("got ID %q, want ab1", book.ID)
	}
	if book.Name != "Test Audiobook" {
		t.Errorf("got Name %q, want 'Test Audiobook'", book.Name)
	}
	if book.TotalChapters != 12 {
		t.Errorf("got TotalChapters %d, want 12", book.TotalChapters)
	}
}

func TestGetAudiobook_WithMarket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("market") != "AU" {
			t.Errorf("expected market=AU, got %q", r.URL.Query().Get("market"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"ab1"}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetAudiobook(context.Background(), "ab1", WithMarket("AU"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetAudiobook_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"status":404,"message":"Audiobook not found"}}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	_, err := client.GetAudiobook(context.Background(), "bad-id")
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

func TestGetAudiobookChapters_EmptyID_ReturnsError(t *testing.T) {
	client := newTestClient("http://unused")
	_, err := client.GetAudiobookChapters(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id")
	}
}

func TestGetAudiobookChapters_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/audiobooks/ab1/chapters" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":12,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	chapters, err := client.GetAudiobookChapters(context.Background(), "ab1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chapters.Total != 12 {
		t.Errorf("got Total %d, want 12", chapters.Total)
	}
}

func TestGetUserSavedAudiobooks_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/me/audiobooks" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"href":"","total":2,"limit":20,"offset":0,"items":[]}`))
	}))
	defer srv.Close()

	client := newTestClient(srv.URL)
	books, err := client.GetUserSavedAudiobooks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if books.Total != 2 {
		t.Errorf("got Total %d, want 2", books.Total)
	}
}
