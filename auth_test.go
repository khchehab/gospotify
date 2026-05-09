package gospotify

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// ---- helpers ----

func freePort(t *testing.T) int {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("could not find free port: %v", err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	_ = l.Close()
	return port
}

func formatPort(port int) string {
	if port == 0 {
		return "0"
	}
	buf := make([]byte, 0, 6)
	digits := 0
	for tmp := port; tmp > 0; tmp /= 10 {
		digits++
	}
	for i := digits - 1; i >= 0; i-- {
		p := 1
		for j := 0; j < i; j++ {
			p *= 10
		}
		buf = append(buf, byte('0'+port/p%10))
	}
	return string(buf)
}

// ---- newCallbackHandler (via httptest) ----

func TestCallbackHandler_ValidCode(t *testing.T) {
	authCh := make(chan authResult, 1)
	var once sync.Once
	h := newCallbackHandler("teststate1234567", authCh, &once)

	req := httptest.NewRequest(http.MethodGet, "/callback?state=teststate1234567&code=auth_code_abc", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want 200", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Authorization successful") {
		t.Errorf("unexpected body: %q", w.Body.String())
	}
	select {
	case res := <-authCh:
		if res.err != nil {
			t.Errorf("unexpected error: %v", res.err)
		}
		if res.code != "auth_code_abc" {
			t.Errorf("got code %q, want auth_code_abc", res.code)
		}
	default:
		t.Error("no result sent to channel")
	}
}

func TestCallbackHandler_StateMismatch(t *testing.T) {
	authCh := make(chan authResult, 1)
	var once sync.Once
	h := newCallbackHandler("correctstate12345", authCh, &once)

	req := httptest.NewRequest(http.MethodGet, "/callback?state=wrongstate&code=somecode", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("got status %d, want 400", w.Code)
	}
	res := <-authCh
	if res.err == nil || res.err.Error() != "invalid state parameter" {
		t.Errorf("got error %v, want 'invalid state parameter'", res.err)
	}
	if res.code != "" {
		t.Errorf("expected empty code, got %q", res.code)
	}
}

func TestCallbackHandler_OAuthError(t *testing.T) {
	authCh := make(chan authResult, 1)
	var once sync.Once
	state := "stateforoautherr"
	h := newCallbackHandler(state, authCh, &once)

	req := httptest.NewRequest(http.MethodGet, "/callback?state="+state+"&error=access_denied", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("got status %d, want 403", w.Code)
	}
	res := <-authCh
	if res.err == nil || res.err.Error() != "oauth error: access_denied" {
		t.Errorf("got error %v, want 'oauth error: access_denied'", res.err)
	}
	if res.code != "" {
		t.Errorf("expected empty code, got %q", res.code)
	}
}

func TestCallbackHandler_MissingCode(t *testing.T) {
	authCh := make(chan authResult, 1)
	var once sync.Once
	state := "stateformisscode"
	h := newCallbackHandler(state, authCh, &once)

	req := httptest.NewRequest(http.MethodGet, "/callback?state="+state, nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("got status %d, want 400", w.Code)
	}
	res := <-authCh
	if res.err == nil || res.err.Error() != "code not found in query parameters" {
		t.Errorf("got error %v, want 'code not found in query parameters'", res.err)
	}
}

func TestCallbackHandler_OnceGuard_SecondRequestIgnored(t *testing.T) {
	authCh := make(chan authResult, 1)
	var once sync.Once
	state := "stateforonceguard"
	h := newCallbackHandler(state, authCh, &once)

	req1 := httptest.NewRequest(http.MethodGet, "/callback?state="+state+"&code=first_code", nil)
	h(httptest.NewRecorder(), req1)

	req2 := httptest.NewRequest(http.MethodGet, "/callback?state="+state+"&code=second_code", nil)
	h(httptest.NewRecorder(), req2)

	if len(authCh) != 1 {
		t.Errorf("channel has %d items, want exactly 1", len(authCh))
	}
	res := <-authCh
	if res.code != "first_code" {
		t.Errorf("got code %q, want first_code", res.code)
	}
}

// ---- waitForAuthResult ----

func TestWaitForAuthResult_ValidCode(t *testing.T) {
	authCh := make(chan authResult, 1)
	authCh <- authResult{code: "valid_code_xyz", err: nil}

	code, err := waitForAuthResult(authCh, 5*time.Second)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if code != "valid_code_xyz" {
		t.Errorf("got code %q, want valid_code_xyz", code)
	}
}

func TestWaitForAuthResult_ChannelError(t *testing.T) {
	authCh := make(chan authResult, 1)
	authCh <- authResult{code: "", err: errors.New("oauth error: access_denied")}

	code, err := waitForAuthResult(authCh, 5*time.Second)
	if code != "" {
		t.Errorf("got code %q, want empty", code)
	}
	if err == nil || err.Error() != "oauth error: access_denied" {
		t.Errorf("got error %v, want 'oauth error: access_denied'", err)
	}
}

func TestWaitForAuthResult_Timeout(t *testing.T) {
	authCh := make(chan authResult, 1)

	start := time.Now()
	code, err := waitForAuthResult(authCh, 20*time.Millisecond)
	elapsed := time.Since(start)

	if code != "" {
		t.Errorf("got code %q, want empty", code)
	}
	if err == nil || err.Error() != "authorization timed out" {
		t.Errorf("got error %v, want 'authorization timed out'", err)
	}
	if elapsed < 20*time.Millisecond {
		t.Errorf("returned too early: elapsed %v", elapsed)
	}
}

// ---- exchangeCode ----

func TestExchangeCode_Success(t *testing.T) {
	cfg := &oauth2.Config{}
	token := &oauth2.Token{
		AccessToken: "tok_abc",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	var capturedCtx context.Context
	exchange := func(ctx context.Context, code string) (*oauth2.Token, error) {
		capturedCtx = ctx
		return token, nil
	}

	ts, err := exchangeCode(cfg, "auth_code_abc", exchange)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts == nil {
		t.Fatal("expected non-nil token source")
	}

	deadline, ok := capturedCtx.Deadline()
	if !ok {
		t.Error("context passed to exchange has no deadline")
	}
	if time.Until(deadline) > 10*time.Second {
		t.Errorf("deadline too far in future: %v", time.Until(deadline))
	}
}

func TestExchangeCode_ExchangeError(t *testing.T) {
	cfg := &oauth2.Config{}
	exchange := func(ctx context.Context, code string) (*oauth2.Token, error) {
		return nil, errors.New("token endpoint unreachable")
	}

	ts, err := exchangeCode(cfg, "some_code", exchange)
	if ts != nil {
		t.Error("expected nil token source on error")
	}
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if !strings.Contains(err.Error(), "error exchanging code for token:") {
		t.Errorf("error missing wrapping prefix, got: %v", err)
	}
}

func TestExchangeCode_ErrorWrapping(t *testing.T) {
	cfg := &oauth2.Config{}
	sentinel := errors.New("sentinel token error")
	exchange := func(ctx context.Context, code string) (*oauth2.Token, error) {
		return nil, sentinel
	}

	_, err := exchangeCode(cfg, "code", exchange)
	if !errors.Is(err, sentinel) {
		t.Errorf("errors.Is failed: sentinel not in error chain, got: %v", err)
	}
}

// ---- shutdownServer ----

func TestShutdownServer_RunningServer(t *testing.T) {
	port := freePort(t)
	srv := &http.Server{
		Addr:    "127.0.0.1:" + formatPort(port),
		Handler: http.NewServeMux(),
	}
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatalf("could not listen: %v", err)
	}
	go srv.Serve(ln) //nolint:errcheck

	a := NewAuthenticator("id", "secret", "http://localhost/callback", nil)
	a.shutdownServer(srv)

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		_, err = net.DialTimeout("tcp", "127.0.0.1:"+formatPort(port), 100*time.Millisecond)
		if err != nil {
			return
		}
	}
	t.Error("expected connection to be refused after shutdown, but it succeeded")
}

func TestShutdownServer_AlreadyStopped(t *testing.T) {
	srv := &http.Server{
		Addr:    "127.0.0.1:0",
		Handler: http.NewServeMux(),
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("shutdownServer panicked: %v", r)
		}
	}()
	a := NewAuthenticator("id", "secret", "http://localhost/callback", nil)
	a.shutdownServer(srv)
}

// ---- scopesToString ----

func TestScopesToString_EmptySlice_ReturnsEmptySlice(t *testing.T) {
	result := scopesToString([]Scope{})
	if len(result) != 0 {
		t.Errorf("got len %d, want 0", len(result))
	}
}

func TestScopesToString_NilSlice_ReturnsEmptySlice(t *testing.T) {
	result := scopesToString(nil)
	if len(result) != 0 {
		t.Errorf("got len %d, want 0", len(result))
	}
}

func TestScopesToString_SingleScope_ReturnsCorrectString(t *testing.T) {
	result := scopesToString([]Scope{ScopeUserReadEmail})
	if len(result) != 1 {
		t.Fatalf("got len %d, want 1", len(result))
	}
	if result[0] != "user-read-email" {
		t.Errorf("got %q, want user-read-email", result[0])
	}
}

func TestScopesToString_MultipleScopes_PreservesOrder(t *testing.T) {
	scopes := []Scope{
		ScopeUserReadEmail,
		ScopeUserReadPrivate,
		ScopePlaylistReadPrivate,
	}
	result := scopesToString(scopes)
	if len(result) != 3 {
		t.Fatalf("got len %d, want 3", len(result))
	}
	expected := []string{"user-read-email", "user-read-private", "playlist-read-private"}
	for i, want := range expected {
		if result[i] != want {
			t.Errorf("result[%d]: got %q, want %q", i, result[i], want)
		}
	}
}

func TestScopesToString_AllStringsAreStrings(t *testing.T) {
	scopes := []Scope{
		ScopeUgcImageUpload,
		ScopeUserModifyPlaybackState,
		ScopeStreaming,
		ScopeUserFollowModify,
		ScopeUserLibraryRead,
	}
	result := scopesToString(scopes)
	if len(result) != len(scopes) {
		t.Fatalf("got len %d, want %d", len(result), len(scopes))
	}
	for i, scope := range scopes {
		if result[i] != string(scope) {
			t.Errorf("result[%d]: got %q, want %q", i, result[i], string(scope))
		}
	}
}
