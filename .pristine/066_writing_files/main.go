// 066: Writing Files. os.WriteFile creates or truncates a file and writes the
// whole byte slice; an existing file is replaced, not appended to.
//
// Scenario: a nightly job writes a tiny status report. A short report must not
// leave bytes from yesterday's longer report behind.
//
// Rules:
//  1. writeText saves content at path with os.WriteFile and mode 0o644.
//  2. Writing again replaces the file completely, so old bytes disappear.
//  3. Writing "" produces an existing, empty file (not an error).
//  4. Missing parent directories and other failures return a non-nil error.
//
// TODO: implement writeText below. Keep main() as the demo harness.
package main

import (
	"fmt"
)

// writeText writes content to path, replacing any existing file.
func writeText(path, content string) error {
	// TODO
	return nil
}

func main() {
	err := writeText("status.txt", "all systems go\n")
	fmt.Println("write error:", err)
}
