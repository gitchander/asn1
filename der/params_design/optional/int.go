package optional

import (
	"strconv"
)

const (
	//absentValue = "<missing>"
	absentValue = "<absent>"
)

// optional int
type OptInt struct {
	Optional
	Value int
}

func (v OptInt) String() string {
	if v.Present {
		return strconv.Itoa(v.Value)
	}
	return absentValue
}

func Int(v int) OptInt {
	return OptInt{
		Optional: Optional{
			Present: true,
		},
		Value: v,
	}
}
