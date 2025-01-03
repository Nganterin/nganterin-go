package helpers

import "time"

func GetDaysFromCheckInCheckOut(checkIn time.Time, checkOut time.Time) int {
	duration := checkOut.Sub(checkIn)
	days := int(duration.Hours() / 24)

	return days
}
