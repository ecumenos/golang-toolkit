package numbertools

import (
	"fmt"
	"strconv"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Float64ToBytes converts input float64 value into output slice of bytes value.
func Float64ToBytes(in float64) []byte {
	return []byte(Float64ToString(in))
}

// BytesToFloat64 converts input slice of bytes value into output float64 value.
func BytesToFloat64(in []byte) (float64, error) {
	return StringToFloat64(string(in))
}

// Float64ToString converts input float64 value into output string value.
func Float64ToString(in float64) string {
	return fmt.Sprint(in)
}

// StringToFloat64 converts input string value into output float64 value.
func StringToFloat64(in string) (float64, error) {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[StringToFloat64] Can not convert float64 to string")
	}

	return out, nil
}

// StringToInt converts input string value into output int value.
func StringToInt(in string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[StringToInt] Can not convert string to int")
	}

	return out, nil
}

// StringToInt64 converts input string value into output int64 value.
func StringToInt64(in string) (int64, error) {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, customerror.NewToolkitFailure(err, "[StringToInt64] Can not convert string to int64")
	}

	return out, nil
}
