package domain

import "context"

type Claims struct{}

type AuthzService interface {
	CreateToken(ctx context.Context) (string, error)
	VerifyToken(ctx context.Context, token string) (*Claims, error)
}
