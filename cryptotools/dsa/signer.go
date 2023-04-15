package dsa

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha1"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// type for signing data
type Signer interface {
	// sign data with our private key by calling SignHash after hashing the data we are given
	// return signature or nil signature and error if an error happened
	Sign(data []byte) (sig []byte, err error)

	// sign hash of data with our private key
	// return signature or nil signature and error if an error happened
	SignHash(h []byte) (sig []byte, err error)
}

type DSASigner struct {
	k *dsa.PrivateKey
}

func (ds *DSASigner) Sign(data []byte) ([]byte, error) {
	h := sha1.Sum(data)

	return ds.SignHash(h[:])
}

func (ds *DSASigner) SignHash(h []byte) ([]byte, error) {
	r, s, err := dsa.Sign(rand.Reader, ds.k, h)
	if err != nil {
		return nil, customerror.NewError(err, "can not sign with dsa", customerror.CryptoEncryptErrorCode)
	}
	sig := make([]byte, 40)
	rb := r.Bytes()
	rl := len(rb)
	copy(sig[20-rl:20], rb)
	sb := s.Bytes()
	sl := len(sb)
	copy(sig[20+(20-sl):], sb)

	return sig, nil
}

func (k DSAPrivateKey) Len() int {
	return len(k)
}
