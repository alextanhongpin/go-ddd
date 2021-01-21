package router

import (
	"net/http"
	"time"

	"github.com/alextanhongpin/go-ddd/interface/controller"
	"github.com/alextanhongpin/go-ddd/pkg/middleware"
	"github.com/alextanhongpin/pkg/requestid"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func NewRouter(l *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// Custom middlewares.
	r.Use(middleware.Logger(l, time.RFC3339, true))
	r.Use(middleware.RequestIDProvider(
		requestid.RequestID(func() (string, error) {
			return xid.New().String(), nil
		}),
	))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})
	pprof.Register(r)
	return r
}

func New(l *zap.Logger) *gin.Engine {
	r := NewRouter(l)

	buildHealth(r)

	return r
}

func buildHealth(r *gin.Engine) {
	ctl := controller.NewHealth()
	r.GET("/health", ctl.GetHealth)
}
