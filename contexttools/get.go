package contexttools

import "context"

// GetValue get value from context.Context by key
func GetValue(ctx context.Context, key string) string {
	if value, ok := ctx.Value(k(key)).(string); ok {
		return value
	}

	return ""
}
