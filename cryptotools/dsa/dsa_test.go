package dsa

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDSA(t *testing.T) {
	sk, pk, err := GenerateKeys()
	require.NoError(t, err)

	data := make([]byte, 512)
	_, err = io.ReadFull(rand.Reader, data)
	require.NoError(t, err)

	signed, err := Sign(sk, data)
	require.NoError(t, err)

	require.NoError(t, Verify(pk, data, signed))
}
