// 075: Command-Line Subcommands. A real CLI branches on the first argument and
// hands the rest to the matching handler.
//
// Scenario: a notes tool offers `greet NAME` and `add N N ...` subcommands.
//
// Rules:
//  1. dispatch(args, out) writes results to out, never to os.Stdout.
//  2. "greet NAME" writes exactly "Hello, NAME!\n"; a missing or extra name is
//     an error and writes nothing.
//  3. "add N N ..." parses every remaining argument as an int and writes
//     exactly "sum: TOTAL\n"; invalid numbers are errors and write nothing.
//  4. An empty command list or an unknown command returns an error that names
//     the offending command when there is one.
//
// TODO: implement dispatch below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"io"
	"os"
)

// dispatch routes args to the greet or add subcommand.
func dispatch(args []string, out io.Writer) error {
	// TODO
	return nil
}

func main() {
	if err := dispatch(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
