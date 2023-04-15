package randomtools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
)

// GetRandBool genetate bool by probability of true response.
func GetRandBool(trueProb float64) (bool, error) {
	n, err := GetRandFloat64(1)
	if err != nil {
		return false, customerror.NewToolkitFailure(err, "[GetRandomBool] Can not generate random float64 for generation boolean")
	}
	return n < trueProb, nil
}
