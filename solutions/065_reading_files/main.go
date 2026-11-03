// 065: Reading Files. os.ReadFile loads a whole file into memory and returns
// the bytes plus an error; callers decide what to do with the error.
//
// Scenario: a config loader prints a small manifest file. Missing files must
// surface as errors, while an empty file is valid and yields "".
//
// Rules:
//  1. readText returns the exact file contents as a string, byte for byte.
//  2. An empty file returns "" and a nil error.
//  3. A missing path returns "" and a non-nil error.
//  4. Never call os.Exit or log.Fatal; return errors to the caller.
package main

import (
	"fmt"
	"os"
)

// readText returns the contents of the file at path.
func readText(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	text, err := readText("go.mod")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println(text)
}
