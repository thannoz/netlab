package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/thannoz/netlab/reverse-proxy/proxy"
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

func main() {
	targetService := flag.String("target", "http://localhost:9001", "target service")
	proxyService := flag.String("listen", "8080", "proxy service")
	flag.Parse()

	u, err := url.Parse(*targetService)
	if err != nil {
		log.Printf("could not parse url: %v", err)
		return
	}
	if err := checkSchemeHost(u); err != nil {
		log.Printf("invalid url: %v", err)
		return
	}

	p := proxy.New(u, log.New(os.Stderr, "proxy: ", log.LstdFlags))

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", *proxyService),
		Handler:           p,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Proxy service listening on %s", *proxyService)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
