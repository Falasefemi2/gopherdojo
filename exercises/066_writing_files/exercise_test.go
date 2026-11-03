package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteText(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{"simple", "all systems go\n"},
		{"no trailing newline", "ready"},
		{"unicode", "café ☕\n"},
		{"empty", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "report.txt")

			if err := writeText(path, tt.content); err != nil {
				t.Fatalf("writeText(%q, %q) returned error: %v", path, tt.content, err)
			}

			got, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading back %q: %v", path, err)
			}
			if string(got) != tt.content {
				t.Errorf("file contents = %q, want %q", string(got), tt.content)
			}
		})
	}
}

func TestWriteTextOverwrites(t *testing.T) {
	path := filepath.Join(t.TempDir(), "status.txt")

	if err := writeText(path, "the first report is quite long\n"); err != nil {
		t.Fatalf("first write failed: %v", err)
	}
	if err := writeText(path, "short"); err != nil {
		t.Fatalf("second write failed: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading back %q: %v", path, err)
	}
	if string(got) != "short" {
		t.Errorf("after overwrite file = %q, want %q", string(got), "short")
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %q: %v", path, err)
	}
	if info.Size() != int64(len("short")) {
		t.Errorf("file size = %d, want %d", info.Size(), len("short"))
	}
}

func TestWriteTextEmptyTruncates(t *testing.T) {
	path := filepath.Join(t.TempDir(), "status.txt")

	if err := writeText(path, "old contents"); err != nil {
		t.Fatalf("first write failed: %v", err)
	}
	if err := writeText(path, ""); err != nil {
		t.Fatalf("empty write failed: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %q: %v", path, err)
	}
	if info.Size() != 0 {
		t.Errorf("file size = %d, want 0", info.Size())
	}
}

func TestWriteTextMissingDirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing", "report.txt")

	if err := writeText(path, "data"); err == nil {
		t.Errorf("writeText(%q) error = nil, want an error", path)
	}
}
