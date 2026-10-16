package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Logger interface {
	Printf(format string, args ...any)
}

type Proxy struct {
	target *url.URL
	logger Logger
}

func New(target *url.URL, logger Logger) *Proxy {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}
	return &Proxy{
		target: target,
		logger: logger,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := *p.target
	u.Path = r.URL.Path
	u.RawQuery = r.URL.RawQuery

	p.logger.Printf("%s %s -> %s", r.Method, r.URL, u.String())

	fmt.Fprintln(w, u.String())
}
