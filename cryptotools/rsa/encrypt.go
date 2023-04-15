package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Encrypt encrypts input with public key.
func Encrypt(pk *rsa.PublicKey, in, label []byte) ([]byte, error) {
	out, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pk, in, label)
	if err != nil {
		return nil, customerror.NewError(err, "can not encrypt with RSA", customerror.CryptoEncryptErrorCode)
	}

	return out, nil
}
