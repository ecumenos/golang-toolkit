package randomtools

import (
	"crypto/rand"
	"math/big"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// GetRandIntInRange ...
func GetRandIntInRange(min int, max int) (int, error) {
	if min >= max {
		return 0, customerror.NewToolkitFailure(nil, "[GetRandIntInRange] min >= max")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandIntInRange] Can not generate random value")
	}

	return int(int64(min) + n.Int64()), nil
}

// GetRandInt ...
func GetRandInt(max int) (int, error) {
	out, err := GetRandIntInRange(0, max)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[GetRandInt] Can not generate int in range from 0 to max")
	}

	return out, nil
}
