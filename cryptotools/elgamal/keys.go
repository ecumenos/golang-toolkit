package elgamal

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/ecumenos/golang-toolkit/customerror"
	"golang.org/x/crypto/openpgp/elgamal"
)

// ElgamalGenerate generate an elgamal key pair.
func ElgamalGenerate(priv *elgamal.PrivateKey, rand io.Reader) (err error) {
	priv.P = elgp
	priv.G = elgg
	xBytes := make([]byte, priv.P.BitLen()/8)
	_, err = io.ReadFull(rand, xBytes)
	if err != nil {
		return customerror.NewError(err, "can not read x bytes for elgamal generation", customerror.CryptoGenerationErrorCode)
	}
	// set private key
	priv.X = new(big.Int).SetBytes(xBytes)
	// compute public key
	priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)

	return nil
}

// createElgamalPublicKey creates an elgamal public key from byte slice.
func createElgamalPublicKey(data []byte) *elgamal.PublicKey {
	if len(data) != 256 {
		return nil
	}

	return &elgamal.PublicKey{
		G: elgg,
		P: elgp,
		Y: new(big.Int).SetBytes(data),
	}
}

// createElgamalPrivateKey creates an elgamal private key from byte slice.
func createElgamalPrivateKey(data []byte) *elgamal.PrivateKey {
	if len(data) != 256 {
		return nil
	}

	x := new(big.Int).SetBytes(data)
	y := new(big.Int).Exp(elgg, x, elgp)

	return &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			Y: y,
			G: elgg,
			P: elgp,
		},
		X: x,
	}
}

// createElgamalEncryption creates a new elgamal encryption session.
func createElgamalEncryption(pub *elgamal.PublicKey, rand io.Reader) (*ElgamalEncryption, error) {
	kbytes := make([]byte, 256)
	k := new(big.Int)
	var err error
	for err == nil {
		_, err = io.ReadFull(rand, kbytes)
		k = new(big.Int).SetBytes(kbytes)
		k = k.Mod(k, pub.P)
		if k.Sign() != 0 {
			break
		}
	}
	if err != nil {
		return nil, customerror.NewError(err, "can not create elgamal bytes for enc", customerror.CryptoEncryptErrorCode)
	}

	return &ElgamalEncryption{
		p:  pub.P,
		a:  new(big.Int).Exp(pub.G, k, pub.P),
		b1: new(big.Int).Exp(pub.Y, k, pub.P),
	}, nil
}

// ElgPublicKey is elgamal public key.
type ElgPublicKey [256]byte

// ElgPrivateKey is elgamal private key.
type ElgPrivateKey [256]byte

// Len returns elgamal public key length.
func (elg ElgPublicKey) Len() int {
	return len(elg)
}

// NewEncrypter is encryptor constructor.
func (elg ElgPublicKey) NewEncrypter() (Encrypter, error) {
	k := createElgamalPublicKey(elg[:])
	return createElgamalEncryption(k, rand.Reader)
}

// Len returns elgamal private key length.
func (elg ElgPrivateKey) Len() int {
	return len(elg)
}

// NewDecrypter is decryptor constructor.
func (elg ElgPrivateKey) NewDecrypter() (Decrypter, error) {
	return &elgDecrypter{
		k: createElgamalPrivateKey(elg[:]),
	}, nil
}
