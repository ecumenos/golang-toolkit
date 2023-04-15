package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Encrypt encrypts message by key with AES cipher.
func Encrypt(key, msg []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, customerror.NewError(err, "can not create new aes cipher", customerror.CryptoGenerationErrorCode)
	}

	encrypted := make([]byte, aes.BlockSize+len(msg))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, customerror.NewError(err, "can not encrypt with aes", customerror.CryptoEncryptErrorCode)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], msg)

	return encrypted, nil
}

// Decrypt decrypts message by key with AES cipher.
func Decrypt(key, encrypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, customerror.NewError(err, "can not create new aes cipher", customerror.CryptoGenerationErrorCode)
	}

	if len(encrypted) < aes.BlockSize {
		return nil, customerror.NewError(err, "invalid aes ciphertext block size", customerror.CryptoDecryptErrorCode)
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted, nil
}
