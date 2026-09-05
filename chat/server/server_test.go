package server

import (
	"bytes"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

// White Box test for server package

func TestNewWithNilLogger(t *testing.T) {

	server := New(nil)
	if server == nil {
		t.Fatalf("got nil, want non-nil server")
	}

	got := server.logger == nil
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNewInitializesServer(t *testing.T) {
	var logger Logger

	server := New(logger)
	if server == nil {
		t.Fatalf("got nil, want non-nil server")
	}

	gotClientLen := len(server.clients)
	want := 0
	if gotClientLen != want {
		t.Errorf("(server.clients): got %v, want %v", gotClientLen, want)
	}

	gotClients := server.clients
	if gotClients == nil {
		t.Errorf("server.clients map: got nil, want initialized map")
	}
}

type dummyLogger struct {
	ID string
}

func (d *dummyLogger) Printf(format string, args ...any) {}

func TestNewUsesProvidedLogger(t *testing.T) {
	providedLogger := &dummyLogger{ID: "hello123"}
	server := New(providedLogger)
	if server == nil {
		t.Fatalf("got nil, want non-nil server")
	}

	if server.logger == nil {
		t.Errorf("server.logger: got nil, want non-nil logger %v", server.logger)
	}

	gotLogger := server.logger
	wantLogger := providedLogger
	if gotLogger != wantLogger {
		t.Errorf("server.logger does not reference the provide logger")
	}
}

func TestAddClient(t *testing.T) {
	server := New(nil)

	conn, peer := net.Pipe()
	defer func() {
		conn.Close()
		peer.Close()
	}()
	server.add(conn)

	gotClientsLen := len(server.clients)
	wantClientsLen := 1

	if gotClientsLen != wantClientsLen {
		t.Errorf("len(server.clients): got %d, want %d", gotClientsLen, wantClientsLen)
	}
	_, exists := server.clients[conn]
	wantExists := true
	if exists != wantExists {
		t.Errorf("client-key: got %v, want %v", exists, wantExists)
	}
}

func TestRemoveClient(t *testing.T) {
	server := New(nil)

	conn, peer := net.Pipe()
	defer func() {
		conn.Close()
		peer.Close()
	}()

	server.add(conn)
	server.remove(conn)

	gotClientLen := len(server.clients)
	wantClientLen := 0

	if gotClientLen != wantClientLen {
		t.Errorf("len(server.clients): got %d, want %d", gotClientLen, wantClientLen)
	}

	_, exists := server.clients[conn]
	wantExists := false
	if exists != wantExists {
		t.Errorf("client-key: got %v, want %v", exists, wantExists)
	}
}

func TestBroadCast(t *testing.T) {
	var wg sync.WaitGroup
	serverA, clientA := net.Pipe()
	serverB, clientB := net.Pipe()

	defer func() {
		serverA.Close()
		clientA.Close()
		serverB.Close()
		clientB.Close()
	}()

	server := New(nil)

	server.add(serverA)
	server.add(serverB)

	msg := "hello"

	buf := make([]byte, 15)
	var gotMessage string

	wg.Go(func() {
		n, _ := clientB.Read(buf)
		gotMessage = string(buf[:n])
	})
	server.broadcast(msg, serverA)

	wg.Wait()

	wantMessage := "hello\n"
	if gotMessage != wantMessage {
		t.Errorf("message: got %s, want %s", gotMessage, wantMessage)
	}
}

func TestBroadcastToMultipleRecipients(t *testing.T) {
	serverA, clientA := net.Pipe()
	serverB, clientB := net.Pipe()
	serverC, clientC := net.Pipe()
	defer func() {
		serverA.Close()
		clientA.Close()
		serverB.Close()
		clientB.Close()
		serverC.Close()
		clientC.Close()
	}()

	server := New(nil)
	server.add(serverA)
	server.add(serverB)
	server.add(serverC)

	want := []byte("hello\n")
	got := make(chan []byte, 2)
	errs := make(chan error, 2)
	read := func(conn net.Conn) {
		buf := make([]byte, len(want))
		_, err := io.ReadFull(conn, buf)
		if err != nil {
			errs <- err
			return
		}
		got <- buf
	}
	go read(clientB)
	go read(clientC)

	server.broadcast("hello", serverA)

	for range 2 {
		select {
		case err := <-errs:
			t.Fatalf("reading recipient message: %v", err)
		case message := <-got:
			if !bytes.Equal(message, want) {
				t.Errorf("message: got %q, want %q", message, want)
			}
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for recipient message")
		}
	}
}

func TestBroadcastWithoutRecipients(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer func() {
		serverConn.Close()
		clientConn.Close()
	}()

	server := New(nil)
	server.add(serverConn)

	done := make(chan struct{})
	go func() {
		server.broadcast("hello", serverConn)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("broadcast blocked without recipients")
	}
}

func TestBroadcastDoesNotSendToSender(t *testing.T) {
	sender, senderPeer := net.Pipe()
	recipient, recipientPeer := net.Pipe()
	defer func() {
		sender.Close()
		senderPeer.Close()
		recipient.Close()
		recipientPeer.Close()
	}()

	server := New(nil)
	server.add(sender)
	server.add(recipient)

	readDone := make(chan error, 1)
	go func() {
		buf := make([]byte, len("hello\n"))
		_, err := io.ReadFull(recipientPeer, buf)
		if err == nil && !bytes.Equal(buf, []byte("hello\n")) {
			err = io.ErrUnexpectedEOF
		}
		readDone <- err
	}()

	server.broadcast("hello", sender)
	if err := <-readDone; err != nil {
		t.Fatalf("reading recipient message: %v", err)
	}

	_ = senderPeer.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	buf := make([]byte, 1)
	_, err := senderPeer.Read(buf)
	if err == nil {
		t.Fatal("sender received a broadcast message")
	}
	if netErr, ok := err.(net.Error); !ok || !netErr.Timeout() {
		t.Fatalf("reading sender: got %v, want timeout", err)
	}
}
