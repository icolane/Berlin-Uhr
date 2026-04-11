# Berlin-Uhr - Cyber-Monitor v2.0

Dieses Projekt ist eine hochmoderne Implementierung der Berliner Uhr (Mengenlehreuhr). Es handelt sich um eine Fullstack-Anwendung bestehend aus einem performanten Go-Backend, einem modernen Svelte-Webfrontend im "Cyber-Corporate"-Design und einem professionellen CLI-Client.

## Features & Highlights

*   **[NEW] Cyber-Monitor CLI**: Ein flickerfreier Terminal-Client mit ANSI-Farben, Echtzeit-Update-Logik und integrierter Sekundenanzeige im Blinker.
*   **[NEW] Interaktive Web-UI**: Ein minimalistisches, premium Design mit einem "Help-Modus" (Fragezeichen-Icon), der detaillierte Informationen und digitale Zeitwerte per Klick einblendet.
*   **Flicker-Free Rendering**: Sowohl CLI als auch Web nutzen optimierte Update-Zyklen für eine flüssige Visualisierung.
*   **RESTful Core**: Ein robustes Go-Backend berechnet die Zeitsegmente präzise nach der Mengenlehreuhr-Logik.

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

## Lizensierung und Quellen
Dieses Projekt nutzt Logiken der Berlin-Uhr von icolane.
Autor: icolane | Quelle: https://github.com/icolane/Berlin-Uhr/

## :camera: Screenshots
<p>
  <img src="https://github.com/user-attachments/assets/8d1c0225-1e57-41b0-91f6-aa48b2750538" width="200" />
  <img src="https://github.com/user-attachments/assets/d039c662-32d4-4f55-83d3-a7f804143ae4" width="200" />
  <img src="https://github.com/user-attachments/assets/7160dfe6-aeb4-4adf-81e2-40bee12ca2c2" width="250" />
</p>

