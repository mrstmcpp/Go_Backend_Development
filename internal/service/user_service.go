package service

import "time"

func CalculateAge(dob time.Time) int {

	today := time.Now()
	age := today.Year() - dob.Year()

	//if bday is incoming
	if today.YearDay() < dob.YearDay() {
		age--
	}

	return age

}
