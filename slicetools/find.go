package slicetools

// FindOpts ...
type FindOpts[T any] struct {
	Elements []T
	Callback func(T) bool
	Strict   bool
}

// FindOne ...
func FindOne[T any](opts FindOpts[T]) T {
	for _, el := range opts.Elements {
		if opts.Callback(el) {
			return el
		}
	}
	var zero T

	return zero
}

// Find ...
func Find[T any](opts FindOpts[T]) []T {
	out := make([]T, 0)
	for _, el := range opts.Elements {
		if opts.Callback(el) {
			out = append(out, el)
		}
	}

	return out
}
