// 051: String Functions. The strings package answers questions about text and
// builds new text without hand-written loops.
//
// Scenario: a log summarizer counts markers, rebuilds a record, and normalizes
// a title for display.
//
// Rules:
//  1. matchText reports Contains, Count (non-overlapping) and the first Index,
//     or -1 when the substring is absent.
//  2. repeatJoin returns strings.Repeat(word, times), strings.Join(words, sep)
//     and the result of splitting that joined string on sep. times is >= 0.
//  3. normalizeTitle trims surrounding spaces, uppercases, then replaces every
//     interior space with "-"; it also reports HasPrefix against prefix.
//  4. An empty joined string splits into one empty part, not zero parts.
//
// TODO: implement matchText, repeatJoin and normalizeTitle. Keep main() as
// the demo harness.
package main

import "fmt"

// TextMatch is the result of searching text for one substring.
type TextMatch struct {
	Found bool
	Count int
	Index int
}

// matchText reports whether sub occurs in text, how many non-overlapping
// times, and where it first appears.
func matchText(text, sub string) TextMatch {
	// TODO
	return TextMatch{}
}

// repeatJoin repeats one word, joins a list and splits the join back apart.
func repeatJoin(word string, times int, words []string, sep string) (repeated, joined string, parts []string) {
	// TODO
	return "", "", nil
}

// normalizeTitle returns the display form of title and whether it starts with
// prefix after normalization.
func normalizeTitle(title, prefix string) (normalized string, hasPrefix bool) {
	// TODO
	return "", false
}

func main() {
	fmt.Println(matchText("go gopher go", "go"))
	fmt.Println(repeatJoin("ha", 3, []string{"a", "b"}, "-"))
	fmt.Println(normalizeTitle("  go server  ", "GO-"))
}
