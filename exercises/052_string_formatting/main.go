// 052: String Formatting. fmt verbs control how values render; %v is the
// default form and the extra flags reveal structure and type.
//
// Scenario: an operations dashboard renders a point, a count and a price for
// both humans and log parsers.
//
// Rules:
//  1. Value uses %v, Plus uses %+v, GoSyntax uses %#v and Type uses %T.
//  2. Int uses %d, Float uses %.2f (always two decimals) and Quoted uses %q.
//  3. Label uses %s.
//  4. Every field must match fmt.Sprintf exactly; do not hard-code results.
//
// TODO: implement formatAll. Keep main() as the demo harness.
package main

import "fmt"

// Point is a coordinate pair used to show struct formatting.
type Point struct {
	X, Y int
}

// Formatted holds one rendering per formatting verb.
type Formatted struct {
	Value    string
	Plus     string
	GoSyntax string
	Type     string
	Int      string
	Float    string
	Quoted   string
	Label    string
}

// formatAll renders the same inputs through the verbs named in the rules.
func formatAll(label string, p Point, n int, f float64, s string) Formatted {
	// TODO
	return Formatted{}
}

func main() {
	fmt.Printf("%+v\n", formatAll("demo", Point{X: 1, Y: 2}, 7, 3.14159, "go"))
}
