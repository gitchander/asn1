package random

import (
	"math/rand"
)

func RandBool(r *rand.Rand) bool {
	return ((r.Int() & 1) == 1)
}

func RandIntMinMax(r *rand.Rand, min, max int) int {
	return min + r.Intn(max-min)
}
