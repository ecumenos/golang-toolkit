package stringtools

// GetLen returns string length with default value if len is zero
func GetLen(in string, def int) int {
	if l := len(in); l > 0 {
		return l
	}

	return def
}
