package lb

import (
	"net/http"
	"net/url"

	"github.com/thannoz/netlab/reverse-proxy/proxy"
)

type Backend struct {
	URL *url.URL
	p   *proxy.Proxy
}

type LoadBalancer struct {
	backends []*Backend
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
	lb.backends[0].p.ServeHTTP(w, r)
}

func newBackend(u *url.URL, logger proxy.Logger) *Backend {
	p := proxy.New(u, logger, nil)

	return &Backend{
		URL: u,
		p:   p,
	}
}
