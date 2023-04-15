package elgamal

import (
	"crypto/rand"

	"github.com/ecumenos/golang-toolkit/customerror"
	"golang.org/x/crypto/openpgp/elgamal"
)

// GenerateKeys returns elgamal private-public keys pair.
func GenerateKeys() (*elgamal.PrivateKey, *elgamal.PublicKey, error) {
	privKey := &elgamal.PrivateKey{}
	if err := ElgamalGenerate(privKey, rand.Reader); err != nil {
		return nil, nil, customerror.NewError(err, "can not generate elgamal private-public keys pair", customerror.CryptoGenerationErrorCode)
	}
	pubKey := createElgamalPublicKey(privKey.Y.Bytes())

	return privKey, pubKey, nil
}

// Encrypt encrypts byte slice with elgamal public key.
func Encrypt(pubKey *elgamal.PublicKey, input []byte) ([]byte, error) {
	var err error
	for i := 0; i < 3; i++ {
		enc, err := encrypt(pubKey, input)
		if err == nil {
			return enc, nil
		}
	}

	return nil, customerror.NewError(err, "can not encrypt input with elgamal public key", customerror.CryptoEncryptErrorCode)
}

func encrypt(pubKey *elgamal.PublicKey, data []byte) ([]byte, error) {
	encrypter, err := createElgamalEncryption(pubKey, rand.Reader)
	if err != nil {
		panic(err.Error())
	}
	encrypted, err := encrypter.Encrypt(data)
	if err != nil {
		return nil, customerror.NewError(err, "can not encrypt input with elgamal public key", customerror.CryptoEncryptErrorCode)
	}

	return encrypted, nil
}

// Decrypt decrypts byte slice with elgamal private key.
func Decrypt(privKey *elgamal.PrivateKey, d []byte) ([]byte, error) {
	decrypter := &elgDecrypter{
		k: privKey,
	}
	decrypted, err := decrypter.Decrypt(d)
	if err != nil {
		return nil, customerror.NewError(err, "can not decrypt input with elgamal private key", customerror.CryptoDecryptErrorCode)
	}
	return decrypted, nil
}
