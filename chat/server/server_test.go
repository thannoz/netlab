package server

import (
	"bytes"
	"errors"
	"fmt"
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

func TestHandleConnReadOneMessage(t *testing.T) {
	serverA, clientA := net.Pipe()
	serverB, clientB := net.Pipe()
	server := New(nil)

	defer func() {
		serverA.Close()
		serverB.Close()
		clientB.Close()
	}()

	server.add(serverB)

	buf := make([]byte, 6)
	gotMessage := make(chan string, 1)
	errs := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		_, err := io.ReadFull(clientB, buf)
		if err != nil {
			errs <- err
			return
		}
		gotMessage <- string(buf)
	}()

	go func() {
		server.handleConn(serverA)
		close(done)
	}()

	_, err := clientA.Write([]byte("hello\n"))
	if err != nil {
		t.Fatalf("writing sender message: %v", err)
	}

	wantMessage := "hello\n"
	select {
	case err := <-errs:
		t.Fatalf("reading recipient message: %v", err)
	case message := <-gotMessage:
		if message != wantMessage {
			t.Errorf("message: got %q, want %q", message, wantMessage)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for recipient message")
	}

	if err := clientA.Close(); err != nil {
		t.Fatalf("closing sender: %v", err)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("handleConn did not terminate")
	}
}

func TestHandleConnReadsMultipleMessages(t *testing.T) {
	serverA, clientA := net.Pipe()
	serverB, clientB := net.Pipe()
	defer func() {
		serverA.Close()
		serverB.Close()
		clientA.Close()
		clientB.Close()
	}()

	server := New(nil)
	server.add(serverB)

	want := []byte("hello\nworld\n")
	gotMessage := make(chan []byte, 1)
	readErr := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		buf := make([]byte, len(want))
		_, err := io.ReadFull(clientB, buf)
		if err != nil {
			readErr <- err
			return
		}
		gotMessage <- buf
	}()

	go func() {
		server.handleConn(serverA)
		close(done)
	}()

	if _, err := clientA.Write(want); err != nil {
		t.Fatalf("writing messages: %v", err)
	}
	_ = clientA.Close()

	select {
	case err := <-readErr:
		t.Fatalf("reading messages: %v", err)
	case message := <-gotMessage:
		if !bytes.Equal(message, want) {
			t.Errorf("messages: got %q, want %q", message, want)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for messages")
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("handleConn did not terminate")
	}
}

func TestHandleConnRemovesConnection(t *testing.T) {
	serverA, clientA := net.Pipe()
	defer func() {
		serverA.Close()
		clientA.Close()
	}()

	server := New(nil)
	done := make(chan struct{})

	go func() {
		server.handleConn(serverA)
		close(done)
	}()

	_ = clientA.Close()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("handleConn did not terminate")
	}

	if _, exists := server.clients[serverA]; exists {
		t.Error("server.clients still contains the closed connection")
	}
}

type trackingConn struct {
	net.Conn
	closed bool
}

func (c *trackingConn) Close() error {
	c.closed = true
	return c.Conn.Close()
}

func TestHandleConnClosesConnection(t *testing.T) {
	pipeConn, peer := net.Pipe()
	conn := &trackingConn{Conn: pipeConn}
	defer peer.Close()

	server := New(nil)
	done := make(chan struct{})

	go func() {
		server.handleConn(conn)
		close(done)
	}()

	_ = peer.Close()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("handleConn did not terminate")
	}

	if !conn.closed {
		t.Error("connection was not closed by handleConn")
	}
}

type recordingLogger struct {
	messages []string
}

func (l *recordingLogger) Printf(format string, args ...any) {
	l.messages = append(l.messages, fmt.Sprintf(format, args...))
}

type failingConn struct {
	net.Conn
	err error
}

func (c *failingConn) Read([]byte) (int, error) {
	return 0, c.err
}

func TestHandleConnLogsReadError(t *testing.T) {
	pipeConn, peer := net.Pipe()
	defer peer.Close()

	logger := &recordingLogger{}
	server := New(logger)
	conn := &failingConn{Conn: pipeConn, err: errors.New("read failed")}
	done := make(chan struct{})

	go func() {
		server.handleConn(conn)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("handleConn did not terminate")
	}

	if len(logger.messages) != 1 {
		t.Fatalf("logged messages: got %d, want 1", len(logger.messages))
	}

	want := "scan error read failed"
	if logger.messages[0] != want {
		t.Errorf("logged message: got %q, want %q", logger.messages[0], want)
	}
}
