package rsa

import (
	"crypto"
	"crypto/rsa"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Decrypt decrypts input with private key.
func Decrypt(sk *rsa.PrivateKey, in []byte) ([]byte, error) {
	out, err := sk.Decrypt(nil, in, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, customerror.NewError(err, "can not decrypt input with RSA", customerror.CryptoDecryptErrorCode)
	}

	return out, nil
}
