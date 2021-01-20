package usercrud

import (
	"database/sql"

	"github.com/alextanhongpin/go-ddd/domain"
)

// Repository implementation.
type Repository struct {
	domain.UserRepository

	db *sql.DB
}

// NewRepository factory.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
