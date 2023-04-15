package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash hashes input string with static salt & dynamic salt.
func Hash(in, ss, ds string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ss+in+ds), 14)
	return string(bytes), err
}

// ValidateHash validates input with static slat & dynamic salt. It returns true if it is ok.
func ValidateHash(in, ss, ds, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(ss+in+ds))
	return err == nil
}
