package dsa

import (
	"github.com/ecumenos/golang-toolkit/customerror"
)

// GenerateKeys returns private-public DSA key.
func GenerateKeys() (DSAPrivateKey, DSAPublicKey, error) {
	var sk DSAPrivateKey
	sk, err := sk.Generate()
	if err != nil {
		return DSAPrivateKey{}, DSAPublicKey{}, customerror.NewError(err, "can not generate DSA secret key", customerror.CryptoGenerationErrorCode)
	}

	zeros := 0
	for _, b := range sk {
		if b == 0 {
			zeros++
		}
	}
	if zeros == len(sk) {
		return DSAPrivateKey{}, DSAPublicKey{}, customerror.NewError(err, "dsa key generation yielded all zeros", customerror.DefaultCryptoErrorCode)
	}

	pk, err := sk.Public()
	if err != nil {
		return DSAPrivateKey{}, DSAPublicKey{}, customerror.NewError(err, "can not generate public key from public", customerror.CryptoGenerationErrorCode)
	}

	return sk, pk, nil
}

// Sign signs input byte slice by DSA private key.
func Sign(sk DSAPrivateKey, input []byte) ([]byte, error) {
	signer, err := sk.NewSigner()
	if err != nil {
		return nil, err
	}

	return signer.Sign(input)
}

// Verify verifies signed byte slice with input byte slice and DSA public key.
func Verify(pk DSAPublicKey, input, signed []byte) error {
	verify, err := pk.NewVerifier()
	if err != nil {
		return err
	}

	return verify.Verify(input, signed)
}
