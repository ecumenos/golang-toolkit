package maptools

// Merge maps with primitive values
func Merge[K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64, V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](args ...map[K]V) map[K]V {
	outSum := make(map[K]V)
	outCount := make(map[K]V)
	for _, arg := range args {
		for key, value := range arg {
			val, ok := outSum[key]
			if ok {
				outSum[key] = val + value
				outCount[key]++
			} else {
				outSum[key] = value
				outCount[key] = 1
			}
		}
	}
	out := make(map[K]V, len(outSum))
	for key, sum := range outSum {
		out[key] = sum / outCount[key]
	}

	return out
}
