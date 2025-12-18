package service

import "time"

func CalculateAgeAt(dob time.Time, today time.Time) int {
	age := today.Year() - dob.Year()

	if today.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func CalculateAge(dob time.Time) int {
	return CalculateAgeAt(dob, time.Now())
}
