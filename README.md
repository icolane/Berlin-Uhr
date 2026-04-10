# Berlin-Uhr - Client Server Applikation

Dieses Projekt ist eine Implementierung der klassischen Berliner Uhr (<a href="https://de.wikipedia.org/wiki/Berlin-Uhr">Mengenlehreuhr</a>), entwickelt als Fullstack-Anwendung mit Go (Backend) und Svelte (Frontend). Es demonstriert Architektur-Kompetenz durch eine Client-Server-Trennung und bietet sowohl ein Web-Interface als auch einen CLI-Monitor.

## 🚀 System-Architektur

Das System besteht aus drei entkoppelten Komponenten:

1.  **Backend (Go-Server)**:
    *   REST-API zur Bereitstellung der berechneten Uhrzeit-Daten.
    *   Echtzeit-Konsolenvisualisierung des Systemstatus via Goroutines.
    *   Abgedeckt durch automatisierte Unit-Tests (`main_test.go`).
2.  **Web-Frontend (Svelte)**:
    *   Design-System im Stealth-Look (Gunmetal / Indigo / Ice-Cyan).
    *   Interaktive Wertigkeits-Explikation: Auf Knopfdruck werden die mathematischen Werte direkt in den Segmenten eingeblendet.
    *   CSS-Animationen für organische UI-Elemente.
3.  **CLI-Client (Go)**:
    *   Eigenständige Konsolen-Anwendung, die das Interface direkt in das Terminal rendert (ANSI-Block-Rendering).
    *   Demonstration von Client-Server-Interaktion via HTTP/JSON.

## 🎨 Design-Konzept

Im Gegensatz zum klassischen Rot-Gelben Design nutzt diese Version eine moderne Farbpalette aus **Indigo-Blau** (Stunden/Markierungen) und **Eis-Cyan** (Minuten). Dies unterstreicht den technischen Charakter der Lösung und sorgt für eine exzellente Lesbarkeit im Dunkelmodus. Das Gehäuse ist minimalistisch und scharfkantig gestaltet, um industrielle Präzision zu vermitteln.
Zusätzlich werden die Sekunde auch angezeigt.

## 🛠️ Installation & Start

### 1. Backend & Server
Navigiere in das `BerlinClock` Verzeichnis und starte den Server:
```bash
cd BerlinClock
go run main.go
```
*Der Server ist nun unter `http://localhost:8080/time` erreichbar.*

### 2. Web-Frontend (UI)
Navigiere in das `Client` Verzeichnis, installiere die Abhängigkeiten und starte den Dev-Server:
```bash
cd Client
npm install
npm run dev
```

### 3. Console-Client (CLI)
Navigiere in das `CliClient` Verzeichnis und starte den Monitor:
```bash
cd CliClient
go run main.go
```

## 🧪 UNIT-Tests

Das Projekt beinhaltet Unit-Tests für die Kernlogik der Zeitberechnung.
Ausführung der Tests:
```bash
cd BerlinClock
go test -v ./...
```

## :camera: Screenshots
<p align="center"><img src="https://github.com/user-attachments/assets/69cb3431-8be4-42f9-bfc5-7635d16f37d0" alt="Screenshot Berlin-Uhr" width="600"></p>
<p align="center"><img src="https://github.com/user-attachments/assets/f452bdb6-10b2-4e78-9d2d-0957bcf2bb39" alt="Screenshot Berlin-Uhr CLI" width="600"></p>


## Lizensierung

Credit / Quellenangabe
<p>Dieses Projekt nutzt Code-Teile oder Logiken der Berlin-Uhr von icolane.
Autor: icolane
Quelle: https://github.com/icolane/Berlin-Uhr/
Hinweis: Da das Original-Repository keine explizite Open-Source-Lizenz beinhaltet, liegen alle Urheberrechte am ursprünglichen Code beim Ersteller.</p>
