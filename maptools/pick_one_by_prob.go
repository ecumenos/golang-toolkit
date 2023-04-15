package maptools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
	"github.com/ecumenos/golang-toolkit/randomtools"
)

// PickOneByProb ...
func PickOneByProb[T string](values map[T]float64) (T, error) {
	var zero T
	if len(values) == 0 {
		return zero, customerror.NewToolkitFailure(nil, "[GetRandomFromSeveral] Values count is zero")
	}

	var (
		valuesWithZeroCount int
		valuesWith1Count    int
	)
	for _, prob := range values {
		if prob == 0 {
			valuesWithZeroCount++
		}
		if prob == 1 {
			valuesWith1Count++
		}
	}
	if valuesWithZeroCount == len(values) {
		return zero, customerror.NewToolkitFailure(nil, "[GetRandomFromSeveral] All values are zero")
	}
	if valuesWith1Count > 1 {
		values = PrepareMapToPickRandomValue(values)
	}

	preparedValues := Clone(values)
	for value, prob := range preparedValues {
		preparedValues[value] = randomtools.PrepareProbability(prob - 0.01)
	}

	tempValues := make(map[T]float64)
	for {
		var iterValues map[T]float64
		if len(tempValues) == 0 {
			iterValues = Clone(preparedValues)
		} else {
			iterValues = Clone(tempValues)
		}
		tempValues = make(map[T]float64)

		for value, prob := range iterValues {
			if r, _ := randomtools.GetRandBool(prob); r {
				tempValues[value] = prob
			}
		}

		if len(tempValues) == 1 {
			for value := range tempValues {
				return value, nil
			}
		}
	}
}

// PrepareMapToPickRandomValue ...
func PrepareMapToPickRandomValue[T string](values map[T]float64) map[T]float64 {
	var totalScore float64
	for _, score := range values {
		totalScore += score
	}

	out := make(map[T]float64, len(values))
	for key, value := range values {
		out[key] = value / totalScore
	}

	return out
}
