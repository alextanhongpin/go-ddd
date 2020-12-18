package domain

import "context"

type AuthnService interface {
	Login(ctx context.Context) error
	Register(ctx context.Context) error
	ForgotPassword(ctx context.Context) error
	ResetPassword(ctx context.Context) error
	ConfirmEmail(ctx context.Context) error
}
