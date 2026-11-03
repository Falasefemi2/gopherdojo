// 067: Line Filters. bufio.Scanner reads an io.Reader line by line, and an
// io.Writer receives the transformed lines; neither end has to be a file.
//
// Scenario: a log pipeline keeps only lines mentioning "ERROR" and uppercases
// them for an alert channel, streaming from stdin to stdout.
//
// Rules:
//  1. Keep a line only if it contains substr; the match is case-sensitive.
//  2. Write each kept line uppercased, followed by exactly one "\n".
//  3. Preserve the input order and drop everything that does not match.
//  4. Return any scanner, read, write or flush error to the caller.
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// filterLines copies lines containing substr to w, uppercased.
func filterLines(r io.Reader, w io.Writer, substr string) error {
	scanner := bufio.NewScanner(r)
	buffered := bufio.NewWriter(w)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, substr) {
			continue
		}
		if _, err := buffered.WriteString(strings.ToUpper(line) + "\n"); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return buffered.Flush()
}

func main() {
	var out strings.Builder
	err := filterLines(strings.NewReader("alpha\nbeta\ngamma\n"), &out, "et")
	fmt.Println("err:", err)
	fmt.Print(out.String())
}
