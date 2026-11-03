// 075: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func dispatch(args []string, out io.Writer) error {
	if len(args) == 0 {
		return fmt.Errorf("missing subcommand")
	}

	switch args[0] {
	case "greet":
		if len(args) != 2 {
			return fmt.Errorf("greet wants exactly one name")
		}
		fmt.Fprintf(out, "Hello, %s!\n", args[1])
		return nil
	case "add":
		if len(args) < 2 {
			return fmt.Errorf("add wants at least one number")
		}
		total := 0
		for _, raw := range args[1:] {
			n, err := strconv.Atoi(raw)
			if err != nil {
				return fmt.Errorf("invalid number %q: %w", raw, err)
			}
			total += n
		}
		fmt.Fprintf(out, "sum: %d\n", total)
		return nil
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func main() {
	if err := dispatch(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
