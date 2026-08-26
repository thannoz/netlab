package main

import (
	"bytes"
	"io"
	"net"
	"testing"
)

func TestEchoServer(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go handleConn(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	payload := []byte("hello netlab")
	if _, err := conn.Write(payload); err != nil {
		t.Fatal(err)

	}

	got := make([]byte, len(payload))
	if _, err := io.ReadFull(conn, got); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(payload, got) {
		t.Errorf("Echo falsch: gesendet %q, empfangen %q", payload, got)
	}
}
