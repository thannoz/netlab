package lb

import (
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/thannoz/netlab/reverse-proxy/proxy"
)

type Backend struct {
	URL *url.URL
	p   *proxy.Proxy
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

func (lb *LoadBalancer) nextBackend() *Backend {
	n := lb.next.Add(1)
	index := int(n % uint64(len(lb.backends)))

	return lb.backends[index]
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.nextBackend()
	backend.p.ServeHTTP(w, r)
}

func newBackend(u *url.URL, logger proxy.Logger) *Backend {
	p := proxy.New(u, logger, nil)

	return &Backend{
		URL: u,
		p:   p,
	}
}
