package rsa

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateKeys(t *testing.T) {
	// Create the keys
	keyPair, err := GenerateKeyPair()
	require.NoError(t, err)

	// Export the keys to pem string
	privPem := ExportRSAPrivateKeyAsPEMString(keyPair.Private)
	f, err := os.Create("id_rsa")
	require.NoError(t, err)
	_, err = f.Write([]byte(privPem))
	require.NoError(t, err)
	require.NoError(t, f.Close())

	pubPem, err := ExportRSAPublicKeyAsPEMString(keyPair.Public)
	require.NoError(t, err)
	f, err = os.Create("id_rsa.pub")
	require.NoError(t, err)
	_, err = f.Write([]byte(pubPem))
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// Import the keys from pem string
	privParsed, err := ParseRSAPrivateKeyFromPEMString(privPem)
	require.NoError(t, err)
	pubParsed, err := ParseRSAPublicKeyFromPEMString(pubPem)
	require.NoError(t, err)

	// Export the newly imported keys
	privParsedPem := ExportRSAPrivateKeyAsPEMString(privParsed)
	pubParsedPem, err := ExportRSAPublicKeyAsPEMString(pubParsed)
	require.NoError(t, err)

	t.Log(privParsedPem)
	t.Log(pubParsedPem)

	// Check that the exported/imported keys match the original keys
	if privPem != privParsedPem || pubPem != pubParsedPem {
		t.Log("Failure: Export and Import did not result in same Keys")
		t.Fail()
	} else {
		t.Log("Success")
	}
}
