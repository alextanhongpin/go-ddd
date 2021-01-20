// package usercrud handles the User CRUD.

package usercrud

import (
	"database/sql"

	"github.com/alextanhongpin/go-ddd/domain"
)

// Any configuration for the usercrud lies here.

// UserCrud config.
type UserCrud struct {
	// Public variables.
	Repository domain.UserRepository
	Service    domain.UserService
}

func New(db *sql.DB) *UserCrud {
	repo := NewRepository(db)
	svc := NewService(repo)

	return &UserCrud{
		Repository: repo,
		Service:    svc,
	}
}
