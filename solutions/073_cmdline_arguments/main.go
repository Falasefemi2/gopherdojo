// 073: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumArgs(args []string) (int, error) {
	total := 0
	for _, raw := range args {
		n, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			return 0, fmt.Errorf("invalid number %q: %w", raw, err)
		}
		total += n
	}
	return total, nil
}

func main() {
	total, err := sumArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	fmt.Println("total:", total)
}
