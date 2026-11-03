package main

import "testing"

func TestWeekdayString(t *testing.T) {
	tests := []struct {
		name string
		d    Weekday
		want string
	}{
		{"sunday", Sunday, "Sunday"},
		{"monday", Monday, "Monday"},
		{"tuesday", Tuesday, "Tuesday"},
		{"wednesday", Wednesday, "Wednesday"},
		{"thursday", Thursday, "Thursday"},
		{"friday", Friday, "Friday"},
		{"saturday", Saturday, "Saturday"},
		{"unknown positive", Weekday(99), "Weekday(99)"},
		{"unknown negative", Weekday(-1), "Weekday(-1)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Weekday(%d).String() = %q, want %q", int(tt.d), got, tt.want)
			}
		})
	}
}

func TestParseWeekday(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		want   Weekday
		wantOK bool
	}{
		{"sunday", "Sunday", Sunday, true},
		{"saturday", "Saturday", Saturday, true},
		{"wednesday", "Wednesday", Wednesday, true},
		{"lowercase rejected", "sunday", Sunday, false},
		{"unknown name", "Funday", Sunday, false},
		{"empty", "", Sunday, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := parseWeekday(tt.in)
			if ok != tt.wantOK {
				t.Fatalf("parseWeekday(%q) ok = %v, want %v", tt.in, ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("parseWeekday(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name string
		d    Weekday
		want bool
	}{
		{"sunday", Sunday, true},
		{"saturday", Saturday, true},
		{"monday", Monday, false},
		{"friday", Friday, false},
		{"unknown", Weekday(99), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWeekend(tt.d); got != tt.want {
				t.Errorf("isWeekend(%v) = %v, want %v", tt.d, got, tt.want)
			}
		})
	}
}

func TestNextDay(t *testing.T) {
	tests := []struct {
		name string
		d    Weekday
		want Weekday
	}{
		{"sunday to monday", Sunday, Monday},
		{"friday to saturday", Friday, Saturday},
		{"saturday wraps", Saturday, Sunday},
		{"unknown unchanged", Weekday(99), Weekday(99)},
		{"negative unchanged", Weekday(-1), Weekday(-1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextDay(tt.d); got != tt.want {
				t.Errorf("nextDay(%v) = %v, want %v", tt.d, got, tt.want)
			}
		})
	}
}
