package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alextanhongpin/pkg/requestid"
)

// RequestIDProvider obtains the request id from the X-Request-Id header if present, or
// creates a new one and populates the context with it.
func RequestIDProvider(provider requestid.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := provider(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				NewError(c, err),
			)
			return
		}
		c.Next()
	}
}
