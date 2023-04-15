package either

// Either is data struct that contain any value and error.
// it is useful for goroutines & channels to use single channel
// instead of using channel for value and channel for error
type Either[T any] struct {
	Value T
	Err   error
}
