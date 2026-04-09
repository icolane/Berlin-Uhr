package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// ClockResponse definiert die Struktur der Antwort, die die Werte für die Berliner Uhr enthält
type ClockResponse struct {
	IsLeapSecond bool `json:"isLeapSecond"`
	HoursFive    int  `json:"hoursFive"`
	HoursOne     int  `json:"hoursOne"`
	MinutesFive  int  `json:"minutesFive"`
	MinutesOne   int  `json:"minutesOne"`
}

// Berechnet die Werte für die Berliner Uhr basierend auf der gegebenen Zeit
func calculateBerlinClock(t time.Time) ClockResponse {
	return ClockResponse{
		IsLeapSecond: t.Second()%2 == 0,
		HoursFive:    t.Hour() / 5,
		HoursOne:     t.Hour() % 5,
		MinutesFive:  t.Minute() / 5,
		MinutesOne:   t.Minute() % 5,
	}
}

// HTTP-Handler, der die aktuelle Zeit in das Format der Berliner Uhr umwandelt und als JSON zurückgibt
func clockHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(calculateBerlinClock(time.Now()))
}

// Startet den HTTP-Server und bindet den Handler an die Route "/time"
func main() {
	http.HandleFunc("/time", clockHandler)
	http.ListenAndServe(":8080", nil)
}
