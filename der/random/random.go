package random

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

func Bool(r *rand.Rand) bool {
	return ((r.Int() & 1) == 1)
}

func RangeInt(r *rand.Rand, min, max int) int {
	return min + r.Intn(max-min)
}
