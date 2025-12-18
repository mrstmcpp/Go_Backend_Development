package service

import (
	"testing"
	"time"
)

func TestCalculatedAge_BdayArrived(t *testing.T) {
	dob := time.Date(2007, time.May, 10, 0, 0, 0, 0, time.UTC)
	today := time.Date(2025, time.June, 19, 0, 0, 0, 0, time.UTC)

	age := CalculateAgeAt(dob, today)

	expected := 18
	if age != expected {
		t.Errorf("expected %d, got %d", expected, age)
	}
}

func TestCalculatedAge_BdayNotArrived(t *testing.T) {
	dob := time.Date(2007, time.December, 22, 0, 0, 0, 0, time.UTC)
	today := time.Date(2025, time.June, 19, 0, 0, 0, 0, time.UTC)

	age := CalculateAgeAt(dob, today)

	expected := 17
	if age != expected {
		t.Errorf("expected %d, got %d", expected, age)
	}
}

func TestCalculatedAge_BdayToday(t *testing.T) {
	dob := time.Date(2007, time.June, 19, 0, 0, 0, 0, time.UTC)
	today := time.Date(2025, time.June, 19, 0, 0, 0, 0, time.UTC)

	age := CalculateAgeAt(dob, today)

	expected := 18
	if age != expected {
		t.Errorf("expected %d, got %d", expected, age)
	}
}
