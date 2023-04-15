package numbertools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat64(t *testing.T) {
	tCases := map[string]struct {
		in  float64
		dec int
		out float64
	}{
		"10.123456 -> 10.12": {
			in:  10.123456,
			dec: 2,
			out: 10.12,
		},
		"101.23456 -> 101.235": {
			in:  101.23456,
			dec: 3,
			out: 101.235,
		},
		"1012.3456 -> 1012.3456": {
			in:  1012.3456,
			dec: 4,
			out: 1012.3456,
		},
		"10.123456 -> 10.1": {
			in:  10.123456,
			dec: 1,
			out: 10.1,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			assert.Equal(tt, tc.out, RoundFloat64(tc.in, tc.dec))
		})
	}
}
