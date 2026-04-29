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

// ---- requireNonEmpty ----

func TestRequireNonEmpty_EmptyString_ReturnsError(t *testing.T) {
	err := requireNonEmpty("my field", "")
	if err == nil {
		t.Fatal("expected non-nil error for empty string")
	}
	want := "my field cannot be empty"
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestRequireNonEmpty_NonEmptyString_ReturnsNil(t *testing.T) {
	err := requireNonEmpty("my field", "value")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRequireNonEmpty_FieldNameInError(t *testing.T) {
	cases := []struct {
		field string
		want  string
	}{
		{"id", "id cannot be empty"},
		{"playlist id", "playlist id cannot be empty"},
		{"query", "query cannot be empty"},
	}
	for _, tc := range cases {
		t.Run(tc.field, func(t *testing.T) {
			err := requireNonEmpty(tc.field, "")
			if err == nil {
				t.Fatalf("expected error for field %q", tc.field)
			}
			if err.Error() != tc.want {
				t.Errorf("got %q, want %q", err.Error(), tc.want)
			}
		})
	}
}

// ---- requireNonEmptyArray ----

func TestRequireNonEmptyArray_EmptySlice_ReturnsError(t *testing.T) {
	err := requireNonEmptyArray("ids", []string{})
	if err == nil {
		t.Fatal("expected non-nil error for empty slice")
	}
	want := "ids cannot be empty"
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestRequireNonEmptyArray_NilSlice_ReturnsError(t *testing.T) {
	err := requireNonEmptyArray("ids", []string(nil))
	if err == nil {
		t.Fatal("expected non-nil error for nil slice")
	}
}

func TestRequireNonEmptyArray_NonEmptySlice_ReturnsNil(t *testing.T) {
	err := requireNonEmptyArray("ids", []string{"a", "b"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRequireNonEmptyArray_IntSlice_EmptyReturnsError(t *testing.T) {
	err := requireNonEmptyArray("values", []int{})
	if err == nil {
		t.Fatal("expected non-nil error for empty int slice")
	}
}

func TestRequireNonEmptyArray_IntSlice_NonEmptyReturnsNil(t *testing.T) {
	err := requireNonEmptyArray("values", []int{1, 2, 3})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// ---- requiredQueryParam.String ----

func TestRequiredQueryParam_String_StringValue(t *testing.T) {
	p := requiredQueryParam{key: "market", value: "US"}
	got := p.String()
	want := "market=US"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequiredQueryParam_String_StringValueWithSpecialChars(t *testing.T) {
	p := requiredQueryParam{key: "q", value: "hello world"}
	got := p.String()
	want := "q=hello+world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequiredQueryParam_String_StringSliceValue(t *testing.T) {
	p := requiredQueryParam{key: "ids", value: []string{"a", "b", "c"}}
	got := p.String()
	want := "ids=a%2Cb%2Cc"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequiredQueryParam_String_BoolTrueValue(t *testing.T) {
	p := requiredQueryParam{key: "state", value: true}
	got := p.String()
	want := "state=true"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequiredQueryParam_String_BoolFalseValue(t *testing.T) {
	p := requiredQueryParam{key: "state", value: false}
	got := p.String()
	want := "state=false"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRequiredQueryParam_String_IntValue(t *testing.T) {
	cases := []struct {
		name  string
		value any
		want  string
	}{
		{"int", 42, "n=42"},
		{"int8", int8(8), "n=8"},
		{"int16", int16(16), "n=16"},
		{"int32", int32(32), "n=32"},
		{"int64", int64(64), "n=64"},
		{"uint", uint(10), "n=10"},
		{"uint8", uint8(8), "n=8"},
		{"uint16", uint16(16), "n=16"},
		{"uint32", uint32(32), "n=32"},
		{"uint64", uint64(64), "n=64"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := requiredQueryParam{key: "n", value: tc.value}
			got := p.String()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// ---- appendQueryParams ----

func TestAppendQueryParams_NoParams_ReturnsEndpoint(t *testing.T) {
	got := appendQueryParams("/me/player")
	want := "/me/player"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_SingleParam_AppendsCorrectly(t *testing.T) {
	got := appendQueryParams("/tracks", requiredQueryParam{key: "market", value: "US"})
	want := "/tracks?market=US"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_MultipleParams_JoinedWithAmpersand(t *testing.T) {
	got := appendQueryParams("/search",
		requiredQueryParam{key: "q", value: "beatles"},
		requiredQueryParam{key: "limit", value: 10},
	)
	want := "/search?q=beatles&limit=10"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_SpecialCharsInValue_AreEncoded(t *testing.T) {
	got := appendQueryParams("/me/player/queue",
		requiredQueryParam{key: "uri", value: "spotify:track:abc"},
	)
	want := "/me/player/queue?uri=spotify%3Atrack%3Aabc"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppendQueryParams_SliceValue_JoinedByComma(t *testing.T) {
	got := appendQueryParams("/me/library",
		requiredQueryParam{key: "uris", value: []string{"spotify:track:a", "spotify:track:b"}},
	)
	want := "/me/library?uris=spotify%3Atrack%3Aa%2Cspotify%3Atrack%3Ab"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
