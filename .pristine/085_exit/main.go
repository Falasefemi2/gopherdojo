// 085: Exit. os.Exit ends the process immediately with a status code; mapping
// errors to codes belongs in a pure function tests can call.
//
// Scenario: a lint tool exits 0 on success, 2 on bad usage and 1 on any other
// failure.
//
// Rules:
//  1. exitCode(nil) is 0.
//  2. exitCode returns 2 for errUsage, including wrapped values (errors.Is must
//     find it).
//  3. Every other non-nil error returns 1.
//  4. exitCode stays pure: no printing, no os.Exit, no panics.
//
// TODO: implement exitCode below. main() is the only place that calls os.Exit.
package main

import (
	"errors"
	"fmt"
	"os"
)

// errUsage marks a bad invocation rather than an operational failure.
var errUsage = errors.New("usage")

// exitCode maps an error to a process exit status.
func exitCode(err error) int {
	// TODO
	return 0
}

// run performs the demo work and reports usage problems with errUsage.
func run() error {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		return fmt.Errorf("no help here: %w", errUsage)
	}
	return nil
}

func main() {
	os.Exit(exitCode(run()))
}
