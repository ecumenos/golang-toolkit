package httptools

import (
	"net/http"
	"strings"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// ExtractJWTBearerToken returns bearer JWT token from request.
func ExtractJWTBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", customerror.NewFailure(nil, customerror.Failure{
			Code:        customerror.IncorrectTokenFormatFailureCode,
			Message:     "Authorization header missing",
			Description: "Authorization header missing",
			StatusCode:  http.StatusUnauthorized,
		})
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", customerror.NewFailure(nil, customerror.Failure{
			Code:        customerror.IncorrectTokenFormatFailureCode,
			Message:     "token is incorrect.",
			Description: "token must be JWT bearer token.",
			StatusCode:  http.StatusUnauthorized,
		})
	}

	return authHeaderParts[1], nil
}
