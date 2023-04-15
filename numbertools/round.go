package numbertools

import (
	"math"
)

// RoundFloat64 raounds value
func RoundFloat64(val float64, dec int) float64 {
	mul := math.Pow10(dec)
	return math.Round(val*mul) / mul
}
