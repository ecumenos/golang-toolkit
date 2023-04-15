package randomtools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
)

const (
	asciiFirstLowcaseLetterCode = 97
	asciiLastLowcaseLetterCode  = 122
)

// GetRandString generate random string
func GetRandString(len int) (string, error) {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		n, err := GetRandIntInRange(asciiFirstLowcaseLetterCode, asciiLastLowcaseLetterCode)
		if err != nil {
			return "", customerror.NewToolkitFailure(err, "[GetRandString] Can not generate random int value for string generation")
		}
		bytes[i] = byte(n)
	}

	return string(bytes), nil
}
