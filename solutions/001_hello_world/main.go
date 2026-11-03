// 001: Reference solution. Read this only after your own attempt passes.
package main

import "fmt"

func greet(app string) string {
	if app == "" {
		return "Hello, stranger!"
	}
	return fmt.Sprintf("Hello, %s!", app)
}

func quoted(s string) string {
	return fmt.Sprintf("%q", s)
}

func banner(app, version string) string {
	return fmt.Sprintf("=== %s %s ===\nDeploying to production\nHave a nice day!", app, version)
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
