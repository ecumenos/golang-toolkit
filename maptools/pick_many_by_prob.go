package maptools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
)

// PickManyByProb ...
func PickManyByProb[T string](values map[T]float64, l int) ([]T, error) {
	result := make([]T, 0, l)
	for {
		val, err := PickOneByProb(values)
		if err != nil {
			return nil, customerror.NewToolkitFailure(err, "[PickManyByProb] can not pick several random from input")
		}
		if includes(result, val) {
			continue
		}
		result = append(result, val)
		if len(result) == l {
			break
		}
	}

	return result, nil
}

// includes ...
func includes[T string](values []T, v T) bool {
	for _, val := range values {
		if val == v {
			return true
		}
	}

	return false
}
