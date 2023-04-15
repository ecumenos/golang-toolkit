package slicetools

// SliceToAnySlice casts any slice to slice of any type
func SliceToAnySlice[T any](in []T) []any {
	out := make([]any, len(in))
	for i := range out {
		out[i] = in[i]
	}

	return out
}
