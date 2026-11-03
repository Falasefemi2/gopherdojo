// 074: Reference solution. Read this only after your own attempt passes.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// Config holds the parsed command-line options.
type Config struct {
	Name    string
	Count   int
	Verbose bool
}

func parseFlags(args []string) (Config, error) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	name := fs.String("name", "world", "service name")
	count := fs.Int("count", 1, "repeat count")
	verbose := fs.Bool("verbose", false, "verbose output")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}

	return Config{Name: *name, Count: *count, Verbose: *verbose}, nil
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	fmt.Printf("name=%s count=%d verbose=%t\n", cfg.Name, cfg.Count, cfg.Verbose)
}
