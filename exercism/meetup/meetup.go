package meetup

import "time"

type WeekSchedule int

const (
	First = WeekSchedule(iota)
	Second
	Third
	Fourth
	Teenth
	Last
)

//"First Monday of January 2022"
//Day returns the day of the week for a meeting for the  given restrictions
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var options []int

	date := time.Date(year, month, 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))

	for date.Month() == month {
		if date.Weekday() == wDay {
			options = append(options, date.Day())
		}
		date = date.AddDate(0, 0, 1) // increment 1 day
	}

	switch wSched {
	case First:
		return options[0]
	case Second:
		return options[1]
	case Third:
		return options[2]
	case Fourth:
		return options[3]
	case Last:
		return options[len(options)-1]
	case Teenth:
		for _, v := range options {
			if v > 12 {
				return v
			}
		}
	}

	panic("Invalid wSched")

}
