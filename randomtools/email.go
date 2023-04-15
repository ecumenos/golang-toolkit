package randomtools

import (
	"fmt"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// GetRandEmail generate random email
func GetRandEmail() (string, error) {
	username, err := GetRandString(100)
	if err != nil {
		return "", customerror.NewToolkitFailure(err, "[GetRandEmail] Can not random username for email")
	}

	domainname, err := GetRandString(5)
	if err != nil {
		return "", customerror.NewToolkitFailure(err, "[GetRandEmail] Can not random domainname for email")
	}

	extension, err := GetRandString(3)
	if err != nil {
		return "", customerror.NewToolkitFailure(err, "[GetRandEmail] Can not random extension for email")
	}

	return fmt.Sprintf("%s@%s.%s", username, domainname, extension), nil
}

// GetRandEmails generate random emails
func GetRandEmails(len int) ([]string, error) {
	emails := make([]string, len)

	for i := 0; i < len; i++ {
		email, err := GetRandEmail()
		if err != nil {
			return nil, customerror.NewToolkitFailure(err, "[GetRandEmails] Can not random email for emails generation")
		}
		emails[i] = email
	}

	return emails, nil
}
