# 08 — API Gateway

## Worum es geht

Ein API Gateway ist die **eine Eingangstür** vor vielen Diensten. Es routet
Requests zum richtigen Backend und erledigt Querschnittsaufgaben zentral:
Authentifizierung, Rate-Limiting, Logging, CORS. Ein Proxy (Stufe 06) mit Gehirn.

**Beantwortet deine Frage „was ist ein API Gateway?".**

## Warum es wichtig ist

In Microservice-Architekturen ist das Gateway der zentrale Kontrollpunkt. Es
bündelt alle deine bisherigen Stufen (Proxy, Routing, Middleware, Nebenläufigkeit)
zu einem produktionsnahen System.

## Was du baust

- **Routing:** `/users/*` → User-Service, `/orders/*` → Order-Service.
- **Middleware-Kette:** Auth → Rate-Limit → Logging → Proxy (Stufe 06/07).
- **Auth:** z. B. API-Keys oder JWT prüfen, bevor der Request durchgeht.
- **Rate-Limiting:** pro Client maximal X Requests/Sekunde (Token-Bucket).
- **Stretch:** Config-Datei für Routen, Metriken, Request-Aggregation.

## Kernkonzepte

- **Middleware-Muster in Go** (`func(http.Handler) http.Handler`) — Ketten von
  Handlern
- **Routing** (Pfad-Matching, Präfixe, Parameter)
- **Rate-Limiting-Algorithmen:** Token-Bucket, Leaky-Bucket (`golang.org/x/time/rate`)
- **Auth-Grundlagen:** API-Keys, JWT, warum das Gateway der richtige Ort dafür ist
- Cross-Cutting Concerns: einmal zentral statt in jedem Service

## Literatur

- **Sam Newman — „Building Microservices"** — das API-Gateway-Muster und wann man
  es (nicht) einsetzt.
- **Kong / NGINX API-Gateway-Doku** — was echte Gateways können.
- **Go Blog / Woodbeck, Kap. 9** — HTTP-Middleware sauber komponieren.
- **`golang.org/x/time/rate`** — fertiger Rate-Limiter; lies den Quellcode für
  das Token-Bucket-Verständnis, nachdem du deinen eigenen gebaut hast.
