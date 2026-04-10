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
	SecondsFive  int  `json:"secondsFive"`
	SecondsOne   int  `json:"secondsOne"`
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

		//printClockToConsole(clock)
	}
}

// Ausgabe der Uhrzeit auf der Konsole. mit Hilfestellungen
func printClockToConsole(c ClockResponse) {

	reset := "\033[0m"
	fgBlue := "\033[34m"
	fgCyan := "\033[36m"
	fgMagenta := "\033[35m"
	fgWhite := "\033[97m"
	fgGray := "\033[90m"

	// Terminal leeren
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%sBERLIN-UHR CONSOLE-CLIENT%s\n", fgWhite, reset)
	fmt.Println("======================================")

	s := fgGray + "░░" + reset
	if c.IsLeapSecond {
		s = fgCyan + "██" + reset
	}
	fmt.Printf("Blinker:            %s  [ %02d:%02d:%02d ]\n\n", s, (c.HoursFive*5 + c.HoursOne), (c.MinutesFive*5 + c.MinutesOne), c.Seconds)

	printBlockRow("5-Stunden:  ", c.HoursFive, 4, fgBlue, fgGray, "██████")
	printBlockRow("1-Stunde:   ", c.HoursOne, 4, fgBlue, fgGray, "██████")

	fmt.Printf("5-Minuten:  ")
	for i := 1; i <= 11; i++ {
		char := "██"
		color := fgGray
		if i <= c.MinutesFive {
			color = fgCyan
			if i%3 == 0 {
				color = fgBlue
			}
		} else {
			char = "░░"
		}
		fmt.Printf("%s%s%s ", color, char, reset)
	}

	printBlockRow("1-Minute:   ", c.MinutesOne, 4, fgCyan, fgGray, "██████")

	fmt.Println("--------------------------------------")

	fmt.Printf("5-Sekunden: ")
	for i := 1; i <= 11; i++ {
		char := "██"
		color := fgGray
		if i <= c.SecondsFive {
			color = fgMagenta
		} else {
			char = "░░"
		}
		fmt.Printf("%s%s%s ", color, char, reset)
	}

	printBlockRow("1-Sekunde:  ", c.SecondsOne, 4, fgMagenta, fgGray, "██████")

	fmt.Println("======================================")
	fmt.Printf("%sSoftware by Daniel Layne%s\n", fgCyan, reset)
}

// Hilfsfunktion zur Ausgabe der farbigen Blöcke
func printBlockRow(label string, active, total int, activeColor, offColor string, blockChar string) {
	reset := "\033[0m"
	fmt.Print(label)
	for i := 0; i < total; i++ {
		color := offColor
		char := "░░░░░░"
		if i < active {
			color = activeColor
			char = blockChar
		}
		fmt.Printf("%s%s%s  ", color, char, reset)
	}
}
