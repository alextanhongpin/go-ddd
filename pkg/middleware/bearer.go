package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/alextanhongpin/pkg/gojwt"
)

const (
	// Capitalization matters.
	AuthorizationBearer = "Bearer"
	AuthorizationBasic  = "Basic"

	UserContext = contextKey("user")
)

func BearerAuthorizer(sign gojwt.Signer) gin.HandlerFunc {
	checkAuthorization := func(auth string) (*gojwt.Claims, error) {
		paths := strings.Split(auth, " ")
		if len(paths) != 2 {
			return nil, errors.New("missing authorization header")
		}
		bearer, token := paths[0], paths[1]
		if valid := strings.EqualFold(bearer, AuthorizationBearer); !valid {
			return nil, errors.New("invalid authorization header")
		}
		claims, err := sign.Verify(token)
		if err != nil {
			return nil, fmt.Errorf("middleware verify token failed: %w", err)
		}
		return claims, nil
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, err := checkAuthorization(c.GetHeader("Authorization"))
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(c, err),
			)
			return
		}

		var (
			user = claims.StandardClaims.Subject
		)

		// Set the context for the next request.
		ctx = UserContext.WithValue(ctx, user)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
