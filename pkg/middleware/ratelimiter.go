package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	ratelimit "github.com/alextanhongpin/pkg/ratelimiter"
)

// RateLimiter middleware.
func RateLimiter(limiter ratelimit.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		// For per path, consider concatenating the c.Request.URL.Path
		// with the client IP.
		if !limiter.Allow(fmt.Sprintf("%s/%s", c.Request.URL.Path, clientIP)) {
			err := errors.Errorf(`client ip "%s" has too many requests`, clientIP)
			c.Error(err)
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				NewError(c, err),
			)
			return
		}
		c.Next()
	}
}
