package date

func YearIsLeap(year int) bool {
	if year < passageYear {
		if year < 1 {
			year++
		}
		return ((year % 4) == 0)
	}
	return (((year%4 == 0) && (year%100 != 0)) || (year%400 == 0))
}

func DaysInYear(year int) int {
	if YearIsLeap(year) {
		return 366
	}
	return 365
}
