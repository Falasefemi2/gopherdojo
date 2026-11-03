// 082: Spawning Processes. exec.Command builds a process; Start launches it
// without waiting, and Wait reaps it and reports the exit status.
//
// Scenario: a build tool records which Go toolchain produced an artefact.
//
// Rules:
//  1. goVersion runs exactly "go version" with no shell.
//  2. It uses Start followed by Wait, not Run or Output.
//  3. stdout is captured; the returned string has no trailing whitespace.
//  4. Start or Wait failures return an error and an empty result.
//
// TODO: implement goVersion below. Keep main() as the demo harness.
package main

import "fmt"

// goVersion returns the output of "go version" without a trailing newline.
func goVersion() (string, error) {
	// TODO
	return "", nil
}

func main() {
	v, err := goVersion()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
}
