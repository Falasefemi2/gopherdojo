// 081: Reference solution. Read this only after your own attempt passes.
package main

import (
	"context"
	"fmt"
	"time"
)

func fetch(ctx context.Context, delay time.Duration) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(delay):
		return "payload", nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	out, err := fetch(ctx, 100*time.Millisecond)
	fmt.Println(out, err)
}
