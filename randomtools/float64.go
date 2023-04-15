package randomtools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
	expRand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// GetRandFloat64InRange ...
func GetRandFloat64InRange(min, max float64) (float64, error) {
	if min >= max {
		return 0, customerror.NewToolkitFailure(nil, "[GetRandFloat64InRange] min >= max")
	}

	rInt, err := GetRandIntInRange(0, 100)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandFloat64InRange] Can not generate int in range 0 to 100")
	}
	s := expRand.NewSource(uint64(rInt))
	return min + expRand.New(s).Float64()*(max-min), nil
}

// GetRandFloat64 ...
func GetRandFloat64(max float64) (float64, error) {
	out, err := GetRandFloat64InRange(0, max)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandFloat64] Can not generate float64 in range 0 to max")
	}
	return out, nil
}

// GetRandFloat64Norm generate random float64 with norm
// stdDev -standart deviation - σ^2; default = 1
// mean - μ (In probability theory, the expected value is a generalization of the weighted average.); default = 0
func GetRandFloat64Norm(stdDev, mean float64) (float64, error) {
	rInt, err := GetRandIntInRange(0, 100)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandFloat64Norm] Can not generate int in range 0 to 100")
	}
	s := expRand.NewSource(uint64(rInt))

	dist := distuv.Normal{
		Mu:    mean,   // Mean of the normal distribution
		Sigma: stdDev, // Standard deviation of the normal distribution
		Src:   s,
	}

	return dist.Rand(), nil
}

// GetRandFloat64NormInRange ...
func GetRandFloat64NormInRange(min, max, stdDev, mean float64) (float64, error) {
	if min >= max {
		return 0, customerror.NewToolkitFailure(nil, "[GetRandFloat64NormInRange] min >= max")
	}

	out, err := getRandFloat64NormInRange(min, max, stdDev, mean, 10)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandFloat64NormInRange] Can not generate float64 in normal distribution")
	}

	return out, nil
}

// getRandFloat64NormInRange ...
func getRandFloat64NormInRange(min, max, stdDev, mean float64, count int) (float64, error) {
	var multiplier float64 = 10

	min *= multiplier
	max *= multiplier
	mean *= multiplier

	r, err := GetRandFloat64Norm(stdDev, mean)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[getRandFloat64NormInRange] Can not generate value by norm distributed")
	}

	if r < min {
		if count == 0 {
			return min / multiplier, nil
		}
		return getRandFloat64NormInRange(min/multiplier, max/multiplier, stdDev, mean/multiplier, count-1)
	}
	if r > max {
		if count == 0 {
			return max / multiplier, nil
		}
		return getRandFloat64NormInRange(min/multiplier, max/multiplier, stdDev, mean/multiplier, count-1)
	}

	return r / multiplier, nil
}
