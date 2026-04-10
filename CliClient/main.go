package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClockResponse struct {
	IsLeapSecond bool `json:"isLeapSecond"`
	HoursFive    int  `json:"hoursFive"`
	HoursOne     int  `json:"hoursOne"`
	MinutesFive  int  `json:"minutesFive"`
	MinutesOne   int  `json:"minutesOne"`
	Seconds      int  `json:"seconds"`
}

// Pollt einmal die Sekunde nach der aktuellen Uhrzeit vom Server und gibt sie auf der Konsole aus.
func main() {
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("Warte auf Server (http://localhost:8080)...")

	for range ticker.C {
		resp, err := http.Get("http://localhost:8080/time")
		if err != nil {
			fmt.Printf("\rFehler: Der Server ist nicht erreichbar. Es wird jede Sekunde erneut versucht eine Verbindung herzustellen.")
			continue
		}

		var clock ClockResponse
		err = json.NewDecoder(resp.Body).Decode(&clock)
		resp.Body.Close()

		if err != nil {
			fmt.Println("Fehler beim Dekodieren.")
			continue
		}

		printClockToConsole(clock)
	}
}

// Ausgabe der Uhrzeit auf der Konsole. mit Hilfestellungen
func printClockToConsole(c ClockResponse) {

	reset := "\033[0m"
	bgBlue := "\033[44m"
	bgCyan := "\033[46m"
	bgOff := "\033[1;100m"
	fgCyan := "\033[36m"
	fgWhite := "\033[97m"

	// Terminal leeren
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%sBERLIN-UHR CONSOLE-CLIENT%s\n", fgWhite, reset)
	fmt.Println("======================================")

	s := "  "
	if c.IsLeapSecond {
		s = fgCyan + "██" + reset
	}
	fmt.Printf("Sekunden:           %s  [ %02d:%02d:%02d ]\n\n", s, (c.HoursFive*5 + c.HoursOne), (c.MinutesFive*5 + c.MinutesOne), c.Seconds)

	printColorRow("5-Stunden:  ", c.HoursFive, 4, bgBlue, bgOff)
	printColorRow("1-Stunde:   ", c.HoursOne, 4, bgBlue, bgOff)

	fmt.Printf("5-Minuten:  ")
	for i := 1; i <= 11; i++ {
		if i <= c.MinutesFive {
			color := bgCyan
			if i%3 == 0 {
				color = bgBlue
			}
			fmt.Printf("%s  %s ", color, reset)
		} else {
			fmt.Printf("%s  %s ", bgOff, reset)
		}
	}
	fmt.Println("\n")

	printColorRow("1-Minute:   ", c.MinutesOne, 4, bgCyan, bgOff)
	fmt.Println("======================================")
	fmt.Printf("%sSoftware by Daniel Layne%s\n", fgCyan, reset)
}

// Hilfsfunktion zur Ausgabe der farbigen Blöcke
func printColorRow(label string, active, total int, activeColor, offColor string) {
	reset := "\033[0m"
	fmt.Print(label)
	for i := 0; i < total; i++ {
		color := offColor
		if i < active {
			color = activeColor
		}
		fmt.Printf("%s      %s  ", color, reset) // Breite Blöcke
	}
	fmt.Println("\n")
}
