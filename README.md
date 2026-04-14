# Berlin-Uhr (Mengenlehreuhr)

## Projektbeschreibung
Dieses Projekt ist eine digitale Nachbildung der Berliner Mengenlehreuhr (auch Berlin-Uhr genannt). Die Anzeige der Uhrzeit erfolgt in einem Stellenwertsystem zur Basis 5, dargestellt durch insgesamt 24 Leuchtsegmente und ein Sekunden-Blinklicht.

Das Ziel dieser Case Study war die Implementierung einer entkoppelten Architektur, bestehend aus einem Backend-Service zur Logikberechnung und einem Frontend-Client zur grafischen Visualisierung.

### Funktionsweise
Die Zeit wird durch Addition der leuchtenden Segmente in vier Zeilen dargestellt:
* **Sekundenanzeige**: Ein rundes Blinklicht (oberhalb der Zeilen), das im Sekundentakt den Zustand wechselt.
* **1. Zeile (Stunden)**: 4 blaue Segmente; jedes Segment steht für 5 volle Stunden.
* **2. Zeile (Stunden)**: 4 cyan Segmente; jedes Segment steht für 1 volle Stunde.
* **3. Zeile (Minuten)**: 11 Segmente; jedes Segment steht für 5 Minuten. Alle 15 Minuten ein blaues Feld zur besseren Übersicht. 
* **4. Zeile (Minuten)**: 4 cyan Segmente; jedes Segment steht für 1 Minute.

## Architektur & Technologiewahl
Entsprechend der Aufgabenstellung wurden zwei Programmiersprachen gewählt, die außerhalb meines bisherigen Repertoires liegen, um die Einarbeitungsfähigkeit in neue Ökosysteme zu demonstrieren.

### Backend: Service "BerlinClock"
* **Sprache**: Go (Golang)
* **Aufgabe**: Berechnung der Zeitintervalle und Bereitstellung der Daten über eine REST-Schnittstelle.
* **Begründung**: Go wurde gewählt, um einen performanten und typsicheren Service zu realisieren. Die Kompilierbarkeit in eine statische Binärdatei gewährleistet ein unkompliziertes Deployment.

### Frontend: Client
* **Sprache**: Svelte
* **Aufgabe**: Konsumierung der API-Daten und reaktive grafische Darstellung der Uhr.
* **Begründung**: Svelte bietet einen compiler-basierten Ansatz für Reaktivität ohne Virtual-DOM-Overhead. Dies ermöglicht eine effiziente, sekündliche Aktualisierung der UI mit minimalem Code-Footprint.



## Qualitätssicherung
* **Unit-Tests**: Die mathematische Umrechnungslogik im Backend ist durch automatisierte Tests in Go abgesichert (`main_test.go`).
* **Separation of Concerns**: Strikte Trennung zwischen Geschäftslogik (Backend) und Präsentationsschicht (Frontend).

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
