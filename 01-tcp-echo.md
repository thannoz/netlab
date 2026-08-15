# 01 — TCP Echo Server + Client

## Worum es geht

Zwei Programme reden über einen **TCP-Socket**: Der Client schickt Text, der
Server schickt ihn *zurück* („echo"). Das ist das „Hallo Welt" der
Netzwerk-Programmierung — und der Moment, in dem zwei Prozesse zum ersten Mal
über das Netzwerk kommunizieren.

## Warum es wichtig ist

Jede höhere Technik (HTTP, Proxy, Datenbanken) sitzt auf genau diesem Fundament:
eine Verbindung öffnen, Bytes rein, Bytes raus. Wer das versteht, versteht die
Basis von *allem*.

## Was du baust

- **Server:** lauscht auf einem Port, nimmt Verbindungen an, liest Daten und
  schickt sie unverändert zurück.
- **Client:** verbindet sich, sendet, was du tippst, und zeigt die Antwort.
- **Stretch:** mehrere Clients gleichzeitig (führt direkt zu Stufe 02).

## Kernkonzepte

- `net.Listen("tcp", ...)` und `net.Dial("tcp", ...)`
- Was eine `net.Conn` ist (ein `io.Reader` UND `io.Writer`)
- Ports, Adressen, `localhost` / `127.0.0.1`
- Der Unterschied zwischen „Verbindung annehmen" (`Accept`) und „Daten lesen"
- TCP ist ein **Bytestrom** — kein „Nachrichten"-Konzept (wichtig für Stufe 03!)

## Literatur

- **Beej's Guide to Network Programming** — Kapitel zu `socket`, `bind`,
  `listen`, `accept`, `connect`. Die mentale Grundlage.
- **Woodbeck — „Network Programming with Go", Kap. 3** (Reliable TCP Data Streams).
- **Go-Doku:** das [`net`-Package](https://pkg.go.dev/net), besonders `Listener`
  und `Conn`.
- **The Go Programming Language, Kap. 8.1–8.2** — ein Echo-/Clock-Server als Beispiel.
