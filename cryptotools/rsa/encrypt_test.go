package rsa

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	privPEM, err := os.ReadFile("id_rsa")
	require.NoError(t, err)

	pubPEM, err := os.ReadFile("id_rsa.pub")
	require.NoError(t, err)

	privParsed, err := ParseRSAPrivateKeyFromPEMString(string(privPEM))
	require.NoError(t, err)
	pubParsed, err := ParseRSAPublicKeyFromPEMString(string(pubPEM))
	require.NoError(t, err)

	msg := "message"
	encrypted, err := Encrypt(pubParsed, []byte(msg), nil)
	require.NoError(t, err)

	decrypted, err := Decrypt(privParsed, encrypted)
	require.NoError(t, err)

	assert.Equal(t, msg, string(decrypted))
}
