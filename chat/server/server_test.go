package server

import (
	"net"
	"testing"
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
