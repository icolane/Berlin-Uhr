package main

import (
	"testing"
	"time"
)

// Testet die Berechnungslogik der Berliner Uhr mit einer spezifischen Zeit
func TestBerlinClockLogic(t *testing.T) {
	testTime := time.Date(2026, 1, 1, 16, 37, 2, 0, time.UTC)
	res := calculateBerlinClock(testTime)

	if res.HoursFive != 3 || res.HoursOne != 1 || res.MinutesFive != 7 || res.MinutesOne != 2 || !res.IsLeapSecond {
		t.Errorf("Berechnungsfehler: %+v", res)
	}
}

// Testet die Berechnungslogik der Berliner Uhr mit der Zeit 00:00:00
func TestZeroTime(t *testing.T) {
	// Testet die Uhrzeit 00:00:00
	testTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	res := calculateBerlinClock(testTime)

	if res.HoursFive != 0 || res.HoursOne != 0 || res.MinutesFive != 0 || res.MinutesOne != 0 || res.IsLeapSecond {
		t.Errorf("Berechnungsfehler: %+v", res)
	}
}
