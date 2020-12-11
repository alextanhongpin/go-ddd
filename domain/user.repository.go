package domain

import "context"

// UserRepository interface.
type UserRepository interface {
	FindOne(ctx context.Context, id string) (User, error)
}
