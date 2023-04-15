package customerror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	const (
		msg  = "error msg"
		code = DefaultErrorCode
	)
	err := NewError(nil, msg, code)
	assert.NotNil(t, err)
	ce := CastStrict(err)
	assert.NotNil(t, ce)
	assert.Equal(t, ce.errorString(), fmt.Sprintf("code: %d, msg: %s", code, msg))
}
