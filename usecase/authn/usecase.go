// package authn handles the authentication usecase/flow.

package authn

import (
	"context"

	"github.com/alextanhongpin/go-ddd/domain"
	"github.com/alextanhongpin/go-ddd/usecase"
)

// UseCase struct.
type UseCase struct {
	usecase.Authn

	userService      domain.UserService
	userTokenService domain.UserTokenService
}

// Login verifies the user identity and returns an access token.
func (u *UseCase) Login(ctx context.Context, dto usecase.LoginDto) (string, error) {
	usr, err := u.userService.FindByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	if err := u.userService.ComparePassword(ctx, usr, dto.Password); err != nil {
		return "", err
	}

	token, err := u.userTokenService.Sign(usr)
	if err != nil {
		return "", err
	}
	return token, nil
}
