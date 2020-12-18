package domain

import "context"

type AuthnService interface {
	Login(ctx context.Context, dto LoginDto) error
	Register(ctx context.Context, dto RegisterDto) error
	ForgotPassword(ctx context.Context) error
	ResetPassword(ctx context.Context) error
	ConfirmEmail(ctx context.Context) error
}
