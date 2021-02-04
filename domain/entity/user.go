package entity

import "github.com/google/uuid"

// User entity.
type User struct {
	ID       uuid.UUID
	Email    string
	Name     string
	Password string
}
