package main

import (
	"github.com/alextanhongpin/go-ddd/domain"
	"github.com/gin-gonic/gin"
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
func (ctl *UserController) GetUser(c *gin.Context) {
	// Call service.
}
