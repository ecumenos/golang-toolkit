package stringtools

// GetCharByIndex returns char from string by index with default value
func GetCharByIndex(in string, index int, def string) string {
	if len(in) <= index {
		return def
	}

	for i, char := range in {
		if i == index {
			return string(char)
		}
	}

	return def
}
