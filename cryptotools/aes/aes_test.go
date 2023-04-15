package aes

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAES(t *testing.T) {
	key, err := os.ReadFile("aes_key")
	require.NoError(t, err)

	msg := []byte("hello world")
	enc, err := Encrypt(key, msg)
	require.NoError(t, err)

	dec, err := Decrypt(key, enc)
	require.NoError(t, err)

	if !bytes.Equal(msg, dec) {
		t.Fail()
	}
}
