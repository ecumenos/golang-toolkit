package maptools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tCases := map[string]struct {
		in  []map[string]float64
		out map[string]float64
	}{
		"": {
			in: []map[string]float64{
				{
					"a": 0.1,
					"b": 0.2,
					"c": 0.3,
					"d": 0.4,
				},
				{
					"c": 0.9,
					"d": 0.4,
					"e": 0.5,
					"f": 0.6,
				},
			},
			out: map[string]float64{
				"a": 0.1,
				"b": 0.2,
				"c": 0.6,
				"d": 0.4,
				"e": 0.5,
				"f": 0.6,
			},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := Merge(tc.in...)
			assert.True(tt, Equal(out, tc.out))
		})
	}
}
