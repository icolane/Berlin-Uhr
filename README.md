# Berlin-Uhr (Mengenlehreuhr)
# Projektbeschreibung
Dieses Projekt ist eine digitale Nachbildung der Berliner Mengenlehreuhr, einer öffentlichen Uhr, die 1975 von Dieter Binninger entwickelt wurde. Die Anzeige der Uhrzeit erfolgt über ein Stellenwertsystem zur Basis 5, dargestellt durch insgesamt 24 Leuchtsegmente in vier horizontalen Zeilen sowie einem Sekunden-Blinklicht

## Funktionsweise der Anzeige
Die Zeit wird durch Addition der leuchtenden Segmente berechnet:
*   **Sekundenanzeige**: Ein rundes Blinklicht über den Zeilen, das im Sekundentakt blinkt.
*   **Stunden(Zeile 1&2)**: Die erste Zeile zeigt Blöcke von je 5 Stunden, die zweite Zeile die einzelnen Stunden.
*   **Minuten(Zeile 3&4)**: Die dritte Zeile zeigt Blöcke von je 5 Minuten (gelbe Segmente, wobei die Markierungen für 15, 30 und 45 Minuten zur besseren Lesbarkeit rot gefärbt sind). Die vierte Zeile zeigt einzelne Minuten in Gelb.

## Kern-Features & Architektur

bereitstellt.
*   **Zentrales Backend**: Service "BerlinClock" (Backend): Ein in Go implementierter Service, der die aktuelle Systemzeit berechnet und im geforderten Format für die Mengenlehreuhr bereitstellt
*   **Client-Frontend (Svelte)**: : Eine Svelte-Webapplikation, welche die Daten des Services konsumiert und die Uhrzeit grafisch analog zum Original visualisiert.
*   **CLI (Go)**: Zusätzlich eine CLI in GO.

## Voraussetzungen

*   **Go**: Version 1.20 oder höher.
*   **Node.js**: Version 18.x oder höher (für das Web-Frontend).
*   **npm**: Version 9.x oder höher.

## System-Architektur

1.  **Backend (Go-Server)**:
    *   Liefert Zeitdaten als JSON über eine REST-API.
    *   Berechnung der Stunden- (5h/1h), Minuten- (5m/1m) und Sekundensegmente.
    *   Unit-Tests für die Validierung der Zeitlogik.
2.  **Web-Frontend (Svelte)**:
    *   *Design**: Standardmäßig fokussiert auf die reine Licht-Visualisierung.
    *   **Help-Modus**: Blendet per Klick die Sekunden-Segmente und die digitale Uhrzeit ein.
    
3.  **CLI-Client (Go)**:
    *   **Design**: Hochwertige ASCII/ANSI-Visualisierung.
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
<p align="center">
   <img src="https://github.com/user-attachments/assets/20890c76-5a38-4a44-8c11-ca871d737ef8" width="600" />
</p>
