// 068: File Paths. path/filepath joins and splits paths with the separator of
// the current OS, so code stays correct on Windows and Unix alike.
//
// Scenario: a backup tool builds a report path from folders, then splits it
// again to show the folder, file name and extension in its log.
//
// Rules:
//  1. joinPath joins every part with filepath.Join, which also cleans the result.
//  2. splitPath returns dir, base and ext using filepath.Dir/Base/Ext.
//  3. cleanPath normalises separators and "." / ".." with filepath.Clean.
//  4. Never hard-code "/" or "\": the same code must run on any OS.
package main

import (
	"fmt"
	"path/filepath"
)

// joinPath joins path parts for the current OS.
func joinPath(parts ...string) string {
	return filepath.Join(parts...)
}

// splitPath returns the directory, base name and extension of p.
func splitPath(p string) (dir, base, ext string) {
	return filepath.Dir(p), filepath.Base(p), filepath.Ext(p)
}

// cleanPath returns the shortest cleaned form of p.
func cleanPath(p string) string {
	return filepath.Clean(p)
}

func main() {
	p := joinPath("data", "2026", "report.txt")
	dir, base, ext := splitPath(p)
	fmt.Println("path:", p)
	fmt.Println("dir:", dir, "base:", base, "ext:", ext)
	fmt.Println("cleaned:", cleanPath("a/../b/./c"))
}
