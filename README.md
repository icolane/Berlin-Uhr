# Professional Berlin Clock - Multi-Interface System

Dieses Projekt ist eine hochmoderne Implementierung der klassischen Berliner Uhr (Mengenlehreuhr), entwickelt als Fullstack-Anwendung mit Go (Backend) und Svelte (Frontend). Es demonstriert Architektur-Kompetenz durch eine Client-Server-Trennung und bietet sowohl ein High-End Web-Interface als auch einen technoiden CLI-Monitor.

## 🚀 System-Architektur

Das System besteht aus drei entkoppelten Komponenten:

1.  **Backend (Go-Server)**:
    *   REST-API zur Bereitstellung der berechneten Uhrzeit-Daten.
    *   Echtzeit-Konsolenvisualisierung des Systemstatus via Goroutines.
    *   Abgedeckt durch automatisierte Unit-Tests (`main_test.go`).
2.  **Web-Frontend (Svelte)**:
    *   "Cyber-Tech" Design-System im Stealth-Look (Gunmetal / Indigo / Ice-Cyan).
    *   Interaktive Wertigkeits-Explikation: Auf Knopfdruck werden die mathematischen Werte direkt in den Segmenten eingeblendet.
    *   High-End CSS-Animationen für organische UI-Elemente.
3.  **CLI-Client (Go)**:
    *   Eigenständige Konsolen-Anwendung, die das Interface direkt in das Terminal rendert (ANSI-Block-Rendering).
    *   Demonstration von Client-Server-Interaktion via HTTP/JSON.

## 🎨 Design-Konzept

Im Gegensatz zum klassischen Rot-Gelben Design nutzt diese Version eine moderne Farbpalette aus **Indigo-Blau** (Stunden/Markierungen) und **Eis-Cyan** (Minuten). Dies unterstreicht den technischen Charakter der Lösung und sorgt für eine exzellente Lesbarkeit im Dunkelmodus. Das Gehäuse ist minimalistisch und scharfkantig gestaltet, um industrielle Präzision zu vermitteln.

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

## 🧪 Qualitätssicherung

Das Projekt beinhaltet Unit-Tests für die Kernlogik der Zeitberechnung.
Ausführung der Tests:
```bash
cd BerlinClock
go test -v ./...
```

---
*Entwickelt als technischer Beitrag für den Einstellungsprozess bei JTL / Greyhound.*
