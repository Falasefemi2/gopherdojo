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
package main

import (
	"fmt"
	"unicode/utf8"
)

func byteLen(s string) int {
	return len(s)
}

func runeLen(s string) int {
	return utf8.RuneCountInString(s)
}

func reverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	label := "héllo"
	fmt.Println("bytes:", byteLen(label), "runes:", runeLen(label))
	fmt.Println("reversed:", reverseRunes(label))
	fmt.Println("japanese:", reverseRunes("日本語"))
}
