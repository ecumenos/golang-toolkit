package customerror

import "errors"

// CustomError is union of error & failure.
type CustomError struct {
	Class       Class   `json:"class"`
	Fail        Failure `json:"failure"`
	Err         Error   `json:"error"`
	ExtendedErr *CustomError
}

// Error is method that allows *CustomError matches with error interface.
func (ce *CustomError) Error() string {
	if ce == nil {
		return "<nil> errror"
	}

	switch ce.Class {
	case ErrorClass:
		return ce.errorString()
	case FailureClass:
		return ce.failureString()
	}

	return "unexpected error class"
}

// CastStrict cases error into *CustomError. If input is not *CustomError it will returns <nil>.
func CastStrict(err error) *CustomError {
	if err == nil {
		return nil
	}

	var ce *CustomError
	if errors.As(err, &ce) {
		return ce
	}

	return nil
}

// Cast cases error into *CustomError. If input is not *CustomError it will creates new *CustomError.
func Cast(err error) *CustomError {
	if err == nil {
		return nil
	}

	if ce := CastStrict(err); ce != nil {
		return ce
	}

	return &CustomError{
		Class: ErrorClass,
		Err: Error{
			Message: err.Error(),
			Code:    DefaultErrorCode,
		},
	}
}

// Equals check if input values are equals
func Equals(left, right error) bool {
	if left == nil || right == nil {
		return errors.Is(left, right)
	}

	return left.Error() == right.Error()
}
