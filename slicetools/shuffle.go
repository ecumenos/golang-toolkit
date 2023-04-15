package slicetools

import (
	"math/rand"
)

// Shuffle slice values order
func Shuffle[T any](in []T) []T {
	out := make([]T, len(in))
	for i, v := range rand.Perm(len(in)) {
		out[v] = in[i]
	}

	return out
}
