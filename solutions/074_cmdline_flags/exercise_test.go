package main

import "testing"

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Config
	}{
		{"defaults", nil, Config{Name: "world", Count: 1}},
		{"name only", []string{"-name", "dojo"}, Config{Name: "dojo", Count: 1}},
		{
			"all flags",
			[]string{"-name=api", "-count=3", "-verbose"},
			Config{Name: "api", Count: 3, Verbose: true},
		},
		{"explicit false", []string{"-verbose=false"}, Config{Name: "world", Count: 1}},
		{"count zero", []string{"-count", "0"}, Config{Name: "world", Count: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFlags(tt.args)
			if err != nil {
				t.Fatalf("parseFlags(%q) unexpected error: %v", tt.args, err)
			}
			if got != tt.want {
				t.Errorf("parseFlags(%q) = %+v, want %+v", tt.args, got, tt.want)
			}
		})
	}
}

func TestParseFlagsErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"unknown flag", []string{"-nope"}},
		{"missing value", []string{"-name"}},
		{"bad count", []string{"-count", "many"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFlags(tt.args)
			if err == nil {
				t.Fatalf("parseFlags(%q) error = nil, want an error", tt.args)
			}
			if got != (Config{}) {
				t.Errorf("parseFlags(%q) = %+v on error, want the zero Config", tt.args, got)
			}
		})
	}
}
