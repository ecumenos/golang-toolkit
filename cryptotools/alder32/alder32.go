package alder32

import "hash/adler32"

// Checksum returns the Adler-32 checksum of data.
func Checksum(input []byte) uint32 {
	return adler32.Checksum(input)
}
