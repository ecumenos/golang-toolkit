package numbertools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64ToBytes(t *testing.T) {
	tCases := map[string]struct {
		in  float64
		out []byte
	}{
		"should convert 1": {
			in:  1,
			out: []byte(fmt.Sprint(1)),
		},
		"should convert 1.1": {
			in:  1.1,
			out: []byte(fmt.Sprint(1.1)),
		},
		"should convert 0.1": {
			in:  0.1,
			out: []byte(fmt.Sprint(0.1)),
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := Float64ToBytes(tc.in)
			assert.Equal(tt, string(tc.out), string(out))
		})
	}
}

func TestBytesToFloat64(t *testing.T) {
	tCases := []struct {
		name string
		in   []byte
		out  float64
	}{
		{
			name: "should convert 1",
			in:   Float64ToBytes(1),
			out:  1,
		},
		{
			name: "should convert 1.1",
			in:   Float64ToBytes(1.1),
			out:  1.1,
		},
		{
			name: "should convert 0.1",
			in:   Float64ToBytes(0.1),
			out:  0.1,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			out, err := BytesToFloat64(tc.in)
			if err != nil {
				tt.Errorf("unexpected error (err=%+v)", err)
			}
			if out != tc.out {
				tt.Errorf("unexpected return (actual=%f, expected=%f)", out, tc.out)
			}
		})
	}
}
