# Berlin-Uhr (Mengenlehreuhr) - Fullstack Application

Dieses Projekt ist eine digitale Nachbildung der berühmten Berliner "Mengenlehreuhr", entwickelt von Dieter Binninger (1975). Es kombiniert ein robustes Backend-System mit einem hochmodernen, interaktiven Frontend, um die Zeitvisualisierung im Basis-5-System ästhetisch ansprechend darzustellen.

## 🚀 Projektbeschreibung

Die Berlin-Uhr zeigt die Zeit über ein System von 24 Leuchten an:
- **Sekundenanzeige**: Ein pulsierender Glas-Orb an der Spitze (blinkt im 2-Sekunden-Takt).
- **Stunden (5h)**: Die oberste Zeile zeigt 5-Stunden-Blöcke (Rot).
- **Stunden (1h)**: Die zweite Zeile zeigt einzelne Stunden (Rot).
- **Minuten (5m)**: Die dritte Zeile zeigt 5-Minuten-Blöcke (Gelb, mit roten Markierungen für 15, 30, 45 Min).
- **Minuten (1m)**: Die unterste Zeile zeigt einzelne Minuten (Gelb).

Die aktuelle Zeit ergibt sich aus der Addition der leuchtenden Segmente.

## 🛠 Technologie-Stack & Architektur

Gemäß der Aufgabenstellung wurden zwei für das Team neue Technologien gewählt, um den Fokus auf Lernfähigkeit und saubere Codequalität zu legen:

- **Backend: Go (Golang)**
  - Ein performanter Service, der die Systemzeit in das spezifische Berlin-Uhr-Format umrechnet.
  - Bereitstellung einer REST-API über den Endpunkt `/time`.
- **Frontend: Svelte**
  - Ein reaktives Framework für eine flüssige, komponentenbasierte Darstellung.
  - **Kreatives Design**: Realistisches "Studio-Look"-Interface mit gebürsteten Metalltexturen, Chrom-Halterung und interaktiven Elementen.
  - **Interaktivität**: Ein "atmender" Hilfe-Button blendet bei Bedarf eine digitale Zeitanzeige ein.

### Architektur
Das Projekt folgt einer strikten Trennung zwischen Präsentation (Frontend) und Logik (Backend). Das Backend berechnet die Zustände der Lampen, während das Frontend ausschließlich für die visuelle Shader-Logik und die Benutzerinteraktion zuständig ist.

## 📦 Installation & Start

### Voraussetzungen
- [Go](https://golang.org/dl/) (mind. v1.18)
- [Node.js](https://nodejs.org/) (mind. v16)

### 1. Backend starten
```bash
cd BerlinClock
go run main.go
```
Der Server läuft nun auf `http://localhost:8080`.

### 2. Frontend starten
```bash
cd Client
npm install
npm run dev
```
Die Anwendung ist nun unter `http://localhost:5173` erreichbar.
---
*Entwickelt als Demonstration für saubere Architektur und kreatives UI/UX-Design.*
