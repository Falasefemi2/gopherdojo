// 070: Temporary Files and Directories. os.CreateTemp and os.MkdirTemp build
// unique names from a pattern; the caller owns them and should clean up.
//
// Scenario: a test runner stages each job in its own scratch file and build
// folder, then removes both when the job finishes.
//
// Rules:
//  1. makeTempFile creates a unique file in dir using pattern, writes content,
//     closes it, and returns its full path.
//  2. makeTempDir creates a unique directory in dir using pattern.
//  3. Two calls with the same pattern must produce different names.
//  4. A bad dir returns "" and a non-nil error; no partial file is left behind.
package main

import (
	"fmt"
	"os"
)

// makeTempFile creates a named temp file containing content and returns its path.
func makeTempFile(dir, pattern, content string) (string, error) {
	f, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}
	name := f.Name()

	if _, err := f.WriteString(content); err != nil {
		f.Close()
		os.Remove(name)
		return "", err
	}
	if err := f.Close(); err != nil {
		os.Remove(name)
		return "", err
	}
	return name, nil
}

// makeTempDir creates a named temp directory and returns its path.
func makeTempDir(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

func main() {
	dir, err := os.MkdirTemp("", "dojo-temp-")
	if err != nil {
		fmt.Println("temp error:", err)
		return
	}
	defer os.RemoveAll(dir)

	name, err := makeTempFile(dir, "demo-*.log", "hello\n")
	fmt.Println("file:", name, "err:", err)
}
