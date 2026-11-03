package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMakeTreeAndListNames(t *testing.T) {
	base := t.TempDir()

	dirs := []string{"alpha", filepath.Join("alpha", "beta"), "gamma"}
	if err := makeTree(base, dirs...); err != nil {
		t.Fatalf("makeTree(%q, %q) returned error: %v", base, dirs, err)
	}

	for _, dir := range dirs {
		info, err := os.Stat(filepath.Join(base, dir))
		if err != nil {
			t.Errorf("stat %q: %v", filepath.Join(base, dir), err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("%q is not a directory", filepath.Join(base, dir))
		}
	}

	if err := os.WriteFile(filepath.Join(base, "zeta.txt"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: writing zeta.txt: %v", err)
	}

	got, err := listNames(base)
	if err != nil {
		t.Fatalf("listNames(%q) returned error: %v", base, err)
	}
	want := []string{"alpha", "gamma", "zeta.txt"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("listNames(%q) = %v, want %v (sorted)", base, got, want)
	}

	nested, err := listNames(filepath.Join(base, "alpha"))
	if err != nil {
		t.Fatalf("listNames(nested) returned error: %v", err)
	}
	if wantNested := []string{"beta"}; !reflect.DeepEqual(nested, wantNested) {
		t.Errorf("listNames(nested) = %v, want %v", nested, wantNested)
	}
}

func TestMakeTreeIdempotent(t *testing.T) {
	base := t.TempDir()
	dirs := []string{"logs", filepath.Join("logs", "2026")}

	if err := makeTree(base, dirs...); err != nil {
		t.Fatalf("first makeTree failed: %v", err)
	}
	if err := makeTree(base, dirs...); err != nil {
		t.Errorf("second makeTree failed: %v", err)
	}
}

func TestListNamesMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "nope")

	got, err := listNames(missing)
	if err == nil {
		t.Errorf("listNames(%q) error = nil, want an error", missing)
	}
	if len(got) != 0 {
		t.Errorf("listNames(%q) = %v, want no names on error", missing, got)
	}
}
