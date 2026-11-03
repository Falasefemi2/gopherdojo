package main

import (
	"math"
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"plain", "42", 42},
		{"negative", "-7", -7},
		{"plus sign", "+13", 13},
		{"leading zeros", "007", 7},
		{"zero", "0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInt(tt.in)
			if err != nil {
				t.Fatalf("parseInt(%q) returned unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("parseInt(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseIntErrors(t *testing.T) {
	bad := []string{"", " 42", "4.2", "abc", "1_000", "999999999999999999999999"}

	for _, in := range bad {
		t.Run(in, func(t *testing.T) {
			got, err := parseInt(in)
			if err == nil {
				t.Errorf("parseInt(%q) error = nil, want an error", in)
			}
			if got != 0 {
				t.Errorf("parseInt(%q) = %d, want the zero value 0", in, got)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want float64
	}{
		{"decimal", "19.95", 19.95},
		{"negative", "-0.25", -0.25},
		{"exponent", "1e3", 1000},
		{"integer text", "8", 8},
		{"zero", "0.0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFloat(tt.in)
			if err != nil {
				t.Fatalf("parseFloat(%q) returned unexpected error: %v", tt.in, err)
			}
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("parseFloat(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseFloatErrors(t *testing.T) {
	bad := []string{"", "abc", "3.5.1", "1,5", " 2.0"}

	for _, in := range bad {
		t.Run(in, func(t *testing.T) {
			got, err := parseFloat(in)
			if err == nil {
				t.Errorf("parseFloat(%q) error = nil, want an error", in)
			}
			if got != 0 {
				t.Errorf("parseFloat(%q) = %v, want the zero value 0", in, got)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"true", "true", true},
		{"false", "false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBool(tt.in)
			if err != nil {
				t.Fatalf("parseBool(%q) returned unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("parseBool(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseBoolErrors(t *testing.T) {
	bad := []string{"", "TRUE", "False", "1", "0", "yes", " true"}

	for _, in := range bad {
		t.Run(in, func(t *testing.T) {
			got, err := parseBool(in)
			if err == nil {
				t.Errorf("parseBool(%q) error = nil, want an error", in)
			}
			if got != false {
				t.Errorf("parseBool(%q) = %v, want the zero value false", in, got)
			}
		})
	}
}
