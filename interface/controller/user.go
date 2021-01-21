package controller

import (
	"github.com/alextanhongpin/go-ddd/domain"
	"github.com/gin-gonic/gin"
)

// User endpoint.
type User struct {
	service domain.UserService
}

// NewUser factory.
func NewUser(svc domain.UserService) *User {
	return &User{
		service: svc,
	}
}

// GetUser route.
func (ctl *User) GetUser(c *gin.Context) {
	// Call service.
}
