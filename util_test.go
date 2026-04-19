package gospotify

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// ---- RandomString ----

func TestRandomString_Length(t *testing.T) {
	cases := []struct {
		name string
		n    int
	}{
		{"zero", 0},
		{"one", 1},
		{"typical_16", 16},
		{"large_10000", 10000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RandomString(tc.n)
			if len(got) != tc.n {
				t.Errorf("got len %d, want %d", len(got), tc.n)
			}
		})
	}
}

func TestRandomString_AllCharsInAllowedSet(t *testing.T) {
	result := RandomString(10_000)
	for i, ch := range result {
		if !strings.ContainsRune(allowedChars, ch) {
			t.Errorf("char %q at index %d is not in allowedChars", ch, i)
			return
		}
	}
}

func TestRandomString_NoDisallowedChars(t *testing.T) {
	result := RandomString(1_000)
	disallowed := "!@#$%^&*()_+-=[]{}|;':\",./<>? \t\n\r"
	for _, ch := range disallowed {
		if strings.ContainsRune(result, ch) {
			t.Errorf("result contains disallowed char %q", ch)
		}
	}
}

func TestRandomString_NoPanicOnIndexBounds(t *testing.T) {
	// 100k iterations stress-tests that r.IntN(62) never returns >= 62.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RandomString panicked: %v", r)
		}
	}()
	RandomString(100_000)
}

func TestRandomString_Uniqueness_TwoCalls(t *testing.T) {
	a := RandomString(32)
	b := RandomString(32)
	if a == b {
		t.Errorf("two successive RandomString(32) calls returned the same value: %q", a)
	}
}

func TestRandomString_Uniqueness_Batch(t *testing.T) {
	const n = 1000
	seen := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		s := RandomString(16)
		if _, exists := seen[s]; exists {
			t.Errorf("duplicate string produced at iteration %d: %q", i, s)
			return
		}
		seen[s] = struct{}{}
	}
}

func TestRandomString_ValidUTF8(t *testing.T) {
	result := RandomString(100)
	if !utf8.ValidString(result) {
		t.Error("RandomString result is not valid UTF-8")
	}
}

// ---- browserCommand (extracted helper — all platforms testable) ----

func TestBrowserCommand_Darwin(t *testing.T) {
	url := "https://example.com"
	cmd, args, err := browserCommand("darwin", url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd != "open" {
		t.Errorf("got cmd %q, want open", cmd)
	}
	if len(args) != 1 || args[0] != url {
		t.Errorf("got args %v, want [%q]", args, url)
	}
}

func TestBrowserCommand_Windows(t *testing.T) {
	url := "https://example.com"
	cmd, args, err := browserCommand("windows", url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd != "rundll32" {
		t.Errorf("got cmd %q, want rundll32", cmd)
	}
	if len(args) != 2 || args[0] != "url.dll,FileProtocolHandler" || args[1] != url {
		t.Errorf("got args %v", args)
	}
}

func TestBrowserCommand_Linux(t *testing.T) {
	url := "https://example.com"
	cmd, args, err := browserCommand("linux", url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd != "xdg-open" {
		t.Errorf("got cmd %q, want xdg-open", cmd)
	}
	if len(args) != 1 || args[0] != url {
		t.Errorf("got args %v, want [%q]", args, url)
	}
}

func TestBrowserCommand_UnsupportedPlatform_ReturnsError(t *testing.T) {
	_, _, err := browserCommand("freebsd", "https://example.com")
	if err == nil {
		t.Fatal("expected non-nil error for unsupported platform")
	}
}

func TestBrowserCommand_UnsupportedPlatform_ErrorMessageFormat(t *testing.T) {
	goos := "freebsd"
	_, _, err := browserCommand(goos, "https://example.com")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	want := "unsupported platform: " + goos
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestBrowserCommand_UnsupportedPlatform_NoCommandReturned(t *testing.T) {
	cmd, args, err := browserCommand("plan9", "https://example.com")
	if err == nil {
		t.Fatal("expected error")
	}
	if cmd != "" {
		t.Errorf("expected empty cmd, got %q", cmd)
	}
	if args != nil {
		t.Errorf("expected nil args, got %v", args)
	}
}
