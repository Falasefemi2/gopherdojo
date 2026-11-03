// 084: Signals. os.Signal values notify a program about external events;
// signal.Notify delivers them on a channel so handling stays ordinary Go.
//
// Scenario: a daemon prints which shutdown signal it received.
//
// Rules:
//  1. signalLabel maps os.Interrupt to "interrupt", SIGTERM to "terminate",
//     SIGKILL to "kill" and SIGHUP to "hup".
//  2. nil or any unknown signal returns "unknown"; signalLabel never panics.
//  3. waitForSignal returns the first value received from ch.
//  4. waitForSignal returns nil when ch is closed.
//
// TODO: implement signalLabel and waitForSignal below. Keep main() as the demo
// harness; never send real signals from a test.
package main

import (
	"fmt"
	"os"
	"os/signal"
)

// signalLabel returns a short human label for sig.
func signalLabel(sig os.Signal) string {
	// TODO
	return ""
}

// waitForSignal blocks until one signal arrives on ch.
func waitForSignal(ch <-chan os.Signal) os.Signal {
	// TODO
	return nil
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	fmt.Println("press Ctrl+C")
	sig := waitForSignal(ch)
	fmt.Println("received:", signalLabel(sig))
}
