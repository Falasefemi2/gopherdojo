// 076: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}

func main() {
	fmt.Println("env:", envOr("DOJO_ENV", "development"))
	fmt.Println("port:", envInt("PORT", 8080))
}
