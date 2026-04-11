# Berlin-Uhr - Client-Server Applikation

Dieses Projekt ist eine Implementierung der Berliner Uhr (Mengenlehreuhr). Es handelt sich um eine Fullstack-Anwendung bestehend aus einem Go-Backend, einem Svelte-Webfrontend und einem kommandozeilenbasierten Client.

## Screenshots

<p>
  <img src="https://github.com/user-attachments/assets/8d1c0225-1e57-41b0-91f6-aa48b2750538" width="200" />
</p>
<p>
<img src="https://github.com/user-attachments/assets/d039c662-32d4-4f55-83d3-a7f804143ae4" />
</p>
## Voraussetzungen

Für den Betrieb und die Kompilierung der Anwendung werden folgende Komponenten benötigt:

*   **Go**: Version 1.20 oder höher (für Backend und CLI-Client).
*   **Node.js**: Version 18.x oder höher.
*   **npm**: Version 9.x oder höher (für das Web-Frontend).
*   **PowerShell**: (Optional, für die Nutzung des Windows-Build-Skripts).

## System-Architektur

Das System ist in drei Komponenten unterteilt:

1.  **Backend (Go-Server)**:
    *   Stellt eine REST-API bereit, die die berechneten Uhrzeit-Daten als JSON liefert.
    *   Berechnet die Zeitsegmente (Stunden, Minuten, Sekunden) gemäß der Logik der Berliner Uhr.
    *   Beinhaltet Unit-Tests für die Kernlogik.
2.  **Web-Frontend (Svelte)**:
    *   Visualisiert die Uhrzeit in einer grafischen Oberfläche.
    *   Abfrage der Daten erfolgt über die REST-API des Backends.
    *   Bietet eine optionale numerische Anzeige der Zeitwerte.
3.  **CLI-Client (Go)**:
    *   Eine Konsolen-Anwendung, die die Uhrzeit in das Terminal rendert.
    *   Nutzt ANSI-Escape-Sequenzen für die farbige Darstellung der Zeitblöcke.

## Installation & Betrieb

### 1. Backend & Server
Navigieren Sie in das Verzeichnis `BerlinClock` und starten Sie den Server:
```bash
go run main.go
```
*Der Server ist standardmäßig unter `http://localhost:8080/time` erreichbar.*

### 2. Web-Frontend
Navigieren Sie in das Verzeichnis `Client`, installieren Sie die Abhängigkeiten und starten Sie den Entwicklungs-Server:
```bash
npm install
npm run dev
```

### 3. CLI-Client
Navigieren Sie in das Verzeichnis `CliClient` und starten Sie die Anwendung:
```bash
go run main.go
```

## Kompiliervorgang (Windows)

Zur Erstellung von ausführbaren Dateien (.exe) kann das bereitgestellte PowerShell-Skript genutzt werden:
```powershell
.\build.ps1
```
Die Dateien werden im Ordner `build` abgelegt.

## Unit-Tests

Ausführung der Tests für die Zeitberechnung:
```bash
cd BerlinClock
go test -v ./...
```

## Lizensierung und Quellen

Dieses Projekt nutzt Logiken der Berlin-Uhr von icolane.
Autor: icolane
Quelle: https://github.com/icolane/Berlin-Uhr/
