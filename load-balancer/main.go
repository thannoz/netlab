package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	lb "github.com/thannoz/netlab/load-balancer/lb"
)

func checkSchemeHost(u *url.URL) error {
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("invalid url scheme: %s", u.String())
	}
	if u.Host == "" {
		return fmt.Errorf("invalid url host: %s", u.String())
	}
	return nil
}

func run() error {
	port := flag.Int("listen", 8090, "Load Balancer port")
	backends := flag.String("backends", "http://localhost:9001,http://localhost:9002", "Backend service to load balance to")
	flag.Parse()

	urls := strings.Split(*backends, ",")
	backendURLs := make([]*url.URL, 0, len(urls))

	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("parse url %q: %w", u, err)
		}
		if err := checkSchemeHost(parsedURL); err != nil {
			return fmt.Errorf("backend %q: %w", u, err)
		}
		backendURLs = append(backendURLs, parsedURL)
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
