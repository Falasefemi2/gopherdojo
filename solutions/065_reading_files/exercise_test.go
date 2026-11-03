package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("setup: writing %s: %v", path, err)
	}
	return path
}

func TestReadText(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{"simple line", "alpha\nbeta\n"},
		{"no trailing newline", "single line"},
		{"empty file", ""},
		{"unicode", "café ☕\n日本語\n"},
		{"blank lines", "a\n\nb\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeTempFile(t, "input.txt", tt.content)
			got, err := readText(path)
			if err != nil {
				t.Fatalf("readText(%q) returned unexpected error: %v", path, err)
			}
			if got != tt.content {
				t.Errorf("readText(%q) = %q, want %q", path, got, tt.content)
			}
		})
	}
}

func TestReadTextMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does-not-exist.txt")

	got, err := readText(path)
	if err == nil {
		t.Errorf("readText(%q) error = nil, want an error", path)
	}
	if got != "" {
		t.Errorf("readText(%q) = %q, want the empty string on error", path, got)
	}
}

func TestReadTextEmptyPath(t *testing.T) {
	got, err := readText("")
	if err == nil {
		t.Errorf("readText(\"\") error = nil, want an error")
	}
	if got != "" {
		t.Errorf("readText(\"\") = %q, want the empty string on error", got)
	}
}
