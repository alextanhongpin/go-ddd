package domain

import "context"

type AuthService interface {
	Login(ctx context.Context) error
	Register(ctx context.Context) error
	ForgotPassword(ctx context.Context) error
	ResetPassword(ctx context.Context) error
	CreateToken(ctx context.Context) error
}
