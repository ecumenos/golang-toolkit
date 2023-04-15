package slicetools

// Chunk split slice into slices n length
func Chunk[T interface{}](total int, preferred []T) [][]T {
	if len(preferred) == 0 {
		return [][]T{}
	}

	if total == 0 {
		return [][]T{}
	}

	divided := make([][]T, (len(preferred)+total-1)/total)
	prev := 0
	i := 0
	till := len(preferred) - total
	for prev < till {
		next := prev + total
		divided[i] = preferred[prev:next]
		prev = next
		i++
	}
	divided[i] = preferred[prev:]

	return divided
}

// ChunkFor split slice into n slices
func ChunkFor[T interface{}](items []T, itemsAmount int) [][]T {
	if len(items) < itemsAmount {
		return nil
	}

	var (
		out   = make([][]T, itemsAmount)
		index int
	)
	for _, item := range items {
		if index == itemsAmount {
			index = 0
		}
		out[index] = append(out[index], item)

		index++
	}

	return out
}
