package main

import (
	"testing"
	"time"
)

func utcDate(y int, m time.Month, d, hh, mm int) time.Time {
	return time.Date(y, m, d, hh, mm, 0, 0, time.UTC)
}

func TestPlus(t *testing.T) {
	base := utcDate(2026, time.March, 2, 9, 0)
	tests := []struct {
		name string
		d    time.Duration
		want time.Time
	}{
		{"add minutes", 90 * time.Minute, utcDate(2026, time.March, 2, 10, 30)},
		{"subtract hours", -3 * time.Hour, utcDate(2026, time.March, 2, 6, 0)},
		{"cross midnight", 16 * time.Hour, utcDate(2026, time.March, 3, 1, 0)},
		{"zero shift", 0, base},
		{"cross month", 30 * 24 * time.Hour, utcDate(2026, time.April, 1, 9, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plus(base, tt.d)
			if !got.Equal(tt.want) {
				t.Errorf("plus(2026-03-02 09:00 UTC, %v) = %v, want %v", tt.d, got, tt.want)
			}
			if got.Location() != time.UTC {
				t.Errorf("plus(...) location = %v, want UTC", got.Location())
			}
		})
	}
}

func TestPlusLeavesInputUnchanged(t *testing.T) {
	base := utcDate(2026, time.March, 2, 9, 0)
	want := base
	_ = plus(base, 48*time.Hour)
	if !base.Equal(want) {
		t.Errorf("plus mutated its input: got %v, want %v", base, want)
	}
}

func TestGap(t *testing.T) {
	start := utcDate(2026, time.March, 2, 9, 0)
	tests := []struct {
		name string
		from time.Time
		to   time.Time
		want time.Duration
	}{
		{"two hours", start, start.Add(2 * time.Hour), 2 * time.Hour},
		{"across days", utcDate(2026, time.March, 1, 23, 0), utcDate(2026, time.March, 3, 1, 0), 26 * time.Hour},
		{"negative", start, start.Add(-30 * time.Minute), -30 * time.Minute},
		{"zero", start, start, 0},
		{"sub-minute", start, start.Add(90 * time.Second), 90 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gap(tt.from, tt.to); got != tt.want {
				t.Errorf("gap(%v, %v) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}

func TestEarlier(t *testing.T) {
	base := utcDate(2026, time.March, 2, 9, 0)
	tests := []struct {
		name string
		a    time.Time
		b    time.Time
		want bool
	}{
		{"earlier", base, base.Add(time.Minute), true},
		{"later", base.Add(time.Minute), base, false},
		{"equal", base, base, false},
		{"across years", utcDate(2025, time.December, 31, 23, 59), utcDate(2026, time.January, 1, 0, 0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earlier(tt.a, tt.b); got != tt.want {
				t.Errorf("earlier(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
