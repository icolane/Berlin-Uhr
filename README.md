# Berlin-Uhr - Coding Challenge / Technical Demo

Dieses Projekt ist eine Implementierung der Berliner Uhr (Mengenlehreuhr) im Rahmen einer technischen Case Study (Job-Bewerbung). Es dient zur Demonstration von Fullstack-Kompetenzen, sauberer Software-Architektur und performanter Implementierung in Go und Svelte.

## Projekt-Kontext & Zielsetzung
Dieses Repository wurde als Teil einer Bewerbung erstellt, um Fähigkeiten in den folgenden Bereichen zu demonstrieren:
- **Clean Code & Unit Testing** (Backend-Validierung).
- **Backend-Entwicklung mit Go** (REST-API, Performance-Optimierung).
- **Frontend-Entwicklung mit Svelte** (Modernes UI/UX, CSS-Animationen).
- **Systemprogrammierung** (CLI-Client, ANSI-Terminal-Handling).
- **Tooling** (CI/CD-Ansätze, Build-Skripte).

## Kern-Features & Architektur

*   **Zentrales Backend**: Ein Go-Server, der die Berlin-Uhr-Logik präzise abbildet und per REST-API bereitstellt.
*   **Präzisions-CLI (Go)**: Ein flickerfreier Terminal-Client mit ANSI-Farben, optimiertem Update-Zyklus und Echtzeit-Logik.
*   **Premium Web-Frontend (Svelte)**: Responsive UI mit minimalistischem Glassmorphism-Design, CSS-Gradients und interaktivem "Help-Modus".
*   **Performance & Stabilität**: Konsequente Verwendung von optimierten Render-Zyklen (CLI & Web) sowie eine saubere Trennung der Concerns.

## Voraussetzungen

*   **Go**: Version 1.20 oder höher.
*   **Node.js**: Version 18.x oder höher (für das Web-Frontend).
*   **npm**: Version 9.x oder höher.

## System-Architektur

1.  **Backend (Go-Server)**:
    *   Liefert Zeitdaten als JSON über eine REST-API.
    *   Präzise Berechnung der Stunden- (5h/1h), Minuten- (5m/1m) und Sekundensegmente.
    *   Unit-Tests für die Validierung der Zeitlogik.
2.  **Web-Frontend (Svelte)**:
    *   **Minimalist Design**: Standardmäßig fokussiert auf die reine Licht-Visualisierung.
    *   **Help-Modus**: Blendet per Klick die Sekunden-Segmente und die digitale Uhrzeit ein.
    *   **Dynamic Colors**: Amber-Blinker mit integrierter Sekunden-Zahl.
3.  **CLI-Client (Go)**:
    *   **Cyber-Monitor Look**: Hochwertige ASCII/ANSI-Visualisierung.
    *   **Clearscreen-Logik**: Verhindert Flackern und zeigt Status-Informationen übersichtlich an.

## Installation & Betrieb

### 1. Backend
```bash
cd BerlinClock
go run main.go
```
*API: `http://localhost:8080/time`*

### 2. Web-Frontend
```bash
cd Client
npm install
npm run dev
```

### 3. CLI-Client
```bash
cd CliClient
go run main.go
```

## Kompiliervorgang (Windows)
Nutzen Sie das Build-Skript für Standalone-Executables:
```powershell
.\build.ps1
```

## Unit-Tests
```bash
cd BerlinClock
go test -v ./...
```

## Lizenz & Hinweis
Dieses Projekt wurde als technische Demo erstellt und steht nicht zum Verkauf. Es basiert auf der Logik der Berlin-Uhr (icolane).
Autor: icolane | Quelle: https://github.com/icolane/Berlin-Uhr/

## :camera: Screenshots

<p align="center">
   <img src="https://github.com/user-attachments/assets/6f80669a-62d2-4842-a63e-35f5e12ca315" width="600" />
</p>
