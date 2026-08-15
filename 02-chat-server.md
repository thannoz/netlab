# 02 — Concurrent Chat Server

## Worum es geht

Ein Server, mit dem sich **viele Clients gleichzeitig** verbinden. Was einer
schreibt, sehen alle anderen (Broadcast). Wie ein winziges IRC.

## Warum es wichtig ist

Hier trifft Networking auf **Nebenläufigkeit** — Gos Superkraft. Du lernst, wie
ein Server tausende Verbindungen parallel bedient, ohne für jede einen eigenen
Betriebssystem-Thread zu brauchen. Das ist das Herz jedes echten Servers.

## Was du baust

- Ein Server, der pro Verbindung eine **Goroutine** startet.
- Ein **Broadcast**: eine Nachricht an alle verbundenen Clients verteilen.
- Sauberes Handling von Beitritt/Verlassen (Client trennt die Verbindung).
- **Stretch:** Nicknames, Räume/Channels, private Nachrichten.

## Kernkonzepte

- **Goroutine pro Verbindung** — `go handleConn(conn)`
- **Channels** als sichere Kommunikation zwischen Goroutines (statt geteilter
  Variablen mit Locks)
- Das **Broadcast-Muster:** eine zentrale Goroutine hält die Client-Liste
- **Race Conditions** und wie man sie vermeidet (`go test -race`!)
- Ressourcen aufräumen: `defer conn.Close()`

## Literatur

- **The Go Programming Language, Kap. 8.10** — enthält einen kompletten
  Chat-Server. Fast eine Blaupause für dieses Projekt.
- **Katherine Cox-Buday — „Concurrency in Go"** — das Standardwerk zu
  Goroutines, Channels und nebenläufigen Mustern.
- **Go Blog:** „Share Memory By Communicating" und „Go Concurrency Patterns".
- **Woodbeck, Kap. 3** — Verbindungen nebenläufig annehmen.
