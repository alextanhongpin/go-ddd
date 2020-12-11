package usercrud

import (
	"net/http"

	"github.com/alextanhongpin/go-ddd/domain"
)

// Controller endpoint.
type Controller struct {
	service domain.UserService
}

// NewController factory.
func NewController(svc domain.UserService) *Controller {
	return &Controller{
		service: svc,
	}
}

// GetUser route.
func (ctl *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	// Call service.
}
