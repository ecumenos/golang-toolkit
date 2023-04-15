package dsa

import (
	"crypto/dsa"
	"crypto/rand"
	"io"
	"math/big"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// generateDSA generates a dsa keypair.
func generateDSA(priv *dsa.PrivateKey, rand io.Reader) error {
	// put our paramters in
	priv.P = param.P
	priv.Q = param.Q
	priv.G = param.G
	// generate the keypair
	return dsa.GenerateKey(priv, rand)
}

// createDSAPublicKey creates i2p dsa public key given its public component.
func createDSAPublicKey(Y *big.Int) *dsa.PublicKey {
	return &dsa.PublicKey{
		Parameters: param,
		Y:          Y,
	}
}

// createDSAPrivkey creates i2p dsa private key given its public component.
func createDSAPrivkey(X *big.Int) *dsa.PrivateKey {
	if X.Cmp(dsap) != -1 {
		return nil
	}

	Y := new(big.Int)
	Y.Exp(dsag, X, dsap)

	return &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: param,
			Y:          Y,
		},
		X: X,
	}
}

// DSAPrivateKey is DSA private key that has 20 bytes.
type DSAPrivateKey [20]byte

// NewSigner create a new dsa signer.
func (k DSAPrivateKey) NewSigner() (Signer, error) {
	return &DSASigner{
		k: createDSAPrivkey(new(big.Int).SetBytes(k[:])),
	}, nil
}

// Public returns public key from DSA private key.
func (k DSAPrivateKey) Public() (DSAPublicKey, error) {
	p := createDSAPrivkey(new(big.Int).SetBytes(k[:]))
	if p == nil {
		return DSAPublicKey{}, customerror.NewError(nil, "invalid dsa private key format", customerror.DefaultCryptoErrorCode)
	}

	var pk DSAPublicKey
	copy(pk[:], p.Y.Bytes())

	return pk, nil
}

// Generate returns DSA private key  from parent DSA private key.
func (k DSAPrivateKey) Generate() (s DSAPrivateKey, err error) {
	dk := new(dsa.PrivateKey)
	err = generateDSA(dk, rand.Reader)
	if err == nil {
		copy(k[:], dk.X.Bytes())
		s = k
	}

	return
}
