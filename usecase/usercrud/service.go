package usercrud

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-ddd/domain"
)

// Service implementation.
type Service struct {
	domain.UserService
	repo domain.UserRepository
}

// NewService factory.
func NewService(repo domain.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create a new User.
func (u *Service) Create(ctx context.Context, dto domain.CreateUserDto) (domain.User, error) {
	var usr domain.User
	// Do validation.
	// Read/save data to/from repository.
	return usr, nil
}

// FindOne returns the User with the given id.
func (u *Service) FindOne(ctx context.Context, id string) (domain.User, error) {
	var usr domain.User
	fmt.Println("finding user with id", id)
	return usr, nil
}
