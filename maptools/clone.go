package maptools

// Clone map with primitive values
func Clone[K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64, V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m map[K]V) map[K]V {
	clonedMap := make(map[K]V)
	for key, value := range m {
		clonedMap[key] = value
	}

	return clonedMap
}
