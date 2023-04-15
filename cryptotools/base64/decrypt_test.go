package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	str := "string"
	enc := Encrypt(str)
	dec, err := Decrypt(enc)
	require.NoError(t, err)
	assert.Equal(t, dec, str)
}
