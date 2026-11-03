// 001: Hello World. A Go program starts at package main and func main.
// fmt is how we build and print text: Sprintf returns a string,
// Println writes a line to stdout.
//
// TODO: implement greet, quoted and banner below so every check prints PASS
// and `go test` is green. The rules live in README.md and exercise_test.go.
package main

import "fmt"

// greet returns "Hello, <app>!".
// Edge case: an empty app returns "Hello, stranger!".
func greet(app string) string {
	return ""
}

// quoted returns s wrapped in double quotes, escaped the Go way.
func quoted(s string) string {
	return ""
}

// banner returns exactly three lines, joined by "\n", with no trailing newline:
//
//	=== <app> <version> ===
//	Deploying to production
//	Have a nice day!
func banner(app, version string) string {
	return ""
}

func check(label, got, want string) {
	if got == want {
		fmt.Printf("PASS %-8s %q\n", label, got)
		return
	}
	fmt.Printf("FAIL %-8s\n  expected: %q\n  got:      %q\n", label, want, got)
}

func main() {
	check("greet", greet("ledger-api"), "Hello, ledger-api!")
	check("greet", greet(""), "Hello, stranger!")

	check("quoted", quoted(`path "with" quotes`), `"path \"with\" quotes"`)

	check("banner", banner("ledger-api", "v1.4.2"),
		"=== ledger-api v1.4.2 ===\nDeploying to production\nHave a nice day!")
	check("banner", banner("payments", "v0.9.0"),
		"=== payments v0.9.0 ===\nDeploying to production\nHave a nice day!")

	fmt.Println(banner("ledger-api", "v1.4.2"))
}
