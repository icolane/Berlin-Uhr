package main

import (
	"testing"
	"time"
)

func TestBerlinClockLogic(t *testing.T) {
	// Teste 23:59:59 (Maximale Zeit)
	testCase1 := time.Date(2024, 1, 1, 23, 59, 59, 0, time.UTC)
	result1 := getBerlinClock(testCase1)

	if result1.HoursFive != 4 || result1.HoursOne != 3 {
		t.Errorf("Stunden falsch bei 23:00: erwartet 4x5 + 3, erhalten %d*5 + %d", result1.HoursFive, result1.HoursOne)
	}
	if result1.MinutesFive != 11 || result1.MinutesOne != 4 {
		t.Errorf("Minuten falsch bei :59: erwartet 11x5 + 4, erhalten %d*5 + %d", result1.MinutesFive, result1.MinutesOne)
	}

	// Teste 00:00:00 (Minimum / Mitternacht)
	testCase2 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	result2 := getBerlinClock(testCase2)

	if result2.HoursFive != 0 || result2.HoursOne != 0 || result2.MinutesFive != 0 || result2.MinutesOne != 0 {
		t.Errorf("Nullzeit falsch berechnet: %+v", result2)
	}

	if result2.IsLeapSecond != true {
		t.Errorf("Sekunden-Blinker falsch bei 00:00:00: erwartet true (gerade), erhalten false")
	}
}
