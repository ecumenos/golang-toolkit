package base64

import (
	"encoding/base64"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Decrypt decrypts input from base64 to string.
func Decrypt(in string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "", customerror.NewToolkitFailure(err, "can not decrypt string from base64")
	}
	return string(data), nil
}
