package stringtools

import (
	"strings"
)

// Contain check if string contains substring (ignore case)
func Contain(in, substr string) bool {
	return strings.Contains(strings.ToLower(in), strings.ToLower(substr))
}
