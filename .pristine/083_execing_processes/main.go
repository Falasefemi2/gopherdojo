// 083: Exec'ing Processes. CombinedOutput runs a command to completion and
// captures stdout and stderr in one buffer.
//
// Scenario: an install script branches on the target OS reported by the
// toolchain.
//
// Rules:
//  1. goos runs "go env GOOS" and returns its trimmed output.
//  2. stderr is captured too (CombinedOutput), not discarded.
//  3. A launch failure or a non-zero exit returns an error and an empty string.
//  4. The result contains no surrounding whitespace.
//
// TODO: implement goos below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"runtime"
)

// goos returns the GOOS reported by the Go toolchain.
func goos() (string, error) {
	// TODO
	return "", nil
}

func main() {
	v, err := goos()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("GOOS:", v, "runtime:", runtime.GOOS)
}
