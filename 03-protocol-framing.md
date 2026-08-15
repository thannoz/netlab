# 03 — Eigenes Nachrichten-Protokoll (Framing)

## Worum es geht

TCP ist ein **Bytestrom** — es gibt keine „Nachrichten". Wenn du „Hallo" und
„Welt" sendest, kann der Empfänger „HalloWelt", „Hal" + „loWelt" oder alles auf
einmal lesen. Du musst selbst festlegen, **wo eine Nachricht endet**. Das nennt
man *Framing*.

## Warum es wichtig ist

Das ist der „Aha"-Moment, an dem du verstehst, warum es Protokolle wie HTTP
überhaupt gibt. Jedes Protokoll löst genau dieses Problem. Danach ist HTTP
(Stufe 04) plötzlich entmystifiziert.

## Was du baust

Ein kleines eigenes Protokoll, z. B. **length-prefixed**: vor jede Nachricht
schreibst du ihre Länge (z. B. 4 Bytes), dann liest der Empfänger erst die
Länge, dann genau so viele Bytes.

```
[Länge: 4 Bytes][Nutzdaten: N Bytes][Länge][Nutzdaten]...
```

- Encoder: Nachricht → Bytes auf dem Draht.
- Decoder: Bytes vom Draht → Nachricht.
- **Stretch:** ein Typ-Feld (verschiedene Nachrichtentypen), Versionierung.

## Kernkonzepte

- **Bytestrom vs. Nachrichtengrenzen** — das Kernproblem
- **Framing-Strategien:** length-prefix, Delimiter (z. B. `\n`), fixe Größe
- `encoding/binary` (Big-/Little-Endian), `io.ReadFull`
- `bufio.Reader`/`bufio.Scanner` und warum Puffer helfen
- Ein Vorgeschmack auf **Serialisierung** (JSON, Protobuf) für später

## Literatur

- **Woodbeck — „Network Programming with Go", Kap. 4** (Sending TCP Data) —
  genau dieses Thema, inkl. length-prefix Framing in Go.
- **W. Richard Stevens — „TCP/IP Illustrated, Vol. 1"** — was TCP wirklich tut
  (Segmente, warum es keine Nachrichtengrenzen gibt).
- **Go-Doku:** `encoding/binary`, `bufio`.
- Artikel-Suche: „message framing tcp" — viele gute Blogposts erklären die
  Strategien im Vergleich.
