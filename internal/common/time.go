package common

import "time"

func ParseDate(date string) time.Time {
	d, _ := time.Parse("2006-01-02", date)
	return d
}
