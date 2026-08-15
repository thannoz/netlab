# 06 — Reverse Proxy

## Worum es geht

Ein Proxy nimmt Requests entgegen und **reicht sie an einen anderen Server
(Backend) weiter** — die Antwort schickt er zurück an den Client. Der Client
merkt nicht, dass hinter dem Proxy noch jemand steht.

**Jetzt beantwortest du praktisch deine Frage „was ist ein Proxy?".**

## Warum es wichtig ist

Proxies sind das Fundament von Load Balancern, API Gateways, CDNs und
Service-Meshes. Wer einen Proxy gebaut hat, versteht moderne Infrastruktur.

## Was du baust

- Ein Server, der jeden eingehenden Request **an ein Backend weiterleitet**.
- Response des Backends zurück an den Client kopieren.
- Header korrekt behandeln (z. B. `X-Forwarded-For` setzen).
- **Stretch:** mehrere Backends nach Pfad routen, Request/Response loggen,
  Fehler des Backends abfangen.

## Kernkonzepte

- **Forward Proxy vs. Reverse Proxy** — der Unterschied
- Request lesen → an Backend `Dial`en → weiterleiten → Antwort zurückkopieren
- Welche Header ein Proxy setzen/entfernen sollte (Hop-by-Hop-Header!)
- Danach der Vergleich: Gos `net/http/httputil.ReverseProxy` — schau dir den
  Quellcode an, jetzt verstehst du ihn.

## Literatur

- **`net/http/httputil` Quellcode** (`ReverseProxy`) — überraschend lesbar und
  lehrreich, *nachdem* du deine eigene Version gebaut hast.
- **NGINX-Doku — „Reverse Proxy"** — wie es die Praxis erklärt.
- **Woodbeck, Kap. 9–10** — HTTP-Middleware und Proxying-Muster.
- **RFC 9110, Abschnitt zu Intermediaries** — was ein Proxy laut Standard tun darf.
