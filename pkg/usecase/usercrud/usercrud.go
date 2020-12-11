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
	Controller *Controller // The last layer is usually not an interface, but an implementation.
}

func New(db *sql.DB) *UserCrud {
	repo := NewRepository(db)
	svc := NewService(repo)
	ctl := NewController(svc)

	return &UserCrud{
		Repository: repo,
		Service:    svc,
		Controller: ctl,
	}
}
