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
	client http.RoundTripper
}

func New(target *url.URL, logger Logger, client http.RoundTripper) *Proxy {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}
	if client == nil {
		client = http.DefaultTransport
	}
	return &Proxy{
		target: target,
		logger: logger,
		client: client,
	}
}

func (p *Proxy) forwardRequest(r *http.Request, u *url.URL) (*http.Response, error) {
	req, err := http.NewRequestWithContext(r.Context(), r.Method, u.String(), r.Body)
	if err != nil {
		return nil, err
	}
	req.Header = r.Header.Clone()

	resp, err := p.client.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("round trip: %w", err)
	}
	return resp, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := *p.target
	u.Path = r.URL.Path
	u.RawQuery = r.URL.RawQuery

	resp, err := p.forwardRequest(r, &u)
	if err != nil {
		p.logger.Printf("could not forward request %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			p.logger.Printf("close body: %v", err)
		}
	}()

	p.logger.Printf("%s %s %s -> %s", r.Method, r.RemoteAddr, r.URL, u.String())
	p.logger.Printf("Backend answered with status %d", resp.StatusCode)
	fmt.Fprintln(w, resp.Status)
}
