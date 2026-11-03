// 077: Logging. log.New(w, prefix, flags) builds a logger that writes to any
// io.Writer; flags 0 means no timestamp, which keeps output testable.
//
// Scenario: a worker records informational and error events into one sink.
//
// Rules:
//  1. logInfo writes exactly "INFO: <msg>" followed by one newline.
//  2. logError writes exactly "ERROR: <msg>" followed by one newline.
//  3. Both build their own log.New(w, "", 0); no timestamps, no file names.
//  4. A trailing newline already present in msg must not produce a blank line.
//
// TODO: implement logInfo and logError below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"io"
	"os"
)

// logInfo writes an INFO line to w.
func logInfo(w io.Writer, msg string) {
	// TODO
}

// logError writes an ERROR line to w.
func logError(w io.Writer, msg string) {
	// TODO
}

func main() {
	fmt.Println("events:")
	logInfo(os.Stdout, "worker started")
	logError(os.Stdout, "queue empty")
}
