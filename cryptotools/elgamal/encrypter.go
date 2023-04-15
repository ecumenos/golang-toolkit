package elgamal

import (
	"crypto/sha256"
	"math/big"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Encrypter encrypts data.
type Encrypter interface {
	// encrypt a block of data.
	// return encrypted block or nil and error if an error happened.
	Encrypt(data []byte) ([]byte, error)
}

// PublicEncryptionKey is interface for public encryption key.
type PublicEncryptionKey interface {
	// NewEncrypter returns a new encrypter to encrypt data to this public key.
	NewEncrypter() (Encrypter, error)

	// Len returns length of this public key in bytes.
	Len() int
}

// ElgamalEncryption is elgamal encryption struct.
type ElgamalEncryption struct {
	p, a, b1 *big.Int
}

// Encrypt encrypts input byte slice.
func (elg *ElgamalEncryption) Encrypt(data []byte) ([]byte, error) {
	return elg.EncryptPadding(data, true)
}

// EncryptPadding encrypts input byte slice.
func (elg *ElgamalEncryption) EncryptPadding(data []byte, zeroPadding bool) ([]byte, error) {
	if len(data) > 222 {
		return nil, customerror.NewError(nil, "failed to encrypt data, too big for elgamal", customerror.CryptoEncryptErrorCode)
	}
	mbytes := make([]byte, 255)
	mbytes[0] = 0xFF
	copy(mbytes[33:], data)
	// do sha256 of payload
	d := sha256.Sum256(mbytes[33 : len(data)+33])
	copy(mbytes[1:], d[:])
	m := new(big.Int).SetBytes(mbytes)
	// do encryption
	b := new(big.Int).Mod(new(big.Int).Mul(elg.b1, m), elg.p).Bytes()

	var encrypted []byte
	if zeroPadding {
		encrypted = make([]byte, 514)
		copy(encrypted[1:], elg.a.Bytes())
		copy(encrypted[258:], b)
	} else {
		encrypted = make([]byte, 512)
		copy(encrypted, elg.a.Bytes())
		copy(encrypted[256:], b)
	}
	return encrypted, nil
}
