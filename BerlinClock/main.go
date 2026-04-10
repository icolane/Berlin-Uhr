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

// getBerlinClock berechnet die Werte der Berliner Uhr basierend auf der gegebenen Zeit
func getBerlinClock(t time.Time) ClockResponse {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	isLeapSecond := second%2 == 0
	hoursFive := hour / 5
	hoursOne := hour % 5
	minutesFive := minute / 5
	minutesOne := minute % 5

	return ClockResponse{
		IsLeapSecond: isLeapSecond,
		HoursFive:    hoursFive,
		HoursOne:     hoursOne,
		MinutesFive:  minutesFive,
		MinutesOne:   minutesOne,
		Seconds:      second,
	}
}

// clockHandler nimmt die aktuelle Zeit, berechnet die Berliner Uhr und gibt die Werte als JSON zurück
func clockHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(getBerlinClock(time.Now()))
}

func main() {
	// Goroutine für die Konsolenausgabe (Visualisierung)
	go func() {
		for {
			now := time.Now()
			clock := getBerlinClock(now)
			printClockToConsole(now, clock)
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/time", clockHandler)
	fmt.Println("Berlin-Uhr Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func printClockToConsole(t time.Time, c ClockResponse) {
	// Terminal leeren und Cursor nach oben
	fmt.Print("\033[H\033[2J")
	fmt.Printf("BERLIN-UHR BACKEND STATUS - %s\n", t.Format("15:04:05"))
	fmt.Println("======================================")
	
	s := " "
	if c.IsLeapSecond { s = "*" }
	fmt.Printf("Sekunden-Blinker:  (%s)\n\n", s)
	
	printRow("5-Stunden-Reihe:   ", c.HoursFive, 4, "B") // B für Blue
	printRow("1-Stunden-Reihe:   ", c.HoursOne, 4, "B")
	
	// 5-Minuten-Reihe mit Spezialfarben (B=Blue für 15/30/45)
	fmt.Printf("5-Minuten-Reihe:   ")
	for i := 1; i <= 11; i++ {
		if i <= c.MinutesFive {
			if i%3 == 0 { fmt.Print("[B]") } else { fmt.Print("[C]") } // C für Cyan
		} else {
			fmt.Print("[ ]")
		}
	}
	fmt.Println()
	
	printRow("1-Minuten-Reihe:   ", c.MinutesOne, 4, "C")
	fmt.Println("======================================")
	fmt.Println("API-Endpunkt erreichbar unter /time")
}

func printRow(label string, active, total int, char string) {
	fmt.Print(label)
	for i := 0; i < total; i++ {
		if i < active {
			fmt.Printf("[%s]", char)
		} else {
			fmt.Print("[ ]")
		}
	}
	fmt.Println()
}
