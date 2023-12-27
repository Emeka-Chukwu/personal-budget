package shared

import "time"

func GetNewDate(duration int, periods int) time.Time {
	noOfDays := int(duration / periods)
	noOfHours := duration % periods
	payDate := time.Now().AddDate(0, 0, noOfDays)
	payDate = payDate.Add(time.Hour * time.Duration(noOfHours))
	return payDate
}
