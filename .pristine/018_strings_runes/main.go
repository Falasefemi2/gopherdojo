// 018: Strings and Runes. len(s) counts bytes; ranging over s or converting
// to []rune counts Unicode code points.
//
// Scenario: a terminal tool must report byte and character widths, then
// reverse non-ASCII labels without corrupting them.
//
// Rules:
//  1. byteLen(s) is len(s), the UTF-8 byte length.
//  2. runeLen(s) is the number of Unicode code points in s.
//  3. reverseRunes(s) reverses by rune, preserving every multi-byte character.
//  4. Both length helpers return 0 for the empty string, and reversing ""
//     returns "".
//
// TODO: implement the three functions below.
package main

import "fmt"

// byteLen returns the number of bytes in s.
func byteLen(s string) int {
	return 0
}

// runeLen returns the number of Unicode code points in s.
func runeLen(s string) int {
	return 0
}

// reverseRunes returns s with its runes in reverse order.
func reverseRunes(s string) string {
	return ""
}

func main() {
	label := "héllo"
	fmt.Println("bytes:", byteLen(label), "runes:", runeLen(label))
	fmt.Println("reversed:", reverseRunes(label))
	fmt.Println("japanese:", reverseRunes("日本語"))
}
