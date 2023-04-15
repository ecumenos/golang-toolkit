package contexttools

import "context"

// SetValue set string value with key into context.Context
func SetValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, k(key), value)
}
