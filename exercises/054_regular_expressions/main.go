// 054: Regular Expressions. regexp compiles patterns once and then matches,
// extracts capture groups and rewrites text.
//
// Scenario: configuration lines are key=value pairs, and secrets must be
// masked before the lines reach a log file.
//
// Rules:
//  1. A key starts with a letter or underscore and continues with word
//     characters; a value uses word characters only.
//  2. isKeyValue is true only when the whole string is one key=value pair.
//  3. parseKeyValue returns the two capture groups and ok=false when the whole
//     string does not match; keys and values are not trimmed.
//  4. maskValues replaces the value of every pair with "***", keeping the key
//     and any surrounding text unchanged.
//
// TODO: implement isKeyValue, parseKeyValue and maskValues. Keep main() as
// the demo harness.
package main

import (
	"fmt"
	"regexp"
)

var keyValueRE = regexp.MustCompile(`^([A-Za-z_]\w*)=(\w+)$`)
var keyValueAllRE = regexp.MustCompile(`([A-Za-z_]\w*)=(\w+)`)

// isKeyValue reports whether s is exactly one key=value pair.
func isKeyValue(s string) bool {
	// TODO
	return false
}

// parseKeyValue splits a whole key=value string into its two parts.
func parseKeyValue(s string) (key, value string, ok bool) {
	// TODO
	return "", "", false
}

// maskValues replaces every value in s with ***.
func maskValues(s string) string {
	// TODO
	return ""
}

func main() {
	fmt.Println(isKeyValue("user=alice"))
	fmt.Println(parseKeyValue("user=alice"))
	fmt.Println(maskValues("user=alice pass=hunter2"))
}
