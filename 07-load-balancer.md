# 07 — Load Balancer

## Worum es geht

Ein Load Balancer ist ein **Proxy, der auf mehrere Backends verteilt**. Statt
alle Requests an einen Server zu schicken, streut er sie über viele — damit
keiner überlastet wird und der Dienst nicht ausfällt, wenn einer stirbt.

## Warum es wichtig ist

Das ist der Schritt von „ein Server" zu „skalierbares System". Load Balancing +
Health Checks sind das Rückgrat jeder Produktions-Architektur.

## Was du baust

- Nimm deinen Reverse Proxy (Stufe 06) und verteile auf **N Backends**.
- **Verteilungsstrategie:** Round-Robin (reihum), dann Least-Connections.
- **Health Checks:** regelmäßig prüfen, ob ein Backend lebt; tote überspringen.
- **Stretch:** Gewichtung, „sticky sessions", automatisches Wiederaufnehmen
  genesener Backends.

## Kernkonzepte

- **Algorithmen:** Round-Robin, Least-Connections, Weighted, Random
- **Health Checks** (aktiv vs. passiv) und Failover
- **Nebenläufigkeit:** die Backend-Liste sicher aus mehreren Goroutines lesen/
  ändern (Mutex oder Channel) — Bezug zu Stufe 02
- L4 (TCP) vs. L7 (HTTP) Load Balancing

## Literatur

- **Kasun Vithanage — „Let's Create a Simple Load Balancer With Go"** (Blogpost)
  — baut genau das, sehr gut nachvollziehbar.
- **NGINX-Doku — „HTTP Load Balancing"** — die Algorithmen in der Praxis.
- **Google — „Site Reliability Engineering"** (kostenlos online), Kapitel zu
  Load Balancing — das große Bild.
- **Woodbeck, Kap. 10** — nebenläufige Server-Muster.
