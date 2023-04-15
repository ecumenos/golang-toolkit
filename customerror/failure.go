package customerror

import (
	"fmt"
	"net/http"
)

// Failure is struct that represents error with status code 4XX. It is expected error.
type Failure struct {
	Code        Code                   `json:"code"`
	Message     string                 `json:"message"`
	Description string                 `json:"description"`
	StatusCode  int                    `json:"status_code" desc:"StatusCode is HTTP status code"`
	Data        map[string]interface{} `json:"data"`
}

func (ce *CustomError) failureString() string {
	if ce == nil {
		return "<nil> errror"
	}
	if ce.Class != FailureClass {
		return ""
	}
	var data string
	if ce.Fail.Data != nil {
		data = fmt.Sprintf(", data: %+v", ce.Fail.Data)
	}

	return fmt.Sprintf("code: %d, msg: %s, desc: %s, status_code: %d%s", ce.Fail.Code, ce.Fail.Message, ce.Fail.Description, ce.Fail.StatusCode, data)
}

// NewFailure is constructor for failure custom errors.
func NewFailure(err error, f Failure) error {
	return &CustomError{
		Class:       FailureClass,
		Fail:        f,
		ExtendedErr: Cast(err),
	}
}

// NewRequestValidationFailure is constructor for request validation failure custom errors.
func NewRequestValidationFailure(desc string, data map[string]interface{}) error {
	return NewFailure(nil, Failure{
		Code:        RequestValidationFailureCode,
		Message:     "validation error",
		Description: desc,
		StatusCode:  http.StatusBadRequest,
		Data:        data,
	})
}

// NewResponseValidationFailure is constructor for response validation failure custom errors.
func NewResponseValidationFailure(desc string, data map[string]interface{}) error {
	return NewFailure(nil, Failure{
		Code:        ResponseValidationFailureCode,
		Message:     "validation error",
		Description: desc,
		StatusCode:  http.StatusBadRequest,
		Data:        data,
	})
}

// NewToolkitFailure is constructor for toolkit failure custom errors.
func NewToolkitFailure(err error, desc string) error {
	return NewFailure(err, Failure{
		Code:        ToolkitErrorCode,
		Message:     "toolkit error",
		Description: desc,
		StatusCode:  http.StatusInternalServerError,
	})
}
