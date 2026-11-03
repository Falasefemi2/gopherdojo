// 082: Reference solution. Read this only after your own attempt passes.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func goVersion() (string, error) {
	cmd := exec.Command("go", "version")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func main() {
	v, err := goVersion()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
}
