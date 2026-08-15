# 04 — HTTP/1.1 Server von Hand

## Worum es geht

Ein HTTP-Server, der Requests **selbst parst** — ohne `net/http`. Du liest die
rohen Bytes vom TCP-Socket und interpretierst sie als HTTP: Methode, Pfad,
Header, Body. Dann baust du eine gültige HTTP-Antwort von Hand.

## Warum es wichtig ist

Danach ist `net/http` keine Blackbox mehr. Du weißt, was ein „Request" wirklich
ist: **Text nach festen Regeln über TCP**. Das ist die Grundlage für Proxy,
Load Balancer und Gateway (Stufen 06–08).

## Was du baust

- Einen TCP-Server (Stufe 01), der eingehende Bytes als **HTTP-Request** parst.
- Request-Zeile (`GET /path HTTP/1.1`), Header, optionaler Body.
- Eine korrekte **HTTP-Response** (Statuszeile, Header, Body).
- **Stretch:** mehrere Routen, `Content-Length`, `keep-alive`, einfache Middleware.

## Kernkonzepte

- **Anatomie eines HTTP-Requests/Response** (Startzeile, Header, Leerzeile, Body)
- Warum HTTP „stateless" ist
- `Content-Length` vs. `chunked` — woher weiß der Server, wo der Body endet?
  (Direkter Bezug zu Stufe 03: Framing!)
- Statuscodes, gängige Header
- Danach: Vergleich mit dem echten `net/http` — was nimmt es dir ab?

## Literatur

- **RFC 9112 (HTTP/1.1 Message Syntax)** und **RFC 9110 (HTTP Semantics)** — die
  offiziellen Regeln. Trocken, aber DIE Quelle.
- **MDN Web Docs — „HTTP"** — die zugängliche Erklärung von Requests, Methoden,
  Statuscodes, Headern.
- **Woodbeck, Kap. 8–9** — HTTP-Clients und -Server in Go.
- **David Gourley & Brian Totty — „HTTP: The Definitive Guide"** — falls du HTTP
  richtig tief verstehen willst.
