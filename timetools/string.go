package timetools

import (
	"time"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// TimeToString serialize time.Time to string with default time format
func TimeToString(in time.Time) string {
	return in.Format(DefaultTimeFormat)
}

// StringToTime deserialize string to time.Time with default time format
func StringToTime(in string) (time.Time, error) {
	out, err := time.Parse(DefaultTimeFormat, in)
	if err != nil {
		return time.Time{}, customerror.NewToolkitFailure(err, "[StringToTime] Can not convert string into time.Time")
	}

	return out, nil
}
