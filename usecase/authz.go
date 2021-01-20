package usecase

import (
	"context"

	"github.com/alextanhongpin/go-ddd/domain"
)

type Authz interface {
	Authorize(ctx context.Context, token string) (domain.User, error)
}
