package proxy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"
)

type received struct {
	method   string
	path     string
	rawQuery string
	header   http.Header
	body     string
}

func newBackend(t *testing.T) (*url.URL, <-chan received) {
	t.Helper()

	got := make(chan received, 1)
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		got <- received{
			method:   r.Method,
			path:     r.URL.Path,
			rawQuery: r.URL.RawQuery,
			header:   r.Header.Clone(),
			body:     string(body),
		}
	}))
	t.Cleanup(backend.Close)

	target, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatalf("parse backend url: %v", err)
	}
	return target, got
}

func TestNewSetsDefaults(t *testing.T) {
	target, _ := url.Parse("http://localhost:9001")

	p := New(target, nil, nil)

	if p.logger == nil {
		t.Error("got nil logger, want default logger")
	}
	if p.client == nil {
		t.Error("got nil client, want http.DefaultTransport")
	}
}

func TestServeHTTPForwardsRequest(t *testing.T) {
	target, got := newBackend(t)
	p := New(target, nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/hallo?x=1&x=3", strings.NewReader("hallo"))
	req.Header.Add("X-A", "1")
	req.Header.Add("X-A", "2")
	rec := httptest.NewRecorder()

	p.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("got status %d, want %d", rec.Code, http.StatusOK)
	}

	r := <-got
	if r.method != http.MethodPost {
		t.Errorf("method: got %q, want %q", r.method, http.MethodPost)
	}
	if r.path != "/hallo" {
		t.Errorf("path: got %q, want %q", r.path, "/hallo")
	}
	if r.rawQuery != "x=1&x=3" {
		t.Errorf("query: got %q, want %q", r.rawQuery, "x=1&x=3")
	}
	if want := []string{"1", "2"}; !slices.Equal(r.header.Values("X-A"), want) {
		t.Errorf("header X-A: got %v, want %v", r.header.Values("X-A"), want)
	}
	if r.body != "hallo" {
		t.Errorf("body: got %q, want %q", r.body, "hallo")
	}
}

func TestServeHTTPKeepsTargetUnchanged(t *testing.T) {
	target, _ := newBackend(t)
	before := target.String()
	p := New(target, nil, nil)

	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/a/b?x=1", nil))

	if after := target.String(); after != before {
		t.Errorf("target changed: got %q, want %q", after, before)
	}
}

func TestServeHTTPBackendDown(t *testing.T) {
	backend := httptest.NewServer(http.NotFoundHandler())
	target, _ := url.Parse(backend.URL)
	backend.Close()
	p := New(target, nil, nil)

	rec := httptest.NewRecorder()
	p.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code < 400 {
		t.Errorf("got status %d, want an error status", rec.Code)
	}
}
