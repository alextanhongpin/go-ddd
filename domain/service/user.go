package service

import (
	"context"

	"github.com/alextanhongpin/go-ddd/domain/entity"
	"github.com/alextanhongpin/go-ddd/domain/repository"
)

type UserService struct {
	repo repository.User
}

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
}

func (u *UserService) CreateUser(ctx context.Context, dto CreateUserDto) (entity.User, error) {
	var usr entity.User
	usr.Name = dto.Name
	usr.Email = dto.Email
	usr.Password = dto.Password
	if err := u.repo.CreateUser(ctx, &usr); err != nil {
		return usr, err
	}
	return usr, nil
}
