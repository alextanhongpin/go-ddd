package main

import (
	"net/http"

	"github.com/alextanhongpin/go-ddd/domain"
)

// UserController endpoint.
type UserController struct {
	service domain.UserService
}

// NewUserController factory.
func NewUserController(svc domain.UserService) *UserController {
	return &UserController{
		service: svc,
	}
}

// GetUser route.
func (ctl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Call service.
}
