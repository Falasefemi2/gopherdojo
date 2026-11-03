package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		app  string
		want string
	}{
		{"typical app", "ledger-api", "Hello, ledger-api!"},
		{"empty app", "", "Hello, stranger!"},
		{"single word", "payments", "Hello, payments!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(tt.app); got != tt.want {
				t.Errorf("greet(%q) = %q, want %q", tt.app, got, tt.want)
			}
		})
	}
}

func TestQuoted(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"interior quotes", `path "with" quotes`, `"path \"with\" quotes"`},
		{"plain word", "ledger", `"ledger"`},
		{"empty string", "", `""`},
		{"unicode is not escaped", "café", `"café"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quoted(tt.in); got != tt.want {
				t.Errorf("quoted(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestBanner(t *testing.T) {
	tests := []struct {
		name    string
		app     string
		version string
		want    string
	}{
		{
			name:    "ledger api",
			app:     "ledger-api",
			version: "v1.4.2",
			want:    "=== ledger-api v1.4.2 ===\nDeploying to production\nHave a nice day!",
		},
		{
			name:    "payments",
			app:     "payments",
			version: "v0.9.0",
			want:    "=== payments v0.9.0 ===\nDeploying to production\nHave a nice day!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := banner(tt.app, tt.version)
			if got != tt.want {
				t.Errorf("banner(%q, %q) = %q, want %q", tt.app, tt.version, got, tt.want)
			}
			if strings.HasSuffix(got, "\n") {
				t.Errorf("banner(%q, %q) must not end with a newline", tt.app, tt.version)
			}
			if lines := strings.Split(got, "\n"); len(lines) != 3 {
				t.Errorf("banner(%q, %q) has %d lines, want 3", tt.app, tt.version, len(lines))
			}
		})
	}
}
