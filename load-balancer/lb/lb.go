package lb

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/thannoz/netlab/reverse-proxy/proxy"
)

type Backend struct {
	URL   *url.URL
	p     *proxy.Proxy
	alive atomic.Bool
}

type LoadBalancer struct {
	backends []*Backend
	next     atomic.Uint64
}

func New(urls []*url.URL, logger proxy.Logger) *LoadBalancer {
	backends := make([]*Backend, 0, len(urls))
	for _, u := range urls {
		b := newBackend(u, logger)
		backends = append(backends, b)
	}

	return &LoadBalancer{
		backends: backends,
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.nextBackend()
	if backend == nil {
		http.Error(w, "no backend available", http.StatusServiceUnavailable)
		return
	}
	backend.p.ServeHTTP(w, r)
}

func (lb *LoadBalancer) nextBackend() *Backend {
	for range len(lb.backends) {
		n := lb.next.Add(1)
		index := int(n % uint64(len(lb.backends)))
		b := lb.backends[index]
		if b.alive.Load() {
			return b
		}
	}
	return nil
}

func newBackend(u *url.URL, logger proxy.Logger) *Backend {
	prx := proxy.New(u, logger, nil)

	b := &Backend{
		URL: u,
		p:   prx,
	}
	b.alive.Store(true)

	return b
}

func (lb *LoadBalancer) HealthCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, b := range lb.backends {
			alive := isAlive(b)
			b.alive.Store(alive)
		}
	}
}

func isAlive(b *Backend) bool {
	u := b.URL.JoinPath("health")
	cl := http.Client{Timeout: 2 * time.Second}

	resp, err := cl.Get(u.String())
	if err != nil {
		fmt.Printf("could not issue GET request: %v\n", err)
		return false
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("close body: %v\n", err)
		}
	}()

	return resp.StatusCode == http.StatusOK
}
