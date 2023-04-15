package sha256

import (
	"crypto/sha256"
)

// Hash hashed input byte slice.
func Hash(input []byte) [32]byte {
	return sha256.Sum256(input)
}
