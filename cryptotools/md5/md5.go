package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 hashes with MD5 Sum input string.
func MD5(in string) string {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:])
}
