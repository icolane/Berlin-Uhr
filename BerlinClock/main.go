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

func getBerlinClock(t time.Time) ClockResponse {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	isLeapSecond := second%2 == 1
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
	}
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(getBerlinClock(time.Now()))
}

func main() {
	http.HandleFunc("/time", clockHandler)
	http.ListenAndServe(":8080", nil)
}
