package ecdsa

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
	privPem, err := ExportECPrivateKey(keyPair.Private)
	require.NoError(t, err)
	f, err := os.Create("ecdsa_key")
	require.NoError(t, err)
	_, err = f.Write([]byte(privPem))
	require.NoError(t, err)
	require.NoError(t, f.Close())

	pubPem, err := ExportECPublicKey(keyPair.Public)
	require.NoError(t, err)
	f, err = os.Create("ecdsa_key.pub")
	require.NoError(t, err)
	_, err = f.Write([]byte(pubPem))
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// Import the keys from pem string
	privParsed, err := ParseECPrivateKeyFromPEM(privPem)
	require.NoError(t, err)
	pubParsed, err := ParseECPublicKeyFromPEM(pubPem)
	require.NoError(t, err)

	// Export the newly imported keys
	privParsedPem, err := ExportECPrivateKey(privParsed)
	require.NoError(t, err)
	pubParsedPem, err := ExportECPublicKey(pubParsed)
	require.NoError(t, err)

	t.Log(string(privParsedPem))
	t.Log(string(pubParsedPem))

	// Check that the exported/imported keys match the original keys
	if string(privPem) != string(privParsedPem) || string(pubPem) != string(pubParsedPem) {
		t.Log("Failure: Export and Import did not result in same Keys")
		t.Fail()
	} else {
		t.Log("Success")
	}
}
