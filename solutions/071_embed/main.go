// 071: Reference solution. Read this only after your own attempt passes.
package main

import (
	_ "embed"
	"fmt"
)

//go:embed data/greeting.txt
var greetingFile string

func greeting() string {
	return greetingFile
}

func main() {
	fmt.Print(greeting())
}
