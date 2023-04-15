package stringtools

import (
	"reflect"
)

// IsString check if value is string
func IsString(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.String
}
