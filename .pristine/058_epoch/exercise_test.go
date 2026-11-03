package main

import (
	"testing"
	"time"
)

func epochUTC(y int, m time.Month, d, hh, mm, ss, ns int) time.Time {
	return time.Date(y, m, d, hh, mm, ss, ns, time.UTC)
}

var epochBase = epochUTC(2024, time.February, 29, 12, 34, 56, 789000000)

func TestUnixSeconds(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want int64
	}{
		{"epoch", epochUTC(1970, time.January, 1, 0, 0, 0, 0), 0},
		{"classic", epochUTC(2009, time.November, 10, 23, 0, 0, 0), 1257894000},
		{"leap day", epochBase, 1709210096},
		{"before epoch", epochUTC(1969, time.December, 31, 23, 59, 59, 0), -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unixSeconds(tt.in); got != tt.want {
				t.Errorf("unixSeconds(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestUnixMillis(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want int64
	}{
		{"epoch", epochUTC(1970, time.January, 1, 0, 0, 0, 0), 0},
		{"leap day", epochBase, 1709210096789},
		{"before epoch", epochUTC(1969, time.December, 31, 23, 59, 59, 0), -1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unixMillis(tt.in); got != tt.want {
				t.Errorf("unixMillis(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestUnixNanos(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want int64
	}{
		{"epoch", epochUTC(1970, time.January, 1, 0, 0, 0, 0), 0},
		{"leap day", epochBase, 1709210096789000000},
		{"one nanosecond", epochUTC(1970, time.January, 1, 0, 0, 0, 1), 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unixNanos(tt.in); got != tt.want {
				t.Errorf("unixNanos(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestFromUnix(t *testing.T) {
	tests := []struct {
		name string
		got  time.Time
		want time.Time
	}{
		{"epoch", fromUnixSeconds(0), epochUTC(1970, time.January, 1, 0, 0, 0, 0)},
		{"classic", fromUnixSeconds(1257894000), epochUTC(2009, time.November, 10, 23, 0, 0, 0)},
		{"before epoch", fromUnixSeconds(-1), epochUTC(1969, time.December, 31, 23, 59, 59, 0)},
		{"millis", fromUnixMillis(1709210096789), epochBase},
		{"millis truncate micros", fromUnixMillis(1709210096789), epochBase},
		{"nanos", fromUnixNanos(1709210096789000000), epochBase},
		{"nanos one", fromUnixNanos(1), epochUTC(1970, time.January, 1, 0, 0, 0, 1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.got.Equal(tt.want) {
				t.Errorf("conversion = %v, want %v", tt.got, tt.want)
			}
			if tt.got.Location() != time.UTC {
				t.Errorf("conversion location = %v, want UTC", tt.got.Location())
			}
		})
	}
}

func TestEpochRoundTrips(t *testing.T) {
	times := []time.Time{
		epochUTC(1970, time.January, 1, 0, 0, 0, 0),
		epochUTC(2009, time.November, 10, 23, 0, 0, 0),
		epochBase,
	}

	for _, want := range times {
		t.Run(want.Format(time.RFC3339Nano), func(t *testing.T) {
			if got := fromUnixSeconds(unixSeconds(want)); !got.Equal(want.Truncate(time.Second)) {
				t.Errorf("second roundtrip = %v, want %v", got, want.Truncate(time.Second))
			}
			if got := fromUnixMillis(unixMillis(want)); !got.Equal(want.Truncate(time.Millisecond)) {
				t.Errorf("milli roundtrip = %v, want %v", got, want.Truncate(time.Millisecond))
			}
			if got := fromUnixNanos(unixNanos(want)); !got.Equal(want) {
				t.Errorf("nano roundtrip = %v, want %v", got, want)
			}
		})
	}
}
