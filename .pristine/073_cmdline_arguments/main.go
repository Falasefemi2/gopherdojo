// 073: Command-Line Arguments. os.Args[0] is the program name; everything after
// it is user input and arrives as strings.
//
// Scenario: a tiny invoice tool adds up amounts typed on the command line.
//
// Rules:
//  1. sumArgs receives only the arguments after the program name.
//  2. Each argument is trimmed of surrounding spaces and parsed as a base-10
//     integer; negative numbers are allowed.
//  3. No arguments is a valid call and returns 0 with a nil error.
//  4. Any unparseable argument returns 0 and a non-nil error naming the value.
//
// TODO: implement sumArgs below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"os"
)

// sumArgs adds the numeric arguments and returns their total.
func sumArgs(args []string) (int, error) {
	// TODO
	return 0, nil
}

func main() {
	total, err := sumArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	fmt.Println("total:", total)
}
