// 085: Reference solution. Read this only after your own attempt passes.
package main

import (
	"errors"
	"fmt"
	"os"
)

// errUsage marks a bad invocation rather than an operational failure.
var errUsage = errors.New("usage")

func exitCode(err error) int {
	switch {
	case err == nil:
		return 0
	case errors.Is(err, errUsage):
		return 2
	default:
		return 1
	}
}

func run() error {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		return fmt.Errorf("no help here: %w", errUsage)
	}
	return nil
}

func main() {
	os.Exit(exitCode(run()))
}
