// 081: Context. A context carries cancellation and deadlines across API
// boundaries; select on ctx.Done() to stop early.
//
// Scenario: a fast cache and a slow origin answer the same lookup, but the
// caller's deadline wins.
//
// Rules:
//  1. fetch waits for delay, then returns "payload" and a nil error.
//  2. If ctx is done first, fetch returns "" and ctx.Err() without waiting.
//  3. An already-cancelled ctx returns immediately; no delay is observed.
//  4. fetch never ignores ctx.Done() and never leaks a goroutine.
//
// TODO: implement fetch below. Keep main() as the demo harness.
package main

import (
	"context"
	"fmt"
	"time"
)

// fetch simulates a lookup guarded by ctx.
func fetch(ctx context.Context, delay time.Duration) (string, error) {
	// TODO
	return "", nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	out, err := fetch(ctx, 100*time.Millisecond)
	fmt.Println(out, err)
}
