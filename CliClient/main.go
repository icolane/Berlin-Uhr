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
			printErrorToConsole(err)
			continue
		}

		var clock ClockResponse
		err = json.NewDecoder(resp.Body).Decode(&clock)
		resp.Body.Close()

		if err != nil {
			printErrorToConsole(err)
			continue
		}

		printClockToConsole(clock)
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
	fgGreen := "\033[32m"

	// Window Title setzen & Cursor an den Anfang
	fmt.Print("\033]0;BERLIN-UHR | CYBER-MONITOR v2.0\007")
	fmt.Print("\033[H")

	s := fgGray + "░░" + reset
	if c.IsLeapSecond {
		s = fgCyan + "██" + reset
	}
	fmt.Printf(" [ SYSTEM STATUS ]:  %sONLINE%s     [ TIME ]: %s %02d:%02d:%02d %s\n", fgGreen, reset, fgWhite, (c.HoursFive*5 + c.HoursOne), (c.MinutesFive*5 + c.MinutesOne), c.Seconds, reset)
	fmt.Printf(" [ BLINKER      ]:  %s\n\n", s)

	printBlockRow(" 5-STUNDEN      ", c.HoursFive, 4, fgBlue, fgGray, "██████")
	printBlockRow(" 1-STUNDE       ", c.HoursOne, 4, fgBlue, fgGray, "██████")

	fmt.Printf(" 5-MINUTEN      ")
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
	fmt.Println("\n")

	printBlockRow(" 1-MINUTE       ", c.MinutesOne, 4, fgCyan, fgGray, "██████")

	fmt.Printf("%s----------------------------------------------------%s\n", fgGray, reset)

	fmt.Printf(" 5-SEKUNDEN     ")
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
	fmt.Println("\n")

	printBlockRow(" 1-SEKUNDE      ", c.SecondsOne, 4, fgMagenta, fgGray, "██████")

	fmt.Printf("%s====================================================%s\n", fgCyan, reset)
	fmt.Printf(" %sLEGENDE:%s  %s■ STUNDEN%s  %s■ MINUTEN%s  %s■ SEKUNDEN%s\n", fgWhite, reset, fgBlue, reset, fgCyan, reset, fgMagenta, reset)
	fmt.Printf(" %s(c) 2026 Developed by Daniel Layne%s\n", fgGray, reset)
}

// Zeigt einen  Fehlerbildschirm an, wenn der Server nicht erreichbar ist
func printErrorToConsole(err error) {
	reset := "\033[0m"
	fgRed := "\033[31m"
	fgWhite := "\033[97m"

	// Cursor an den Anfang
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033]0;BERLIN-UHR | CONNECTION ERROR\007")

	fmt.Printf("%s┌──────────────────────────────────────────────────┐%s\n", fgRed, reset)
	fmt.Printf("%s│                                                  │%s\n", fgRed, reset)
	fmt.Printf("%s│   [!] COMMUNICATION ERROR: BACKEND OFFLINE [!]   │%s\n", fgRed, reset)
	fmt.Printf("%s│                                                  │%s\n", fgRed, reset)
	fmt.Printf("%s└──────────────────────────────────────────────────┘%s\n\n", fgRed, reset)

	fmt.Printf("%sSTATUS:%s   %sSUCHE SERVER... (http://localhost:8080)%s\n", fgWhite, reset, fgRed, reset)
	fmt.Printf("%sFEHLER:%s   %v\n\n", fgWhite, reset, err)
	fmt.Printf("%sDas System versucht jede Sekunde die Verbindung neu aufzubauen.%s\n", fgWhite, reset)
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
	fmt.Println("\n")
}
