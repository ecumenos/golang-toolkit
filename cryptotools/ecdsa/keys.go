package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// KeyPair is RSA key pair.
type KeyPair struct {
	Private *ecdsa.PrivateKey
	Public  *ecdsa.PublicKey
}

// GenerateKeyPair returns private & public keys.
func GenerateKeyPair() (KeyPair, error) {
	private, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return KeyPair{}, customerror.NewError(err, "can not generate ecdsa key pair.", customerror.CryptoGenerationErrorCode)
	}

	return KeyPair{
		Private: private,
		Public:  &private.PublicKey,
	}, nil
}

// ParseECPrivateKeyFromPEM Parse PEM encoded Elliptic Curve Private Key Structure.
func ParseECPrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, customerror.NewError(nil, "invalid key: Key must be PEM encoded PKCS1 or PKCS8 private key.", customerror.DefaultCryptoErrorCode)
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParseECPrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, customerror.NewError(nil, "key is not a valid ECDSA private key.", customerror.DefaultCryptoErrorCode)
	}

	return pkey, nil
}

// ExportECPrivateKey converts private key to []byte.
func ExportECPrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, customerror.NewError(err, "can not marshal EC private key", customerror.DefaultCryptoErrorCode)
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return pemEncoded, nil
}

// ParseECPublicKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 public key.
func ParseECPublicKeyFromPEM(key []byte) (*ecdsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, customerror.NewError(nil, "invalid key: Key must be PEM encoded PKCS1 or PKCS8 private key.", customerror.DefaultCryptoErrorCode)
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *ecdsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PublicKey); !ok {
		return nil, customerror.NewError(nil, "key is not a valid ECDSA public key.", customerror.DefaultCryptoErrorCode)
	}

	return pkey, nil
}

// ExportECPublicKey converts public key to []byte.
func ExportECPublicKey(publicKey *ecdsa.PublicKey) ([]byte, error) {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, customerror.NewError(err, "can not marshal EC public key", customerror.DefaultCryptoErrorCode)
	}
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return pemEncodedPub, nil
}
