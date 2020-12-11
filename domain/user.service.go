package domain

import "context"

// UserService interface.
type UserService interface {
	FindOne(ctx context.Context, id string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, dto CreateUserDto) (User, error)
	Update(ctx context.Context, dto UpdateUserDto) (User, error)
	Delete(ctx context.Context, dto DeleteUserDto) (User, error)
	// Putting all domains in the same folder helps prevent circular
	// dependency issue.
	// If we place comment and user entity/service in a separate folder, we can only refer to either one in either direction.
	UserComments(ctx context.Context, dto UserCommentsDto) ([]Comment, error)
}
