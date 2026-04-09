package main

import (
	"testing"
	"time"
)

// BerlinClock represents the time displayed on a Berlin Clock
type BerlinClock struct {
	HoursFive    int
	HoursOne     int
	MinutesFive  int
	MinutesOne   int
	IsLeapSecond bool
}

// calculateBerlinClock calculates the Berlin Clock time representation
func calculateBerlinClock(t time.Time) BerlinClock {
	return BerlinClock{
		HoursFive:    t.Hour() / 5,
		HoursOne:     t.Hour() % 5,
		MinutesFive:  t.Minute() / 5,
		MinutesOne:   t.Minute() % 5,
		IsLeapSecond: t.Second() == 0 && t.Nanosecond() == 0,
	}
}

// Testet die Berechnungslogik der Berliner Uhr mit einer spezifischen Zeit
func TestBerlinClockLogic(t *testing.T) {
	testTime := time.Date(2026, 1, 1, 16, 37, 2, 0, time.UTC)
	res := calculateBerlinClock(testTime)

	if res.HoursFive != 3 || res.HoursOne != 1 || res.MinutesFive != 7 || res.MinutesOne != 2 || !res.IsLeapSecond {
		t.Errorf("Berechnungsfehler: %+v", res)
	}
}
