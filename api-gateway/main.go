package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thannoz/netlab/load-balancer/lb"
)

func run() error {
	mux := http.NewServeMux()

	port := flag.Int("listen", 8075, "API Gateway port")
	flag.Parse()

	users := []string{"http://localhost:9001", "http://localhost:9002"}
	orders := []string{"http://localhost:9011"}

	usersURLs, err := lb.ParseURLs(users...)
	if err != nil {
		return fmt.Errorf("parse users: %w", err)
	}

	ordersURLs, err := lb.ParseURLs(orders...)
	if err != nil {
		return fmt.Errorf("parse orders: %w", err)
	}

	usersLB := lb.New(usersURLs, log.New(os.Stderr, "API Gateway: ", log.LstdFlags))
	ordersLB := lb.New(ordersURLs, log.New(os.Stderr, "API Gateway: ", log.LstdFlags))

	mux.Handle("/users/", usersLB)
	mux.Handle("/orders/", ordersLB)

	go usersLB.HealthCheck(5 * time.Second)
	go ordersLB.HealthCheck(5 * time.Second)

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Printf("API Gateway listening on port %d", *port)
	return srv.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
