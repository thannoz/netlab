# Capstones — the big final projects

Once the foundation (stages 01–08) is solid, you can aim at genuinely impressive
projects. Each of these is portfolio-worthy and deserves its own repo.

---

## 🗄️ Mini-Redis — in-memory key-value store

Your own Redis clone: `SET`, `GET`, `DEL` over a custom protocol, many
concurrent clients, optional persistence to disk.

**Uses:** networking (01) + concurrency (02) + protocol/framing (03) + files (05).

**Reading:**
- **RESP protocol specification** (redis.io → "Protocol specification") — Redis'
  wire protocol, surprisingly simple.
- **Josiah Carlson — "Redis in Action"** — how Redis thinks.
- **CodeCrafters "Build Your Own Redis"** — guided challenge (Go track).

---

## 🧲 BitTorrent client

Real **peer-to-peer**: read a `.torrent` file, ask the tracker, connect to other
peers and swap file pieces until the file is complete.

**Uses:** everything — sockets, protocol, concurrency, files. The classic for
"I really understand networks now," and enormous fun.

**Reading:**
- **Jesse Li — "Building a BitTorrent client from the ground up in Go"** (blog post)
  — the perfect, complete walkthrough.
- **BitTorrent specification BEP 3** (bittorrent.org) — the official protocol.

---

## 🔀 NGINX-lite — reverse proxy + load balancer

Your stages 06+07 built out into something usable: config file, TLS/HTTPS,
health checks, metrics, access logs.

**Reading:**
- **NGINX docs** (architecture, reverse proxy, load balancing).
- **Woodbeck, ch. 9–10** + Go `crypto/tls`.

---

## 🐳 Mini-Docker (later — needs Linux)

Containers from scratch: process isolation with **Linux namespaces**, resource
limits with **cgroups**, your own root filesystem. Needs Linux (on a Mac: a VM).
You'll already have the networking down by then.

**Reading:**
- **Liz Rice — "Containers From Scratch"** (talk + code, in Go).
- **Liz Rice — "Container Security"** (book).
- Article: "Building a container in less than 100 lines of Go".

---

## 🌐 Distributed KV store with replication (mini-etcd) — the final boss

Multiple servers that agree on the same state despite failures — **consensus**
with the Raft algorithm. Distributed systems in their purest form.

**Reading:**
- **Diego Ongaro & John Ousterhout — "In Search of an Understandable Consensus
  Algorithm (Raft)"** — the paper. Plus raft.github.io (visualization!).
- **Martin Kleppmann — "Designing Data-Intensive Applications"** — the best book
  on distributed systems, period.
- **MIT 6.824 Distributed Systems** — lectures + labs, free online.
