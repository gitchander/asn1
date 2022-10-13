package date

import "time"

var monthDays = [...]int{
	31, // January
	28, // February
	31, // March
	30, // April
	31, // May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

func NumberOfDays(year int, month time.Month) int {
	n := monthDays[int(month)-1]
	if YearIsLeap(year) && (month == time.February) {
		n = 29
	}
	return n
}
