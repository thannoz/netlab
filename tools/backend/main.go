package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"time"
)

const maxBodySize = 10 << 20

type Server struct {
	serviceName string
}

func printSorted(w io.Writer, m map[string][]string) {
	for _, key := range slices.Sorted(maps.Keys(m)) {
		fmt.Fprintf(w, "%s: %s\n", key, m[key])
	}
}

func statusFromQuery(r *http.Request) int {
	raw := r.URL.Query().Get("status")
	if raw == "" {
		return http.StatusOK
	}
	code, err := strconv.Atoi(raw)
	if err != nil || code < 200 || code > 599 {
		log.Printf("invalid status %q, using %d", raw, http.StatusOK)
		return http.StatusOK
	}
	return code
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxBodySize))
	if err != nil {
		log.Printf("%s: could not read body from %s: %v", s.serviceName, r.RemoteAddr, err)
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(statusFromQuery(r))

	fmt.Fprintf(w, "Server name: %s\n", s.serviceName)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Addr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Body length: %d\n", len(body))

	fmt.Fprintln(w, "\nHeaders:")
	printSorted(w, r.Header)

	fmt.Fprintln(w, "\nQuery params:")
	printSorted(w, r.URL.Query())
}

func main() {
	port := flag.Int("port", 9001, "port to listen on")
	name := flag.String("name", "backend-server", "name of the backend service")
	flag.Parse()

	server := &Server{serviceName: *name}

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.echo)
	mux.HandleFunc("/health", server.health)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("%s listening on %s", *name, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
