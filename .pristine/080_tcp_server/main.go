// 080: TCP Server. net.Listen reserves a port; Accept blocks until one client
// connects, and net.Conn reads and writes bytes like a file.
//
// Scenario: a one-shot echo service used by a health probe.
//
// Rules:
//  1. serveOnce accepts exactly one connection; an Accept error is returned.
//  2. It reads one line terminated by '\n' and writes the same line back.
//  3. The accepted connection is closed before serveOnce returns; the listener
//     is left open for the caller.
//  4. Read or write failures are returned as errors.
//
// TODO: implement serveOnce below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"net"
)

// serveOnce accepts one connection, echoes one line and closes it.
func serveOnce(ln net.Listener) error {
	// TODO
	return nil
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	defer ln.Close()

	fmt.Println("listening on", ln.Addr())
	if err := serveOnce(ln); err != nil {
		fmt.Println("serve:", err)
	}
}
