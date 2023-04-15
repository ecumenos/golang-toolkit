package randomtools

import (
	"github.com/ecumenos/golang-toolkit/customerror"
	goNanoID "github.com/matoous/go-nanoid"
)

// GetNanoString generate nano string
func GetNanoString(length int) (string, error) {
	id, err := goNanoID.ID(length)
	if err != nil {
		return "", customerror.NewToolkitFailure(err, "[GetNanoString] Can not generate nano id")
	}

	return id, nil
}
