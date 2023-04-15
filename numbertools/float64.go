package numbertools

import (
	"sort"
	"strconv"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// ParseFloat64 ...
func ParseFloat64(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[ParseFloat64] Can not convert string to float64")
	}

	return f, nil
}

// ParseFloat32 ...
func ParseFloat32(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[ParseFloat32] Can not convert string to float32")
	}

	return float32(f), nil
}

// Average ...
func Average(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}

	var total float64
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

// Median ...
func Median(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}

	sort.Float64s(x)

	half := len(x) / 2
	if isOdd(x) {
		return x[half]
	}

	return (x[half-1] + x[half]) / 2
}

// isOdd ...
func isOdd(x []float64) bool {
	return len(x)%2 != 0
}
