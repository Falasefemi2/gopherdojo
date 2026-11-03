package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestServeOnceEchoesLine(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	t.Cleanup(func() { ln.Close() })

	errCh := make(chan error, 1)
	go func() { errCh <- serveOnce(ln) }()

	conn, err := net.DialTimeout("tcp", ln.Addr().String(), 2*time.Second)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("set deadline: %v", err)
	}
	if _, err := conn.Write([]byte("ping\n")); err != nil {
		t.Fatalf("write: %v", err)
	}

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatalf("read echo: %v", err)
	}
	if line != "ping\n" {
		t.Errorf("echo = %q, want %q", line, "ping\n")
	}
	conn.Close()

	select {
	case err := <-errCh:
		if err != nil {
			t.Errorf("serveOnce() error = %v, want nil", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("serveOnce did not return after the connection closed")
	}
}

func TestServeOnceAcceptError(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	ln.Close()

	if err := serveOnce(ln); err == nil {
		t.Fatal("serveOnce() on a closed listener = nil, want an error")
	}
}
