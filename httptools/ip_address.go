package httptools

import (
	"net"
	"net/http"
	"strings"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// ExtractIPAddress returns IP address of request.
func ExtractIPAddress(r *http.Request) (string, error) {
	// Get the IP address from the "X-Forwarded-For" header.
	ipAddress := r.Header.Get("X-Forwarded-For")

	if ipAddress == "" {
		// If the "X-Forwarded-For" header is not set, try to get the IP address from the remote address of the connection.
		ipAddress, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return "", customerror.NewFailure(nil, customerror.Failure{
				Code:        customerror.ExtractIPAddressFromRequestFailureCode,
				Message:     "request is incorrect.",
				Description: "can not extract IP addresses from request's remote address",
				StatusCode:  http.StatusBadRequest,
			})
		}

		return ipAddress, nil
	}

	// If the "X-Forwarded-For" header is set, split it and return the first IP address in the list.
	ipAddresses := strings.Split(ipAddress, ", ")
	if len(ipAddresses) < 1 {
		return "", customerror.NewFailure(nil, customerror.Failure{
			Code:        customerror.ExtractIPAddressFromRequestFailureCode,
			Message:     "request is incorrect.",
			Description: "can not extract IP address from request",
			StatusCode:  http.StatusBadRequest,
		})
	}

	return ipAddresses[0], nil
}
