// 084: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalLabel(sig os.Signal) string {
	switch sig {
	case os.Interrupt:
		return "interrupt"
	case syscall.SIGTERM:
		return "terminate"
	case syscall.SIGKILL:
		return "kill"
	case syscall.SIGHUP:
		return "hup"
	default:
		return "unknown"
	}
}

func waitForSignal(ch <-chan os.Signal) os.Signal {
	return <-ch
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	fmt.Println("press Ctrl+C")
	sig := waitForSignal(ch)
	fmt.Println("received:", signalLabel(sig))
}
