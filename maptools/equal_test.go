package maptools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	tCases := map[string]struct {
		in  []map[string]float64
		out bool
	}{
		"should return true": {
			in: []map[string]float64{
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
			},
			out: true,
		},
		"should return false": {
			in: []map[string]float64{
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
					"d": 0.44,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
				},
			},
			out: false,
		},
		"should return false for map with same keys": {
			in: []map[string]float64{
				{
					"a": 0.22,
					"b": 0.32,
					"c": 0.12,
				},
				{
					"a": 0.23,
					"b": 0.32,
					"c": 0.12,
				},
			},
			out: false,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			assert.Equal(tt, tc.out, Equal(tc.in...))
		})
	}
}
