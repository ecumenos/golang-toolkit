package elgamal

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElGamal(t *testing.T) {
	tCases := []struct {
		name  string
		times int
	}{
		{
			name:  "should be ok for 100 times",
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.times; i++ {
				privKey, pubKey, err := GenerateKeys()
				require.NoError(t, err)

				input := make([]byte, 222)
				_, _ = io.ReadFull(rand.Reader, input)
				encrypted, err := Encrypt(pubKey, input)
				require.NoError(t, err)

				output, err := Decrypt(privKey, encrypted)
				require.NoError(t, err)

				if !bytes.Equal(input, output) {
					t.Fail()
				}
			}
		})
	}
}
