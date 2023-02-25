package random

import (
	"math"
	"math/rand"
)

func NormFloat64(min float64, max float64) float64 {
	return math.Abs(rand.NormFloat64()*((max-min)/3)) + min
}
