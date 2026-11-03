// 031: Channel Synchronization. A done channel announces that a goroutine has
// finished; closing it wakes every waiter at once.
//
// Scenario: a nightly batch must finish before the summary report is printed.
//
// Rules:
//  1. waitFor starts work in a goroutine, closes done after work returns, and
//     returns work's result only after done is closed.
//  2. work must run exactly once per call.
//  3. signalWhenDone closes done after work returns; it never sends on done.
//  4. A closed done channel is ready immediately for every later receive.
package main

import "fmt"

// waitFor runs work in a goroutine and blocks until work has completed.
func waitFor(work func() int) int {
	done := make(chan struct{})
	var result int
	go func() {
		result = work()
		close(done)
	}()
	<-done
	return result
}

// signalWhenDone runs work in a goroutine and closes done once work returns.
func signalWhenDone(work func(), done chan struct{}) {
	go func() {
		work()
		close(done)
	}()
}

func main() {
	fmt.Println(waitFor(func() int { return 21 * 2 }))

	done := make(chan struct{})
	signalWhenDone(func() { fmt.Println("job finished") }, done)
	<-done
}
