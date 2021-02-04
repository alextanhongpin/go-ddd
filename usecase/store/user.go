package store

import (
	"context"

	"github.com/alextanhongpin/go-ddd/domain/entity"
	"github.com/google/uuid"
)

type UserRepository struct {
}

func (r *UserRepository) CreateUser(ctx context.Context, u *entity.User) error {
	u.ID = uuid.New()
	return nil
}
