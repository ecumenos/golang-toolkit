package aes

import (
	"crypto/rand"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// GenerateKey returns random 32 byte key for AES-256 cipher.
func GenerateKey() ([]byte, error) {
	b := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(b); err != nil {
		return nil, customerror.NewError(err, "can not generate aes key", customerror.CryptoGenerationErrorCode)
	}

	return b, nil
}
