package maptools

// Equal compare primitive values
func Equal[K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64, V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](maps ...map[K]V) bool {
	if len(maps) == 0 {
		return true
	}

	for _, mm := range maps {
		for key, value := range mm {
			for _, m := range maps {
				v, ok := m[key]
				if !ok {
					return false
				}
				if v != value {
					return false
				}
			}
		}
	}

	return true
}
