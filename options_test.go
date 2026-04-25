package gospotify

import (
	"math"
	"testing"
)

// helpers

func ptrTimeRange(t TimeRange) *TimeRange { return &t }
func ptrInt(i int) *int                  { return &i }
func ptrString(s string) *string         { return &s }

// ---- WithTimeRange ----

func TestWithTimeRange(t *testing.T) {
	cases := []struct {
		name  string
		input TimeRange
		want  TimeRange
	}{
		{"short_term", ShortTerm, "short_term"},
		{"medium_term", MediumTerm, "medium_term"},
		{"long_term", LongTerm, "long_term"},
		{"arbitrary", TimeRange("all_time"), "all_time"},
		{"empty", TimeRange(""), ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithTimeRange(tc.input)(&p)
			if p.timeRange == nil {
				t.Fatal("timeRange is nil")
			}
			if *p.timeRange != tc.want {
				t.Errorf("got %q, want %q", *p.timeRange, tc.want)
			}
		})
	}
}

func TestWithTimeRange_PointerIndependence(t *testing.T) {
	val := ShortTerm
	p := queryParameters{}
	WithTimeRange(val)(&p)
	val = LongTerm
	if *p.timeRange != ShortTerm {
		t.Errorf("timeRange mutated after caller change: got %q", *p.timeRange)
	}
}

// ---- WithLimit ----

func TestWithLimit(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"lower_boundary", 1, 1},
		{"midpoint", 25, 25},
		{"upper_boundary", 50, 50},
		{"zero_clamped_to_1", 0, 1},
		{"negative_clamped_to_1", -100, 1},
		{"51_clamped_to_50", 51, 50},
		{"large_positive_clamped_to_50", 10000, 50},
		{"min_int", math.MinInt, 1},
		{"max_int", math.MaxInt, 50},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithLimit(tc.input)(&p)
			if p.limit == nil {
				t.Fatal("limit is nil")
			}
			if *p.limit != tc.want {
				t.Errorf("got %d, want %d", *p.limit, tc.want)
			}
		})
	}
}

// ---- WithOffset ----

func TestWithOffset(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 10, 10},
		{"large_positive", 100000, 100000},
		{"negative_one_clamped", -1, 0},
		{"large_negative_clamped", -500, 0},
		{"min_int_clamped", math.MinInt, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithOffset(tc.input)(&p)
			if p.offset == nil {
				t.Fatal("offset is nil")
			}
			if *p.offset != tc.want {
				t.Errorf("got %d, want %d", *p.offset, tc.want)
			}
		})
	}
}

// ---- WithAfter ----

func TestWithAfter(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"normal_cursor", "cursor_abc123", "cursor_abc123"},
		{"empty_string", "", ""},
		{"special_chars", "abc/def?x=1&y=2 z", "abc/def?x=1&y=2 z"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithAfter(tc.input)(&p)
			if p.after == nil {
				t.Fatal("after is nil")
			}
			if *p.after != tc.want {
				t.Errorf("got %q, want %q", *p.after, tc.want)
			}
		})
	}
}

// ---- WithMarket ----

func TestWithMarket(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"us", "US", "US"},
		{"gb", "GB", "GB"},
		{"empty_string", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithMarket(tc.input)(&p)
			if p.market == nil {
				t.Fatal("market is nil")
			}
			if *p.market != tc.want {
				t.Errorf("got %q, want %q", *p.market, tc.want)
			}
		})
	}
}

func TestWithMarket_PointerIndependence(t *testing.T) {
	val := "US"
	p := queryParameters{}
	WithMarket(val)(&p)
	val = "GB"
	if *p.market != "US" {
		t.Errorf("market mutated after caller change: got %q", *p.market)
	}
}

// ---- WithIncludeGroups ----

func TestWithIncludeGroups(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"single", []string{"album"}, []string{"album"}},
		{"multiple", []string{"album", "single", "appears_on"}, []string{"album", "single", "appears_on"}},
		{"empty_slice", []string{}, []string{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithIncludeGroups(tc.input...)(&p)
			if len(p.includeGroups) != len(tc.want) {
				t.Fatalf("got len %d, want %d", len(p.includeGroups), len(tc.want))
			}
			for i, v := range tc.want {
				if p.includeGroups[i] != v {
					t.Errorf("index %d: got %q, want %q", i, p.includeGroups[i], v)
				}
			}
		})
	}
}

// ---- WithIncludeExternal ----

func TestWithIncludeExternal(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"audio", "audio", "audio"},
		{"empty_string", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithIncludeExternal(tc.input)(&p)
			if p.includeExternal == nil {
				t.Fatal("includeExternal is nil")
			}
			if *p.includeExternal != tc.want {
				t.Errorf("got %q, want %q", *p.includeExternal, tc.want)
			}
		})
	}
}

// ---- WithFields ----

func TestWithFields(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"single_field", "name", "name"},
		{"dot_notation", "tracks.items(track(name,id))", "tracks.items(track(name,id))"},
		{"exclusion", "!available_markets", "!available_markets"},
		{"empty_string", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithFields(tc.input)(&p)
			if p.fields == nil {
				t.Fatal("fields is nil")
			}
			if *p.fields != tc.want {
				t.Errorf("got %q, want %q", *p.fields, tc.want)
			}
		})
	}
}

// ---- WithAdditionalTypes ----

func TestWithAdditionalTypes(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"single", []string{"episode"}, []string{"episode"}},
		{"multiple", []string{"track", "episode"}, []string{"track", "episode"}},
		{"empty_slice", []string{}, []string{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithAdditionalTypes(tc.input...)(&p)
			if len(p.additionalTypes) != len(tc.want) {
				t.Fatalf("got len %d, want %d", len(p.additionalTypes), len(tc.want))
			}
			for i, v := range tc.want {
				if p.additionalTypes[i] != v {
					t.Errorf("index %d: got %q, want %q", i, p.additionalTypes[i], v)
				}
			}
		})
	}
}

// ---- WithURIs ----

func TestWithURIs(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"single", []string{"spotify:track:abc"}, []string{"spotify:track:abc"}},
		{"multiple", []string{"spotify:track:abc", "spotify:album:xyz"}, []string{"spotify:track:abc", "spotify:album:xyz"}},
		{"empty_slice", []string{}, []string{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithURIs(tc.input...)(&p)
			if len(p.uris) != len(tc.want) {
				t.Fatalf("got len %d, want %d", len(p.uris), len(tc.want))
			}
			for i, v := range tc.want {
				if p.uris[i] != v {
					t.Errorf("index %d: got %q, want %q", i, p.uris[i], v)
				}
			}
		})
	}
}

// ---- WithPosition ----

func TestWithPosition(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 5, 5},
		{"large", 1000, 1000},
		{"negative", -1, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithPosition(tc.input)(&p)
			if p.position == nil {
				t.Fatal("position is nil")
			}
			if *p.position != tc.want {
				t.Errorf("got %d, want %d", *p.position, tc.want)
			}
		})
	}
}

// ---- WithDeviceID ----

func TestWithDeviceID(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"normal_id", "abc123xyz", "abc123xyz"},
		{"empty_string", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithDeviceID(tc.input)(&p)
			if p.deviceID == nil {
				t.Fatal("deviceID is nil")
			}
			if *p.deviceID != tc.want {
				t.Errorf("got %q, want %q", *p.deviceID, tc.want)
			}
		})
	}
}

func TestWithDeviceID_PointerIndependence(t *testing.T) {
	val := "device1"
	p := queryParameters{}
	WithDeviceID(val)(&p)
	val = "device2"
	if *p.deviceID != "device1" {
		t.Errorf("deviceID mutated after caller change: got %q", *p.deviceID)
	}
}

// ---- WithAfterMs ----

func TestWithAfterMs(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"unix_ms", 1609459200000, 1609459200000},
		{"zero", 0, 0},
		{"negative", -1, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithAfterMs(tc.input)(&p)
			if p.afterMs == nil {
				t.Fatal("afterMs is nil")
			}
			if *p.afterMs != tc.want {
				t.Errorf("got %d, want %d", *p.afterMs, tc.want)
			}
		})
	}
}

// ---- WithBeforeMs ----

func TestWithBeforeMs(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"unix_ms", 1609459200000, 1609459200000},
		{"zero", 0, 0},
		{"negative", -1, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := queryParameters{}
			WithBeforeMs(tc.input)(&p)
			if p.beforeMs == nil {
				t.Fatal("beforeMs is nil")
			}
			if *p.beforeMs != tc.want {
				t.Errorf("got %d, want %d", *p.beforeMs, tc.want)
			}
		})
	}
}

// ---- toQuery ----

func TestToQuery(t *testing.T) {
	cases := []struct {
		name  string
		p     queryParameters
		want  string
	}{
		{
			"empty",
			queryParameters{},
			"",
		},
		{
			"only_time_range",
			queryParameters{timeRange: ptrTimeRange(ShortTerm)},
			"time_range=short_term",
		},
		{
			"only_limit",
			queryParameters{limit: ptrInt(20)},
			"limit=20",
		},
		{
			"only_offset",
			queryParameters{offset: ptrInt(5)},
			"offset=5",
		},
		{
			"only_after",
			queryParameters{after: ptrString("abc123")},
			"after=abc123",
		},
		{
			// url.Values.Encode sorts keys alphabetically
			"all_four",
			queryParameters{
				timeRange: ptrTimeRange(LongTerm),
				limit:     ptrInt(50),
				offset:    ptrInt(0),
				after:     ptrString("xyz"),
			},
			"after=xyz&limit=50&offset=0&time_range=long_term",
		},
		{
			"time_range_and_limit",
			queryParameters{timeRange: ptrTimeRange(MediumTerm), limit: ptrInt(10)},
			"limit=10&time_range=medium_term",
		},
		{
			"limit_and_offset",
			queryParameters{limit: ptrInt(25), offset: ptrInt(50)},
			"limit=25&offset=50",
		},
		{
			"offset_zero_emitted",
			queryParameters{offset: ptrInt(0)},
			"offset=0",
		},
		{
			"limit_min",
			queryParameters{limit: ptrInt(1)},
			"limit=1",
		},
		{
			"limit_max",
			queryParameters{limit: ptrInt(50)},
			"limit=50",
		},
		{
			"after_url_special_chars_encoded",
			queryParameters{after: ptrString("hello world&foo=bar")},
			"after=hello+world%26foo%3Dbar",
		},
		{
			"time_range_empty_string",
			queryParameters{timeRange: ptrTimeRange(TimeRange(""))},
			"time_range=",
		},
		{
			"only_market",
			queryParameters{market: ptrString("US")},
			"market=US",
		},
		{
			"market_gb",
			queryParameters{market: ptrString("GB")},
			"market=GB",
		},
		{
			// url.Values.Encode sorts keys alphabetically: after, limit, market, offset, time_range
			"all_five_fields",
			queryParameters{
				timeRange: ptrTimeRange(LongTerm),
				limit:     ptrInt(50),
				offset:    ptrInt(0),
				after:     ptrString("xyz"),
				market:    ptrString("US"),
			},
			"after=xyz&limit=50&market=US&offset=0&time_range=long_term",
		},
		{
			"market_with_limit",
			queryParameters{market: ptrString("DE"), limit: ptrInt(20)},
			"limit=20&market=DE",
		},
		{
			"only_include_groups_single",
			queryParameters{includeGroups: []string{"album"}},
			"include_groups=album",
		},
		{
			"include_groups_multiple",
			queryParameters{includeGroups: []string{"album", "single", "compilation"}},
			"include_groups=album%2Csingle%2Ccompilation",
		},
		{
			"only_include_external",
			queryParameters{includeExternal: ptrString("audio")},
			"include_external=audio",
		},
		{
			"only_fields",
			queryParameters{fields: ptrString("name,id")},
			"fields=name%2Cid",
		},
		{
			"only_additional_types",
			queryParameters{additionalTypes: []string{"episode"}},
			"additional_types=episode",
		},
		{
			"additional_types_multiple",
			queryParameters{additionalTypes: []string{"track", "episode"}},
			"additional_types=track%2Cepisode",
		},
		{
			"only_uris_single",
			queryParameters{uris: []string{"spotify:track:abc"}},
			"uris=spotify%3Atrack%3Aabc",
		},
		{
			"uris_multiple",
			queryParameters{uris: []string{"spotify:track:abc", "spotify:album:xyz"}},
			"uris=spotify%3Atrack%3Aabc%2Cspotify%3Aalbum%3Axyz",
		},
		{
			"only_position",
			queryParameters{position: ptrInt(3)},
			"position=3",
		},
		{
			"position_zero",
			queryParameters{position: ptrInt(0)},
			"position=0",
		},
		{
			"only_device_id",
			queryParameters{deviceID: ptrString("dev123")},
			"device_id=dev123",
		},
		{
			"only_after_ms",
			queryParameters{afterMs: ptrInt(1609459200000)},
			"after=1609459200000",
		},
		{
			"only_before_ms",
			queryParameters{beforeMs: ptrInt(1609459200000)},
			"before=1609459200000",
		},
		{
			"after_ms_and_limit",
			queryParameters{afterMs: ptrInt(1000), limit: ptrInt(10)},
			"after=1000&limit=10",
		},
		{
			"device_id_and_limit",
			queryParameters{deviceID: ptrString("mydev"), limit: ptrInt(5)},
			"device_id=mydev&limit=5",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.p.toQuery()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// ---- applyQueryParameters ----

func TestApplyQueryParameters(t *testing.T) {
	t.Run("no_options", func(t *testing.T) {
		p := applyQueryParameters()
		if p.timeRange != nil || p.limit != nil || p.offset != nil || p.after != nil || p.market != nil {
			t.Error("expected all fields to be nil")
		}
	})

	t.Run("single_time_range", func(t *testing.T) {
		p := applyQueryParameters(WithTimeRange(ShortTerm))
		if p.timeRange == nil || *p.timeRange != ShortTerm {
			t.Errorf("unexpected timeRange: %v", p.timeRange)
		}
		if p.limit != nil || p.offset != nil || p.after != nil {
			t.Error("expected other fields to be nil")
		}
	})

	t.Run("single_limit", func(t *testing.T) {
		p := applyQueryParameters(WithLimit(30))
		if p.limit == nil || *p.limit != 30 {
			t.Errorf("unexpected limit: %v", p.limit)
		}
	})

	t.Run("single_offset", func(t *testing.T) {
		p := applyQueryParameters(WithOffset(5))
		if p.offset == nil || *p.offset != 5 {
			t.Errorf("unexpected offset: %v", p.offset)
		}
	})

	t.Run("single_after", func(t *testing.T) {
		p := applyQueryParameters(WithAfter("tok123"))
		if p.after == nil || *p.after != "tok123" {
			t.Errorf("unexpected after: %v", p.after)
		}
	})

	t.Run("all_four_options", func(t *testing.T) {
		p := applyQueryParameters(WithTimeRange(LongTerm), WithLimit(50), WithOffset(10), WithAfter("xyz"))
		if *p.timeRange != LongTerm || *p.limit != 50 || *p.offset != 10 || *p.after != "xyz" {
			t.Errorf("unexpected values: %v %v %v %v", *p.timeRange, *p.limit, *p.offset, *p.after)
		}
	})

	t.Run("last_write_wins_time_range", func(t *testing.T) {
		p := applyQueryParameters(WithTimeRange(ShortTerm), WithTimeRange(LongTerm))
		if *p.timeRange != LongTerm {
			t.Errorf("got %q, want long_term", *p.timeRange)
		}
	})

	t.Run("last_write_wins_limit", func(t *testing.T) {
		p := applyQueryParameters(WithLimit(10), WithLimit(40))
		if *p.limit != 40 {
			t.Errorf("got %d, want 40", *p.limit)
		}
	})

	t.Run("last_write_wins_offset", func(t *testing.T) {
		p := applyQueryParameters(WithOffset(100), WithOffset(5))
		if *p.offset != 5 {
			t.Errorf("got %d, want 5", *p.offset)
		}
	})

	t.Run("last_write_wins_after", func(t *testing.T) {
		p := applyQueryParameters(WithAfter("first"), WithAfter("second"))
		if *p.after != "second" {
			t.Errorf("got %q, want second", *p.after)
		}
	})

	t.Run("clamping_applied_per_call_limit", func(t *testing.T) {
		// second WithLimit(0) clamps to 1, overwriting first call's clamped 50
		p := applyQueryParameters(WithLimit(200), WithLimit(0))
		if *p.limit != 1 {
			t.Errorf("got %d, want 1", *p.limit)
		}
	})

	t.Run("interleaved_options_last_limit_wins", func(t *testing.T) {
		p := applyQueryParameters(WithLimit(5), WithTimeRange(MediumTerm), WithOffset(20), WithLimit(15))
		if *p.limit != 15 || *p.timeRange != MediumTerm || *p.offset != 20 || p.after != nil {
			t.Errorf("unexpected values: limit=%v timeRange=%v offset=%v after=%v", *p.limit, *p.timeRange, *p.offset, p.after)
		}
	})

	t.Run("boundary_clamping_combined", func(t *testing.T) {
		p := applyQueryParameters(WithLimit(-1), WithOffset(-1))
		if *p.limit != 1 || *p.offset != 0 {
			t.Errorf("got limit=%d offset=%d, want limit=1 offset=0", *p.limit, *p.offset)
		}
	})

	t.Run("single_market", func(t *testing.T) {
		p := applyQueryParameters(WithMarket("US"))
		if p.market == nil || *p.market != "US" {
			t.Errorf("unexpected market: %v", p.market)
		}
		if p.timeRange != nil || p.limit != nil || p.offset != nil || p.after != nil {
			t.Error("expected other fields to be nil")
		}
	})

	t.Run("last_write_wins_market", func(t *testing.T) {
		p := applyQueryParameters(WithMarket("US"), WithMarket("GB"))
		if *p.market != "GB" {
			t.Errorf("got %q, want GB", *p.market)
		}
	})

	t.Run("single_include_groups", func(t *testing.T) {
		p := applyQueryParameters(WithIncludeGroups("album", "single"))
		if len(p.includeGroups) != 2 || p.includeGroups[0] != "album" || p.includeGroups[1] != "single" {
			t.Errorf("unexpected includeGroups: %v", p.includeGroups)
		}
	})

	t.Run("single_include_external", func(t *testing.T) {
		p := applyQueryParameters(WithIncludeExternal("audio"))
		if p.includeExternal == nil || *p.includeExternal != "audio" {
			t.Errorf("unexpected includeExternal: %v", p.includeExternal)
		}
	})

	t.Run("single_fields", func(t *testing.T) {
		p := applyQueryParameters(WithFields("name,id"))
		if p.fields == nil || *p.fields != "name,id" {
			t.Errorf("unexpected fields: %v", p.fields)
		}
	})

	t.Run("single_additional_types", func(t *testing.T) {
		p := applyQueryParameters(WithAdditionalTypes("track", "episode"))
		if len(p.additionalTypes) != 2 || p.additionalTypes[0] != "track" || p.additionalTypes[1] != "episode" {
			t.Errorf("unexpected additionalTypes: %v", p.additionalTypes)
		}
	})

	t.Run("single_uris", func(t *testing.T) {
		p := applyQueryParameters(WithURIs("spotify:track:abc", "spotify:album:xyz"))
		if len(p.uris) != 2 || p.uris[0] != "spotify:track:abc" || p.uris[1] != "spotify:album:xyz" {
			t.Errorf("unexpected uris: %v", p.uris)
		}
	})

	t.Run("single_position", func(t *testing.T) {
		p := applyQueryParameters(WithPosition(3))
		if p.position == nil || *p.position != 3 {
			t.Errorf("unexpected position: %v", p.position)
		}
	})

	t.Run("single_device_id", func(t *testing.T) {
		p := applyQueryParameters(WithDeviceID("dev1"))
		if p.deviceID == nil || *p.deviceID != "dev1" {
			t.Errorf("unexpected deviceID: %v", p.deviceID)
		}
	})

	t.Run("single_after_ms", func(t *testing.T) {
		p := applyQueryParameters(WithAfterMs(1609459200000))
		if p.afterMs == nil || *p.afterMs != 1609459200000 {
			t.Errorf("unexpected afterMs: %v", p.afterMs)
		}
	})

	t.Run("single_before_ms", func(t *testing.T) {
		p := applyQueryParameters(WithBeforeMs(1609459200000))
		if p.beforeMs == nil || *p.beforeMs != 1609459200000 {
			t.Errorf("unexpected beforeMs: %v", p.beforeMs)
		}
	})

	t.Run("last_write_wins_device_id", func(t *testing.T) {
		p := applyQueryParameters(WithDeviceID("dev1"), WithDeviceID("dev2"))
		if *p.deviceID != "dev2" {
			t.Errorf("got %q, want dev2", *p.deviceID)
		}
	})

	t.Run("last_write_wins_fields", func(t *testing.T) {
		p := applyQueryParameters(WithFields("name"), WithFields("id,uri"))
		if *p.fields != "id,uri" {
			t.Errorf("got %q, want id,uri", *p.fields)
		}
	})

	t.Run("last_write_wins_position", func(t *testing.T) {
		p := applyQueryParameters(WithPosition(1), WithPosition(5))
		if *p.position != 5 {
			t.Errorf("got %d, want 5", *p.position)
		}
	})
}
