// 074: Command-Line Flags. The flag package turns -name value pairs into typed
// variables without slicing os.Args by hand.
//
// Scenario: a deploy tool accepts a service name, a repeat count and a verbose
// switch; tests must run without touching the global flag set.
//
// Rules:
//  1. parseFlags parses args (already without the program name) with its own
//     flag.NewFlagSet("app", flag.ContinueOnError).
//  2. Defaults are Name "world", Count 1 and Verbose false.
//  3. -name, -count and -verbose override those defaults.
//  4. A parse error returns an empty Config and the error, and the FlagSet must
//     not write usage text to the test log.
//
// TODO: implement parseFlags below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"os"
)

// Config holds the parsed command-line options.
type Config struct {
	Name    string
	Count   int
	Verbose bool
}

// parseFlags parses args into a Config using a private FlagSet.
func parseFlags(args []string) (Config, error) {
	// TODO
	return Config{}, nil
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	fmt.Printf("name=%s count=%d verbose=%t\n", cfg.Name, cfg.Count, cfg.Verbose)
}
