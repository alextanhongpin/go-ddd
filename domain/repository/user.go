package repository

import (
	"context"

	"github.com/alextanhongpin/go-ddd/domain/entity"
	"github.com/google/uuid"
)

// NOTE: Create or CreateUser? The latter, it makes composition easier. It's
// okay to stutter.

// User repository
type User interface {
	// Read.
	FindUser(ctx context.Context, id string) (entity.User, error)
	FindUsers(ctx context.Context) ([]entity.User, error)

	// Write.
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
