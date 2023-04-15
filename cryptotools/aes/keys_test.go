package aes

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	require.NoError(t, err)
	t.Log(string(key))

	f, err := os.Create("aes_key")
	require.NoError(t, err)

	_, err = f.Write(key)
	require.NoError(t, err)
	require.NoError(t, f.Close())

}
