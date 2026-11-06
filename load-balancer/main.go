package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	lb "github.com/thannoz/netlab/load-balancer/lb"
)

func run() error {
	port := flag.Int("listen", 8090, "Load Balancer port")
	backends := flag.String("backends", "http://localhost:9001,http://localhost:9002", "Backend service to load balance to")
	flag.Parse()

	urls := strings.Split(*backends, ",")
	backendURLs, err := lb.ParseURLs(urls...)
	if err != nil {
		return fmt.Errorf("parse backends: %w", err)

	}

	balancer := lb.New(backendURLs, log.New(os.Stderr, "lb: ", log.LstdFlags))
	go balancer.HealthCheck(2 * time.Second)

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		Handler:           balancer,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Load Balancer listening on port %d", *port)

	return srv.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
