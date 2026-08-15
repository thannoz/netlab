# Capstones — die großen End-Projekte

Wenn die Basis (Stufen 01–08) sitzt, kannst du auf richtig beeindruckende
Projekte zielen. Jedes davon ist portfolio-tauglich und eigenes Repo wert.

---

## 🗄️ Mini-Redis — In-Memory Key-Value-Store

Ein eigener Redis-Klon: `SET`, `GET`, `DEL` über ein eigenes Protokoll, viele
gleichzeitige Clients, optional Persistenz auf Platte.

**Nutzt:** Networking (01) + Nebenläufigkeit (02) + Protokoll/Framing (03) + Files (05).

**Literatur:**
- **RESP-Protokoll-Spezifikation** (redis.io → „Protocol specification") — Redis'
  Draht-Protokoll, erstaunlich simpel.
- **Josiah Carlson — „Redis in Action"** — wie Redis denkt.
- **CodeCrafters „Build Your Own Redis"** — geführte Challenge (Go-Track).

---

## 🧲 BitTorrent-Client

Echtes **Peer-to-Peer**: eine `.torrent`-Datei lesen, den Tracker fragen, sich
mit anderen Peers verbinden und Datei-Stücke tauschen, bis die Datei komplett ist.

**Nutzt:** alles — Sockets, Protokoll, Nebenläufigkeit, Files. Der Klassiker für
„ich verstehe Netzwerke jetzt wirklich", und macht enorm Spaß.

**Literatur:**
- **Jesse Li — „Building a BitTorrent client from the ground up in Go"** (Blogpost)
  — die perfekte, komplette Anleitung.
- **BitTorrent-Spezifikation BEP 3** (bittorrent.org) — das offizielle Protokoll.

---

## 🔀 NGINX-lite — Reverse Proxy + Load Balancer

Deine Stufen 06+07 zu etwas Benutzbarem ausgebaut: Config-Datei, TLS/HTTPS,
Health-Checks, Metriken, Zugriffs-Logs.

**Literatur:**
- **NGINX-Doku** (Architektur, Reverse Proxy, Load Balancing).
- **Woodbeck, Kap. 9–10** + Go `crypto/tls`.

---

## 🐳 Mini-Docker (später — Linux nötig)

Container von Grund auf: Prozess-Isolation mit **Linux Namespaces**,
Ressourcen-Limits mit **cgroups**, eigenes Root-Filesystem. Braucht Linux (auf
dem Mac: VM). Networking hast du dann schon drauf.

**Literatur:**
- **Liz Rice — „Containers From Scratch"** (Vortrag + Code, in Go).
- **Liz Rice — „Container Security"** (Buch).
- Artikel: „Building a container in less than 100 lines of Go".

---

## 🌐 Verteilter KV-Store mit Replikation (mini-etcd) — der Endgegner

Mehrere Server, die sich trotz Ausfällen auf denselben Zustand einigen —
**Consensus** mit dem Raft-Algorithmus. Distributed Systems in Reinform.

**Literatur:**
- **Diego Ongaro & John Ousterhout — „In Search of an Understandable Consensus
  Algorithm (Raft)"** — das Paper. Plus raft.github.io (Visualisierung!).
- **Martin Kleppmann — „Designing Data-Intensive Applications"** — das beste Buch
  über verteilte Systeme, Punkt.
- **MIT 6.824 Distributed Systems** — Vorlesungen + Labs, kostenlos online.
