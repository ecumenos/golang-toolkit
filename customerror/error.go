package customerror

import "fmt"

// Error is struct that represents error with status code 5XX.
type Error struct {
	Message string `json:"message"`
	Code    Code   `json:"code"`
}

func (ce *CustomError) errorString() string {
	if ce == nil {
		return "<nil> errror"
	}
	if ce.Class != ErrorClass {
		return ""
	}

	return fmt.Sprintf("code: %d, msg: %s", ce.Err.Code, ce.Err.Message)
}

// NewError is constructor for error custom errors.
func NewError(err error, msg string, code Code) error {
	return &CustomError{
		Class:       ErrorClass,
		Err:         Error{Message: msg, Code: code},
		ExtendedErr: Cast(err),
	}
}
