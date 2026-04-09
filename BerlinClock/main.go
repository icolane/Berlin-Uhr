package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ClockResponse struct {
	IsLeapSecond bool `json:"isLeapSecond"`
	HoursFive    int  `json:"hoursFive"`
	HoursOne     int  `json:"hoursOne"`
	MinutesFive  int  `json:"minutesFive"`
	MinutesOne   int  `json:"minutesOne"`
}

func calculateBerlinClock(t time.Time) ClockResponse {
	return ClockResponse{
		IsLeapSecond: t.Second()%2 == 0,
		HoursFive:    t.Hour() / 5,
		HoursOne:     t.Hour() % 5,
		MinutesFive:  t.Minute() / 5,
		MinutesOne:   t.Minute() % 5,
	}
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(calculateBerlinClock(time.Now()))
}

func main() {
	http.HandleFunc("/time", clockHandler)
	http.ListenAndServe(":8080", nil)
}
