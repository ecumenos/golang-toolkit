package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// KeyPair is RSA key pair.
type KeyPair struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

// GenerateKeyPair returns private & public keys.
func GenerateKeyPair() (KeyPair, error) {
	// The GenerateKey method takes in a reader that returns random bits, and
	// the number of bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return KeyPair{}, customerror.NewError(err, "can not generate RSA key", customerror.CryptoGenerationErrorCode)
	}

	// The public key is a part of the *rsa.PrivateKey struct

	return KeyPair{
		Private: privateKey,
		Public:  &privateKey.PublicKey,
	}, nil
}

// ExportRSAPrivateKeyAsPEMString converts private key to string.
func ExportRSAPrivateKeyAsPEMString(privkey *rsa.PrivateKey) string {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)

	return string(privkeyPEM)
}

// ParseRSAPrivateKeyFromPEMString converts string private key to private key.
func ParseRSAPrivateKeyFromPEMString(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, customerror.NewError(nil, "failed to parse PEM block containing the key", customerror.DefaultCryptoErrorCode)
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, customerror.NewError(err, "can not parse PKCS1 private key", customerror.DefaultCryptoErrorCode)
	}

	return priv, nil
}

// ExportRSAPublicKeyAsPEMString converts public key to string.
func ExportRSAPublicKeyAsPEMString(pubkey *rsa.PublicKey) (string, error) {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", customerror.NewError(err, "can not marshal PKIX public key", customerror.DefaultCryptoErrorCode)
	}
	pubkeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPEM), nil
}

// ParseRSAPublicKeyFromPEMString converts string public key to public key.
func ParseRSAPublicKeyFromPEMString(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, customerror.NewError(nil, "failed to parse PEM block containing the key", customerror.DefaultCryptoErrorCode)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, customerror.NewError(err, "can not parse PKIX public key", customerror.DefaultCryptoErrorCode)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}

	return nil, customerror.NewError(nil, "key type is not RSA", customerror.DefaultCryptoErrorCode)
}
