package randomtools

import (
	"github.com/google/uuid"
)

// GetUUID generate uuid.GetUUID
func GetUUID() uuid.UUID {
	return uuid.New()
}

// GetUUIDString generate UUID string
func GetUUIDString() string {
	return uuid.NewString()
}
