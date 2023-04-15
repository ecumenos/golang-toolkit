package numbertools

import "reflect"

// IsInt check if value is integer
func IsInt(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Int
}
