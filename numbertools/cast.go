package numbertools

import (
	"fmt"
	"strconv"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Float64ToBytes ...
func Float64ToBytes(in float64) []byte {
	return []byte(Float64ToString(in))
}

// BytesToFloat64 ...
func BytesToFloat64(in []byte) (float64, error) {
	return StringToFloat64(string(in))
}

// Float64ToString ...
func Float64ToString(in float64) string {
	return fmt.Sprint(in)
}

// StringToFloat64 ...
func StringToFloat64(in string) (float64, error) {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[StringToFloat64] Can not convert float64 to string")
	}

	return out, nil
}

// StringToInt ...
func StringToInt(in string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[StringToFloat64] Can not convert string to int")
	}

	return out, nil
}
