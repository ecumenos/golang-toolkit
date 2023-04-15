package version

import (
	"testing"

	"github.com/blang/semver/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSemverVersion(t *testing.T) {
	tCase := map[string]struct {
		input  string
		output semver.Version
	}{
		"should be successful for 0.0.0": {
			input:  "0.0.0",
			output: semver.MustParse("0.0.0"),
		},
		"should be successful for 0.1.0": {
			input:  "0.1.0",
			output: semver.MustParse("0.1.0"),
		},
		"should be successful for 99.99.99": {
			input:  "99.99.99",
			output: semver.MustParse("99.99.99"),
		},
	}

	for name, tc := range tCase {
		t.Run(name, func(tt *testing.T) {
			out, err := ParseSemverVersion(tc.input)
			require.NoError(tt, err)
			assert.Equal(tt, tc.output.Major, out.Major)
			assert.Equal(tt, tc.output.Minor, out.Minor)
			assert.Equal(tt, tc.output.Patch, out.Patch)
		})
	}
}

func TestValidateSemver(t *testing.T) {
	tCase := map[string]struct {
		input  string
		output bool
	}{
		"should be successful for 0.0.0": {
			input:  "0.0.0",
			output: true,
		},
		"should be successful for 0.1.0": {
			input:  "0.1.0",
			output: true,
		},
		"should be successful for 99.99.99": {
			input:  "99.99.99",
			output: true,
		},
		"should be failure for 99.99.99.0": {
			input:  "99.99.99.0",
			output: false,
		},
		"should be failure for 99.99": {
			input:  "99.99",
			output: false,
		},
		"should be failure for xxx": {
			input:  "xxx",
			output: false,
		},
	}

	for name, tc := range tCase {
		t.Run(name, func(tt *testing.T) {
			assert.Equal(tt, tc.output, ValidateSemver(tc.input))
		})
	}
}

func TestIncrementSemverPatchVersion(t *testing.T) {
	tCase := map[string]struct {
		input  string
		output string
	}{
		"should be successful for 0.0.0": {
			input:  "0.0.0",
			output: "0.0.1",
		},
		"should be successful for 0.1.0": {
			input:  "0.1.0",
			output: "0.1.1",
		},
		"should be successful for 99.99.99": {
			input:  "99.99.99",
			output: "99.99.100",
		},
	}

	for name, tc := range tCase {
		t.Run(name, func(tt *testing.T) {
			out, err := IncrementSemverPatchVersion(tc.input)
			require.NoError(tt, err)
			assert.Equal(tt, tc.output, out)
		})
	}
}

func TestReturnGreaterOrIncreasePatchVersion(t *testing.T) {
	tCase := map[string]struct {
		inputLeft  string
		inputRight string
		output     string
	}{
		"should be returns 0.0.1 for left 0.0.0 and right 0.0.0": {
			inputLeft:  "0.0.0",
			inputRight: "0.0.0",
			output:     "0.0.1",
		},
		"should be returns 0.0.2 for left 0.0.1 and right 0.0.2": {
			inputLeft:  "0.0.1",
			inputRight: "0.0.2",
			output:     "0.0.2",
		},
		"should be returns 0.1.1 for left 0.1.0 and right 0.1.0": {
			inputLeft:  "0.1.0",
			inputRight: "0.1.0",
			output:     "0.1.1",
		},
		"should be returns 0.1.1 for left 0.1.0 and right 0.1.1": {
			inputLeft:  "0.1.0",
			inputRight: "0.1.1",
			output:     "0.1.1",
		},
		"should be returns 99.99.100 for left 99.99.99 and right 99.99.99": {
			inputLeft:  "99.99.99",
			inputRight: "99.99.99",
			output:     "99.99.100",
		},
		"should be returns 99.99.100 for left 99.99.99 and right 99.99.100": {
			inputLeft:  "99.99.99",
			inputRight: "99.99.100",
			output:     "99.99.100",
		},
	}

	for name, tc := range tCase {
		t.Run(name, func(tt *testing.T) {
			out, err := ReturnGreaterOrIncreasePatchVersion(tc.inputLeft, tc.inputRight)
			require.NoError(tt, err)
			assert.Equal(tt, tc.output, out)
		})
	}
}
