package domain

import "context"

// Service is just a layer on top of repository in this case, just to validate the payload.
// This is because other usecases should only call services, and not repository
// directly for validation.
// We can divide service into two layers, basically utility services that
// performs business logic and/or validation and repository service layer, one
// that ensures payload to repository is valid.
// Otherwise, we can just stick to DTOs with validation for validating params
// for repository layer.

// UserService interface.
type UserService interface {
	// Read.
	FindOne(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindAll(ctx context.Context) ([]User, error)

	// Write.
	Create(ctx context.Context, dto CreateUserDto) (User, error)
	Update(ctx context.Context, dto UpdateUserDto) (User, error)
	Delete(ctx context.Context, dto DeleteUserDto) (User, error)
	// Putting all domains in the same folder helps prevent circular
	// dependency issue.
	// If we place comment and user entity/service in a separate folder,
	// we can only refer to either one in either direction.
	UserComments(ctx context.Context, dto UserCommentsDto) ([]Comment, error)

	// Helpers.
	ComparePassword(ctx context.Context, u User, password string) error
}
