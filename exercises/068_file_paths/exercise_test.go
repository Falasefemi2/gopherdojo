package main

import (
	"path/filepath"
	"testing"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{"three parts", []string{"data", "2026", "report.txt"}, filepath.Join("data", "2026", "report.txt")},
		{"one part", []string{"report.txt"}, "report.txt"},
		{"no parts", nil, ""},
		{"empty part skipped", []string{"data", "", "report.txt"}, filepath.Join("data", "report.txt")},
		{"dot segments cleaned", []string{"data", "..", "logs", "app.log"}, filepath.Join("logs", "app.log")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinPath(tt.parts...); got != tt.want {
				t.Errorf("joinPath(%q) = %q, want %q", tt.parts, got, tt.want)
			}
		})
	}
}

func TestSplitPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantDir  string
		wantBase string
		wantExt  string
	}{
		{
			name:     "nested file",
			path:     filepath.Join("data", "2026", "report.txt"),
			wantDir:  filepath.Join("data", "2026"),
			wantBase: "report.txt",
			wantExt:  ".txt",
		},
		{
			name:     "no extension",
			path:     filepath.Join("archive", "README"),
			wantDir:  "archive",
			wantBase: "README",
			wantExt:  "",
		},
		{
			name:     "dotfile has a dot extension",
			path:     filepath.Join("repo", ".gitignore"),
			wantDir:  "repo",
			wantBase: ".gitignore",
			wantExt:  ".gitignore",
		},
		{
			name:     "double extension uses the last",
			path:     "report.tar.gz",
			wantDir:  ".",
			wantBase: "report.tar.gz",
			wantExt:  ".gz",
		},
		{
			name:     "bare name",
			path:     "Makefile",
			wantDir:  ".",
			wantBase: "Makefile",
			wantExt:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, base, ext := splitPath(tt.path)
			if dir != tt.wantDir {
				t.Errorf("splitPath(%q) dir = %q, want %q", tt.path, dir, tt.wantDir)
			}
			if base != tt.wantBase {
				t.Errorf("splitPath(%q) base = %q, want %q", tt.path, base, tt.wantBase)
			}
			if ext != tt.wantExt {
				t.Errorf("splitPath(%q) ext = %q, want %q", tt.path, ext, tt.wantExt)
			}
		})
	}
}

func TestCleanPath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"parent segment", "a/../b/./c", filepath.Join("b", "c")},
		{"duplicate separators", "a//b/", filepath.Join("a", "b")},
		{"empty becomes dot", "", "."},
		{"dot stays dot", ".", "."},
		{"already clean", filepath.Join("data", "report.txt"), filepath.Join("data", "report.txt")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanPath(tt.in); got != tt.want {
				t.Errorf("cleanPath(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
