package customerror

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFailure(t *testing.T) {
	failure := Failure{
		Code:        DefaultErrorCode,
		Message:     "failure message",
		Description: "failure desc",
		StatusCode:  http.StatusBadRequest,
	}
	err := NewFailure(nil, failure)
	assert.NotNil(t, err)
	ce := CastStrict(err)
	assert.NotNil(t, ce)
	assert.Equal(t, ce.failureString(), fmt.Sprintf("code: %d, msg: %s, desc: %s, status_code: %d", failure.Code, failure.Message, failure.Description, failure.StatusCode))
}
