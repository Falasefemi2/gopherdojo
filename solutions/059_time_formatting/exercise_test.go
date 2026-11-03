package main

import (
	"testing"
	"time"
)

func TestFormatStamp(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want string
	}{
		{"typical", time.Date(2026, time.July, 4, 8, 5, 0, 0, time.UTC), "2026-07-04 08:05"},
		{"midnight", time.Date(2026, time.January, 9, 0, 0, 0, 0, time.UTC), "2026-01-09 00:00"},
		{"seconds dropped", time.Date(2026, time.December, 31, 23, 59, 42, 0, time.UTC), "2026-12-31 23:59"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStamp(tt.in); got != tt.want {
				t.Errorf("formatStamp(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseStamp(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want time.Time
	}{
		{"typical", "2026-12-31 23:59", time.Date(2026, time.December, 31, 23, 59, 0, 0, time.UTC)},
		{"leap day", "2024-02-29 00:00", time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)},
		{"epoch", "1970-01-01 00:00", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStamp(tt.in)
			if err != nil {
				t.Fatalf("parseStamp(%q) returned unexpected error: %v", tt.in, err)
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseStamp(%q) = %v, want %v", tt.in, got, tt.want)
			}
			if got.Location() != time.UTC {
				t.Errorf("parseStamp(%q) location = %v, want UTC", tt.in, got.Location())
			}
		})
	}
}

func TestParseStampErrors(t *testing.T) {
	bad := []struct {
		name string
		in   string
	}{
		{"missing time", "2026-12-31"},
		{"wrong order", "31/12/2026 23:59"},
		{"empty", ""},
		{"impossible day", "2026-02-30 10:00"},
		{"impossible month", "2026-13-01 00:00"},
		{"trailing text", "2026-12-31 23:59 UTC"},
	}

	for _, tt := range bad {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStamp(tt.in)
			if err == nil {
				t.Errorf("parseStamp(%q) error = nil, want an error", tt.in)
			}
			if !got.IsZero() {
				t.Errorf("parseStamp(%q) = %v, want the zero time", tt.in, got)
			}
		})
	}
}

func TestStampRoundTrip(t *testing.T) {
	want := time.Date(2026, time.November, 3, 6, 7, 0, 0, time.UTC)
	got, err := parseStamp(formatStamp(want))
	if err != nil {
		t.Fatalf("roundtrip returned unexpected error: %v", err)
	}
	if !got.Equal(want) {
		t.Errorf("roundtrip = %v, want %v", got, want)
	}
}
