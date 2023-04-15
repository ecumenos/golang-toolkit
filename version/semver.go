package version

import (
	"fmt"
	"net/http"

	"github.com/blang/semver/v4"
	"github.com/ecumenos/golang-toolkit/customerror"
)

// ParseSemverVersion returns parsed semver.Version.
func ParseSemverVersion(v string) (semver.Version, error) {
	out, err := semver.Parse(v)
	if err != nil {
		return semver.Version{}, customerror.NewToolkitFailure(err, "[ParseVersion] can not parse version")
	}

	return out, nil
}

// ValidateSemver return true if input version is valid.
func ValidateSemver(v string) bool {
	_, err := ParseSemverVersion(v)
	return err == nil
}

// IncrementSemverPatchVersion returns version with incremented patch version.
func IncrementSemverPatchVersion(v string) (string, error) {
	version, err := ParseSemverVersion(v)
	if err != nil {
		return "", err
	}
	if err := version.IncrementPatch(); err != nil {
		return "", customerror.NewToolkitFailure(err, "[IncrementPatchVersion] can not increment version")
	}

	return version.String(), nil
}

// CompareSemver compares version left to right:
//
// -1 == left is less than right;
//
// 0 == left is equal to right;
//
// 1 == left is greater than right;
func CompareSemver(left, right string) (int, error) {
	leftVersion, err := ParseSemverVersion(left)
	if err != nil {
		return 0, err
	}
	rightVersion, err := ParseSemverVersion(right)
	if err != nil {
		return 0, err
	}

	return leftVersion.Compare(rightVersion), nil
}

// ReturnGreaterOrIncreasePatchVersion returns version.
//
// in case left is greater than right: returns left;
//
// in case left equals to right: returns left with increased patch;
//
// in case right is greater than left: returns right;
func ReturnGreaterOrIncreasePatchVersion(left, right string) (string, error) {
	result, err := CompareSemver(left, right)
	if err != nil {
		return "", err
	}

	switch result {
	case -1:
		return right, nil
	case 0:
		return IncrementSemverPatchVersion(left)
	case 1:
		return left, nil
	}

	return "", customerror.NewFailure(nil, customerror.Failure{
		Code:        customerror.DefaultErrorCode,
		Message:     "incorrect version comparation.",
		Description: fmt.Sprintf("can not compare semver versions. Expects 1, 0, or -1 but received %d", result),
		StatusCode:  http.StatusInternalServerError,
	})
}
