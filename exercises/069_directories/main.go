// 069: Directories. os.MkdirAll creates a directory and every missing parent,
// and os.ReadDir lists entries sorted by name.
//
// Scenario: a workspace bootstrap creates a nested folder layout, then prints
// the top-level entries so the user can confirm what was made.
//
// Rules:
//  1. makeTree creates each relative dir under base, including parents.
//  2. Calling makeTree twice with the same arguments must succeed (idempotent).
//  3. listNames returns entry names only, in the sorted order os.ReadDir gives.
//  4. listNames returns a non-nil error when dir cannot be read.
//
// TODO: implement makeTree and listNames below. Keep main() as the demo.
package main

import (
	"fmt"
	"os"
)

// makeTree creates every relative dir below base.
func makeTree(base string, dirs ...string) error {
	// TODO
	return nil
}

// listNames returns the sorted entry names in dir.
func listNames(dir string) ([]string, error) {
	// TODO
	return nil, nil
}

func main() {
	base, err := os.MkdirTemp("", "dojo-dirs-")
	if err != nil {
		fmt.Println("temp error:", err)
		return
	}
	defer os.RemoveAll(base)

	if err := makeTree(base, "alpha", "beta"); err != nil {
		fmt.Println("mkdir error:", err)
		return
	}
	names, err := listNames(base)
	fmt.Println("entries:", names, "err:", err)
}
