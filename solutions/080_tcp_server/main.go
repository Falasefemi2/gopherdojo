// 080: Reference solution. Read this only after your own attempt passes.
package main

import (
	"bufio"
	"fmt"
	"net"
)

func serveOnce(ln net.Listener) error {
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	if _, err := conn.Write([]byte(line)); err != nil {
		return err
	}
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
