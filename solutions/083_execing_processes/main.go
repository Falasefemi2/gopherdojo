// 083: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func goos() (string, error) {
	out, err := exec.Command("go", "env", "GOOS").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("go env GOOS: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func main() {
	v, err := goos()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("GOOS:", v, "runtime:", runtime.GOOS)
}
