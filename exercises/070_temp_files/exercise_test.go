package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMakeTempFile(t *testing.T) {
	dir := t.TempDir()

	name, err := makeTempFile(dir, "job-*.log", "payload\n")
	if err != nil {
		t.Fatalf("makeTempFile returned error: %v", err)
	}
	if name == "" {
		t.Fatal("makeTempFile returned an empty name")
	}

	if got := filepath.Dir(name); got != dir {
		t.Errorf("file dir = %q, want %q", got, dir)
	}
	base := filepath.Base(name)
	if !strings.HasPrefix(base, "job-") {
		t.Errorf("file name %q does not start with pattern prefix %q", base, "job-")
	}
	if !strings.HasSuffix(base, ".log") {
		t.Errorf("file name %q does not end with pattern suffix %q", base, ".log")
	}
	if base == "job-.log" {
		t.Errorf("file name %q has no unique part", base)
	}

	got, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("reading temp file %q: %v", name, err)
	}
	if string(got) != "payload\n" {
		t.Errorf("temp file contents = %q, want %q", string(got), "payload\n")
	}

	other, err := makeTempFile(dir, "job-*.log", "")
	if err != nil {
		t.Fatalf("second makeTempFile returned error: %v", err)
	}
	if other == name {
		t.Errorf("two calls returned the same name %q", name)
	}

	if err := os.Remove(name); err != nil {
		t.Errorf("removing temp file %q: %v", name, err)
	}
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		t.Errorf("temp file %q still exists after removal, stat err = %v", name, err)
	}
}

func TestMakeTempDir(t *testing.T) {
	dir := t.TempDir()

	path, err := makeTempDir(dir, "build-*")
	if err != nil {
		t.Fatalf("makeTempDir returned error: %v", err)
	}
	if path == "" {
		t.Fatal("makeTempDir returned an empty path")
	}

	if got := filepath.Dir(path); got != dir {
		t.Errorf("temp dir parent = %q, want %q", got, dir)
	}
	base := filepath.Base(path)
	if !strings.HasPrefix(base, "build-") {
		t.Errorf("temp dir name %q does not start with %q", base, "build-")
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat temp dir %q: %v", path, err)
	}
	if !info.IsDir() {
		t.Errorf("%q is not a directory", path)
	}

	inner := filepath.Join(path, "notes.txt")
	if err := os.WriteFile(inner, []byte("scratch"), 0o644); err != nil {
		t.Errorf("writing inside temp dir: %v", err)
	}

	other, err := makeTempDir(dir, "build-*")
	if err != nil {
		t.Fatalf("second makeTempDir returned error: %v", err)
	}
	if other == path {
		t.Errorf("two calls returned the same path %q", path)
	}

	if err := os.RemoveAll(path); err != nil {
		t.Errorf("removing temp dir %q: %v", path, err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("temp dir %q still exists after removal, stat err = %v", path, err)
	}
}

func TestTempFunctionsBadDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")

	name, err := makeTempFile(missing, "job-*.log", "payload")
	if err == nil {
		t.Errorf("makeTempFile(%q) error = nil, want an error", missing)
	}
	if name != "" {
		t.Errorf("makeTempFile(%q) = %q, want an empty name on error", missing, name)
	}

	path, err := makeTempDir(missing, "build-*")
	if err == nil {
		t.Errorf("makeTempDir(%q) error = nil, want an error", missing)
	}
	if path != "" {
		t.Errorf("makeTempDir(%q) = %q, want an empty path on error", missing, path)
	}
}
