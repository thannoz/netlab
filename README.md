# netlab — Network programming in Go, built from scratch

A learning journey: from raw TCP sockets to an API gateway — **every feature
implemented by hand**, to truly understand what happens underneath.

> Guiding principle: I build each project myself. I only reach for libraries
> once I understand what they do internally.

## The through-line

Each stage builds on the previous one:

| # | Project | New core idea | Matches interest |
|---|---------|---------------|------------------|
| 01 | TCP Echo | Sockets: connect, read/write bytes | Networking basics |
| 02 | Chat server | Many clients at once | Concurrency |
| 03 | Protocol & framing | Where does a message end? | Low-level, bytes |
| 04 | HTTP by hand | HTTP is just text over TCP | "under the hood" |
| 05 | File server | Stream & serve files | Files, I/O |
| 06 | Reverse proxy | Forwarding requests | Proxy |
| 07 | Load balancer | Spread load across N backends | Load balancer |
| 08 | API gateway | Routing + middleware | API gateway, servers |

At the end: **capstones** — the big final projects → [capstones.md](capstones.md)

## Foundational reading (companion to every stage)

- **Alan Donovan & Brian Kernighan — "The Go Programming Language"** — the Go book.
  Ch. 8 (Goroutines & Channels) even includes a chat server.
- **Adam Woodbeck — "Network Programming with Go"** (Manning, 2021) — THE book for
  exactly this journey. Sockets, TCP, HTTP, proxies — all in Go.
- **Beej's Guide to Network Programming** (free online) — the classic on sockets.
  Written in C, but the concepts apply everywhere.
- **Kurose & Ross — "Computer Networking: A Top-Down Approach"** — the theory
  behind it all (layering model, TCP/IP, HTTP).

## Meta resources

- **"Build Your Own X"** (GitHub list) — guides to rebuilding X yourself.
- **CodeCrafters** (paid) — guided "Build your own Redis/HTTP/…" challenges.

## Structure

Each stage lives in its own subdirectory with its own `main` package, under the
module `github.com/thannoz/netlab`.
