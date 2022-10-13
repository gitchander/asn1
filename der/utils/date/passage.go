package date

import "time"

const (
	passageYear  = 1582
	passageMonth = time.October
)

type ymdDate struct {
	year  int
	month time.Month
	day   int
}

var (
	// First Julian Day -4713-01-01
	firstJulianDay = ymdDate{
		year:  -4713,
		month: time.January,
		day:   1,
	}

	// Last Julian Day 1582-10-04
	lastJulianDay = ymdDate{
		year:  passageYear,
		month: passageMonth,
		day:   4,
	}

	// First Gregorian Day 1582-10-15
	firstGregorianDay = ymdDate{
		year:  passageYear,
		month: passageMonth,
		day:   15,
	}
)
