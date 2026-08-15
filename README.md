# netlab — Netzwerk-Programmierung in Go, von Grund auf selbst gebaut

Eine Lern-Reise: von rohen TCP-Sockets bis zu einem API Gateway — **jedes Feature
selbst implementiert**, um im Kern zu verstehen, was darunter passiert.

> Leitprinzip: Ich baue jedes Projekt selbst. Bibliotheken benutze ich erst,
> wenn ich verstanden habe, was sie intern tun.

## Der rote Faden

Jede Stufe baut auf der vorigen auf:

| # | Projekt | Neue Kernidee | Trifft Interesse |
|---|---------|---------------|------------------|
| 01 | [TCP Echo](01-tcp-echo.md) | Sockets: Verbindung, Bytes lesen/schreiben | Netzwerk-Basis |
| 02 | [Chat-Server](02-chat-server.md) | Viele Clients gleichzeitig | Nebenläufigkeit |
| 03 | [Protokoll & Framing](03-protocol-framing.md) | Wo endet eine Nachricht? | Low-level, Bytes |
| 04 | [HTTP von Hand](04-http-from-scratch.md) | HTTP ist nur Text über TCP | „unter der Haube" |
| 05 | [File-Server](05-file-server.md) | Dateien streamen & ausliefern | Files, I/O |
| 06 | [Reverse Proxy](06-reverse-proxy.md) | Requests weiterreichen | Proxy |
| 07 | [Load Balancer](07-load-balancer.md) | Last auf N Backends verteilen | Load Balancer |
| 08 | [API Gateway](08-api-gateway.md) | Routing + Middleware | API Gateway, Server |

Am Ende: **Capstones** — die großen End-Projekte → [capstones.md](capstones.md)

## Grundlagen-Literatur (begleitend zu allen Stufen)

- **Alan Donovan & Brian Kernighan — „The Go Programming Language"** — das Go-Buch.
  Kap. 8 (Goroutines & Channels) enthält sogar einen Chat-Server.
- **Adam Woodbeck — „Network Programming with Go"** (Manning, 2021) — DAS Buch für
  genau diese Reise. Sockets, TCP, HTTP, Proxies — alles in Go.
- **Beej's Guide to Network Programming** (kostenlos online) — der Klassiker zu
  Sockets. In C geschrieben, aber die Konzepte gelten überall.
- **Kurose & Ross — „Computer Networking: A Top-Down Approach"** — die Theorie
  hinter allem (Schichtenmodell, TCP/IP, HTTP).

## Meta-Ressourcen

- **„Build Your Own X"** (GitHub-Liste) — Anleitungen, X selbst nachzubauen.
- **CodeCrafters** (kostenpflichtig) — geführte „Build your own Redis/HTTP/…"-Challenges.

## Struktur

Jede Stufe lebt in ihrem eigenen Unterordner mit eigenem `main`-Paket unter dem
Modul `github.com/thannoz/netlab`.
