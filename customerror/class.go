package customerror

// Class is type of error class.
type Class uint32

const (
	// FailureClass is failure error class constant.
	FailureClass Class = iota + 1
	// ErrorClass is error error class constant.
	ErrorClass
)
