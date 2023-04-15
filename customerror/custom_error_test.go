package customerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCast(t *testing.T) {
	errMsg := "error message"
	ce := Cast(errors.New(errMsg))
	assert.NotNil(t, ce)
	assert.Equal(t, ErrorClass, ce.Class)
	assert.Equal(t, DefaultErrorCode, ce.Err.Code)
	assert.Equal(t, errMsg, ce.Err.Message)

	code := NotImplementedErrorCode
	ce = Cast(NewError(nil, errMsg, code))
	assert.NotNil(t, ce)
	assert.Equal(t, ErrorClass, ce.Class)
	assert.Equal(t, code, ce.Err.Code)
	assert.Equal(t, errMsg, ce.Err.Message)

	code = TokenIsInvalidFailureCode
	failure := Failure{
		Code:    code,
		Message: errMsg,
	}
	ce = Cast(NewFailure(nil, failure))
	assert.NotNil(t, ce)
	assert.Equal(t, FailureClass, ce.Class)
	assert.Equal(t, code, ce.Fail.Code)
	assert.Equal(t, errMsg, ce.Fail.Message)
}

func TestCastStrict(t *testing.T) {
	errMsg := "error message"
	ce := CastStrict(errors.New(errMsg))
	assert.Nil(t, ce)

	code := NotImplementedErrorCode
	ce = CastStrict(NewError(nil, errMsg, code))
	assert.NotNil(t, ce)
	assert.Equal(t, ErrorClass, ce.Class)
	assert.Equal(t, code, ce.Err.Code)
	assert.Equal(t, errMsg, ce.Err.Message)

	code = TokenIsInvalidFailureCode
	failure := Failure{
		Code:    code,
		Message: errMsg,
	}
	ce = CastStrict(NewFailure(nil, failure))
	assert.NotNil(t, ce)
	assert.Equal(t, FailureClass, ce.Class)
	assert.Equal(t, code, ce.Fail.Code)
	assert.Equal(t, errMsg, ce.Fail.Message)
}

func TestEquals(t *testing.T) {
	errMsg := "error message"
	assert.True(t, Equals(errors.New(errMsg), errors.New(errMsg)))
	assert.False(t, Equals(errors.New(errMsg), nil))
	assert.False(t, Equals(nil, errors.New(errMsg)))
	assert.True(t, Equals(nil, nil))
	assert.True(t, Equals(NewError(nil, errMsg, DefaultErrorCode), NewError(nil, errMsg, DefaultErrorCode)))
	assert.False(t, Equals(NewError(nil, errMsg, DefaultErrorCode), NewError(nil, errMsg+"1", DefaultErrorCode)))
	assert.True(t, Equals(NewFailure(nil, Failure{Message: errMsg}), NewFailure(nil, Failure{Message: errMsg})))
	assert.False(t, Equals(NewFailure(nil, Failure{Message: errMsg}), NewFailure(nil, Failure{Message: errMsg + "1"})))
}
