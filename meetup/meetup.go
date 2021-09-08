package meetup

import "time"

type WeekSchedule int

const (
	First = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(ordinal WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	bom := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().UTC().Location())
	eom := bom.AddDate(0, 1, 0).Add(-time.Nanosecond)
	lom := bom.AddDate(0, 0, 18)
	switch ordinal {
	case First:
		return lookForFirstWeekDay(bom, weekday, 0, 1)
	case Second:
		return lookForFirstWeekDay(bom, weekday, 1, 1)
	case Third:
		return lookForFirstWeekDay(bom, weekday, 2, 1)
	case Fourth:
		return lookForFirstWeekDay(bom, weekday, 3, 1)
	case Last:
		return lookForFirstWeekDay(eom, weekday, 0, -1)
	default:
		return lookForFirstWeekDay(lom, weekday, 0, -1)
	}
}

func lookForFirstWeekDay(date time.Time, weekday time.Weekday, pos int, dir int) int {
	for {
		for weekday != date.Weekday() {
			date = date.AddDate(0, 0, dir)
		}
		if pos == 0 {
			return date.Day()
		}
		pos--
		date = date.AddDate(0, 0, dir)
	}
}
