package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
)

type Logger interface {
	Printf(format string, args ...any)
}

type ChatServer struct {
	mu      sync.Mutex
	clients map[net.Conn]bool
	logger  Logger
}

func New(logger Logger) *ChatServer {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}
	return &ChatServer{
		clients: make(map[net.Conn]bool),
		logger:  logger,
	}
}

func (s *ChatServer) add(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[conn] = true
}

func (s *ChatServer) remove(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, conn)
}

func (s *ChatServer) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConn(conn)
	}
}

func (s *ChatServer) handleConn(conn net.Conn) {
	s.add(conn)
	defer func() {
		s.remove(conn)
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		s.broadcast(msg, conn)
	}
	if err := scanner.Err(); err != nil {
		s.logger.Printf("scan error %v", err)
	}

}

func (s *ChatServer) broadcast(msg string, sender net.Conn) {
	s.mu.Lock()
	recipients := make([]net.Conn, 0, len(s.clients))

	for cl := range s.clients {
		if cl != sender {
			recipients = append(recipients, cl)
		}
	}
	s.mu.Unlock()

	message := []byte(msg + "\n")
	for _, c := range recipients {
		if _, err := c.Write(message); err != nil {
			s.logger.Printf("write error: %v", err)
		}
	}
}
