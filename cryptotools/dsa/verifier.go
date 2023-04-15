package dsa

import (
	"crypto/dsa"
	"crypto/sha1"
	"math/big"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Verifier is type for verifying signatures.
type Verifier interface {
	// verify hashed data with this signing key
	// return nil on valid signature otherwise error
	VerifyHash(h, sig []byte) error
	// verify an unhashed piece of data by hashing it and calling VerifyHash
	Verify(data, sig []byte) error
}

// DSAVerifier is DSA verifier.
type DSAVerifier struct {
	k *dsa.PublicKey
}

// DSAPublicKey is DSA public key that has 128 bytes.
type DSAPublicKey [128]byte

// NewVerifier create a new DSA verifier.
func (k DSAPublicKey) NewVerifier() (Verifier, error) {
	return &DSAVerifier{
		k: createDSAPublicKey(new(big.Int).SetBytes(k[:])),
	}, nil
}

// Verify verifies data with a DSA public key.
func (v *DSAVerifier) Verify(data, sig []byte) error {
	h := sha1.Sum(data)

	return v.VerifyHash(h[:], sig)
}

// VerifyHash verifies hash of data with a DSA public key.
func (v *DSAVerifier) VerifyHash(h, sig []byte) error {
	if len(sig) != 40 {
		return customerror.NewError(nil, "bad DSA public key signature size", customerror.DefaultCryptoErrorCode)
	}

	r := new(big.Int).SetBytes(sig[:20])
	s := new(big.Int).SetBytes(sig[20:])
	if dsa.Verify(v.k, h, r, s) {
		return nil
	}

	return customerror.NewError(nil, "invalid DSA public key signature", customerror.DefaultCryptoErrorCode)
}

// Len returns length of DSA public key.
func (k DSAPublicKey) Len() int {
	return len(k)
}
