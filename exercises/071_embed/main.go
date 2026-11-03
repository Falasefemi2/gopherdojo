// 071: Embed Directive. //go:embed pulls a file from the package directory
// into a string or []byte variable at build time.
//
// Scenario: a CLI ships a fixed greeting in data/greeting.txt so the copy can
// be edited without touching Go code, yet the binary still works on its own.
//
// Rules:
//  1. Embed the asset with //go:embed data/greeting.txt on a string variable.
//  2. The directive sits directly above the variable; the path is relative to
//     this package directory.
//  3. greeting returns the embedded contents unchanged, including the
//     trailing newline.
//  4. Never read the file from disk at run time.
//
// TODO: implement greeting below. Keep main() as the demo harness.
package main

import (
	_ "embed"
	"fmt"
)

//go:embed data/greeting.txt
var greetingFile string

// greeting returns the embedded contents of data/greeting.txt.
func greeting() string {
	// TODO
	return ""
}

func main() {
	fmt.Print(greeting())
}
