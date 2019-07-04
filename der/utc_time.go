package der

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gitchander/asn1/der/random"
	"github.com/toelsiba/date"
)

/*

YYMMDDhhmmZ
YYMMDDhhmm+hhmm
YYMMDDhhmm-hhmm
YYMMDDhhmmssZ
YYMMDDhhmmss+hhmm
YYMMDDhhmmss-hhmm

*/

func encodeUTCTime(t time.Time) ([]byte, error) {
	data := make([]byte, 0, 17)
	return appendUTCTime(data, t)
}

func decodeUTCTime(data []byte) (time.Time, error) {
	t, _, err := parseUTCTime(data)
	return t, err
}

func appendUTCTime(data []byte, t time.Time) ([]byte, error) {

	year, month, day := t.Date()

	shortYear := yearCollapse(year)
	if shortYear == -1 {
		return nil, fmt.Errorf("bad convert time to UTCTime: invalid year(%d)", year)
	}
	data = appendTwoDigits(data, shortYear)
	data = appendTwoDigits(data, int(month))
	data = appendTwoDigits(data, day)

	hour, min, sec := t.Clock()

	data = appendTwoDigits(data, hour)
	data = appendTwoDigits(data, min)
	data = appendTwoDigits(data, sec)

	_, offset := t.Zone()
	offsetMinutes := offset / 60

	switch {
	case offsetMinutes == 0:
		return append(data, 'Z'), nil
	case offsetMinutes > 0:
		data = append(data, '+')
	case offsetMinutes < 0:
		data = append(data, '-')
	}

	if offsetMinutes < 0 {
		offsetMinutes = -offsetMinutes
	}

	data = appendTwoDigits(data, offsetMinutes/60) // hours
	data = appendTwoDigits(data, offsetMinutes%60) // mins

	return data, nil
}

func parseUTCTime(data []byte) (time.Time, []byte, error) {

	var err error

	ds := make([]int, 6)
	data, err = parseTwoDigitsSeries(data, ds)
	if err != nil {
		return time.Time{}, nil, err
	}

	var (
		shortYear = ds[0]
		month     = time.Month(ds[1])
		day       = ds[2]
	)

	year := yearExpand(shortYear)

	var (
		hour = ds[3]
		min  = ds[4]
		sec  = ds[5]
	)

	if len(data) < 1 {
		return time.Time{}, nil, errors.New("parse UTCTime: insufficient data length")
	}

	b := data[0]
	data = data[1:]

	var negative bool
	switch b {
	case 'Z':
		t := time.Date(year, month, day, hour, min, sec, 0, time.UTC)
		return t, data, nil
	case '-':
		negative = true
	case '+':
		negative = false
	default:
		return time.Time{}, nil, fmt.Errorf("parse UTCTime: invalid character %q", b)
	}

	ds = make([]int, 2)
	data, err = parseTwoDigitsSeries(data, ds)
	if err != nil {
		return time.Time{}, nil, err
	}
	offsetMinutes := int(ds[0])*60 + int(ds[1])
	if negative {
		offsetMinutes = -offsetMinutes
	}

	const timeInLocal = true
	if timeInLocal {
		t := time.Date(year, month, day, hour, min, sec, 0, time.UTC)
		t = t.Add(time.Minute * time.Duration(-offsetMinutes))
		t = t.In(time.Local)
		return t, data, nil
	}

	loc := time.FixedZone("", offsetMinutes*60)
	t := time.Date(year, month, day, hour, min, sec, 0, loc)
	return t, data, nil
}

func appendTwoDigits(data []byte, x int) []byte {
	var (
		lo = '0' + byte(x%10)
		hi = '0' + byte((x/10)%10)
	)
	return append(data, hi, lo)
}

func parseTwoDigits(data []byte) (int, []byte) {
	if len(data) < 2 {
		return -1, data
	}

	hi, ok := byteToDigit(data[0])
	if !ok {
		return -1, data
	}

	lo, ok := byteToDigit(data[1])
	if !ok {
		return -1, data
	}

	x := hi*10 + lo

	return x, data[2:]
}

func parseTwoDigitsSeries(data []byte, ds []int) ([]byte, error) {
	var d int
	for i := range ds {
		d, data = parseTwoDigits(data)
		if d == -1 {
			return nil, errors.New("invalid series of digits")
		}
		ds[i] = d
	}
	return data, nil
}

func RandomUTCTime(r *rand.Rand) time.Time {

	var (
		year  = random.RangeInt(r, 1950, 2050) // [1950..2049]
		month = time.Month(1 + r.Intn(12))     // [1..12]
		day   = 1 + r.Intn(date.NumberOfDays(year, month))

		hour = r.Intn(24) // [0..23]
		min  = r.Intn(60) // [0..59]
		sec  = r.Intn(60) // [0..59]
	)

	const minutesPerHalfDay = 12 * 60
	offsetMin := random.RangeInt(r, -minutesPerHalfDay, minutesPerHalfDay)
	loc := time.FixedZone("", offsetMin*60)

	t := time.Date(year, month, day, hour, min, sec, 0, loc)

	if year == 1950 {
		t = t.Add(24 * time.Hour) // + 1 day (because if convert it to local time it can be equals 1949)
	} else if year == 2049 {
		t = t.Add(-24 * time.Hour) // - 1 day (because ---||--- 2050)
	}

	return t
}

func UTCTimeSerialize(t time.Time, tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTC_TIME
	}

	n := NewNode(class, tag)
	err := n.SetUTCTime(t)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func UTCTimeDeserialize(n *Node, tag int) (time.Time, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTC_TIME
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return time.Time{}, err
	}

	return n.GetUTCTime()
}

// year: [0..99]
func yearExpand(year int) int {
	if inInterval(year, 0, 50) {
		return year + 2000
	}
	if inInterval(year, 50, 100) {
		return year + 1900
	}
	return -1
}

// year: [1950..2049]
func yearCollapse(year int) int {
	if inInterval(year, 2000, 2050) {
		return year - 2000
	}
	if inInterval(year, 1950, 2000) {
		return year - 1900
	}
	return -1
}

// Value a is in [min..max)
func inInterval(a int, min, max int) bool {
	return (min <= a) && (a < max)
}
