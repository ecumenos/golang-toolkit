package base64

import "encoding/base64"

// Encrypt encrypts input by base64.
func Encrypt(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}
