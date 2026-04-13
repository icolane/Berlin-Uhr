package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClockResponse struct {
	IsLeapSecond bool `json:"isLeapSecond"` // Schaltsekunde 60 / 0...
	HoursFive    int  `json:"hoursFive"`
	HoursOne     int  `json:"hoursOne"`
	MinutesFive  int  `json:"minutesFive"`
	MinutesOne   int  `json:"minutesOne"`
	SecondsFive  int  `json:"secondsFive"`
	SecondsOne   int  `json:"secondsOne"`
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
	secondsFive := second / 5
	secondsOne := second % 5

	return ClockResponse{
		IsLeapSecond: isLeapSecond,
		HoursFive:    hoursFive,
		HoursOne:     hoursOne,
		MinutesFive:  minutesFive,
		MinutesOne:   minutesOne,
		SecondsFive:  secondsFive,
		SecondsOne:   secondsOne,
		Seconds:      second,
	}
}

// clockHandler nimmt die aktuelle Zeit, berechnet die Berliner Uhr und gibt die Werte als JSON zurück
func clockHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(getBerlinClock(time.Now()))
}

func main() {

	http.HandleFunc("/time", clockHandler)
	fmt.Println("Berlin-Uhr Server läuft auf http://localhost:8080")
	fmt.Println("Verbinden mit dem Webclient unter http://localhost:5173")
	http.ListenAndServe(":8080", nil)
}
