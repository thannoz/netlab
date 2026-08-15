# 05 — Statischer File-Server

## Worum es geht

Ein HTTP-Server (Stufe 04), der **Dateien von der Festplatte ausliefert**:
HTML, Bilder, was auch immer. Request kommt für `/bild.png` → Server liest die
Datei und schickt sie als Response.

## Warum es wichtig ist

Hier verbindest du Networking mit deinem Interesse an **Dateien & I/O**. Und du
lernst ein zentrales Go-Konzept: **Streaming** — große Dateien ausliefern, ohne
sie komplett in den Speicher zu laden.

## Was du baust

- Pfad aus dem Request → Datei auf der Platte finden.
- Datei **streamen** (`io.Copy`) statt komplett einzulesen.
- Korrekten `Content-Type` setzen (nach Dateiendung).
- Fehlerfälle: Datei nicht da (404), Zugriff außerhalb des Ordners verhindern
  (**Path Traversal** — Sicherheit!).
- **Stretch:** Range-Requests (Teildownloads), Caching-Header, Verzeichnislisting.

## Kernkonzepte

- `os.Open`, `io.Reader`/`io.Writer`, **`io.Copy`** (das Streaming-Herzstück)
- Warum man große Dateien *nicht* mit `os.ReadFile` in den RAM lädt
- MIME-Types / `Content-Type`
- **Path-Traversal-Angriff** (`../../etc/passwd`) und wie man ihn abwehrt
- `os.Stat` für Dateigröße/Existenz

## Literatur

- **The Go Programming Language, Kap. 1.7 & Kap. 7** — `io.Reader`/`io.Writer`,
  das wichtigste Interface-Paar in Go.
- **Go-Doku:** `os`, `io`, `io/fs`, `mime`.
- **Woodbeck, Kap. 9** — Dateien über HTTP ausliefern.
- **OWASP — „Path Traversal"** — die Sicherheitslücke verstehen und vermeiden.
