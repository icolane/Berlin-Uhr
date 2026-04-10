# BerlinClock - Go & Svelte Implementation

Diese Repository enthält eine Fullstack-Implementierung der Berliner Uhr (Mengenlehreuhr). Die Lösung ist in ein Go-Backend und ein Svelte-Frontend unterteilt, was eine saubere Trennung von Geschäftslogik und Präsentationsschicht ermöglicht.

## Architektur & Technologien

Im Rahmen der Aufgabenstellung wurden **Go** und **Svelte** gewählt, um eine moderne und performante Systemarchitektur abzubilden. Beide Technologien wurden für dieses Projekt neu erschlossen.

### Backend (Go)
Das Backend fungiert als zustandsloser Service. Es berechnet die Logik der Berlin-Uhr auf Basis der aktuellen Systemzeit und stellt diese über eine REST-API (`/time`) zur Verfügung.
- **Unit-Testing**: Die Kernlogik der Zeitumrechnung ist über Go-Tests validiert.
- **JSON API**: Standardisierte Ausgabe für die Konsumierung durch beliebige Clients.

### Frontend (Svelte)
Das Frontend visualisiert die Daten des Backends in einem hochwertigen, interaktiven Interface.
- **UI/UX Design**: Umsetzung in einem fotorealistischen Studio-Setting mit CSS-basierten Shadern für Metall- und Glaseffekte.
- **Reaktivität**: Automatische Synchronisation mit dem Backend-Service.
- **Interaktivität**: Optional einblendbare digitale Hilfe für bessere Lesbarkeit.

## Installation & Ausführung

### Backend (Port 8080)
```bash
cd BerlinClock
go test ./...    # Ausführen der Unit-Tests
go run main.go   # Starten des Servers
```

### Frontend (Port 5173)
```bash
cd Client
npm install      # Abhängigkeiten installieren
npm run dev      # Entwicklungs-Server starten
```

## Projektstruktur
- `/BerlinClock`: Go-Quellcode inkl. Unit-Tests.
- `/Client`: Svelte-Frontend (Vite-basiert).
- `/Client/src/App.css`: Zentrales Stylesheet für das Studio-Design.

---
*Anmerkung: Das Projekt legt besonderen Wert auf Code-Klarheit und eine hochwertige visuelle Aufbereitung, um die Brücke zwischen klassischer Mengenlehre und modernen Web-Technologien zu schlagen.*
