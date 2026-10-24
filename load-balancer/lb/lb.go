package lb

import (
	"net/url"

	"github.com/thannoz/netlab/reverse-proxy/proxy"
)

type Backend struct {
	u *url.URL
	p *proxy.Proxy
}

type LoadBalancer struct {
	backend *Backend
}

func New(backend *Backend) *LoadBalancer {
	return &LoadBalancer{
		backend: backend,
	}
}

func newBackend(url *url.URL, proxy *proxy.Proxy) *Backend {
	return &Backend{
		u: url,
		p: proxy,
	}
}
