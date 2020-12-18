package main

import (
	"net/http"
	"time"

	"github.com/alextanhongpin/go-ddd/pkg/logger"
	"github.com/alextanhongpin/go-ddd/pkg/middleware"
	"github.com/alextanhongpin/pkg/requestid"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func newRouter(cfg Config) (*gin.Engine, func()) {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())

	l := logger.New(cfg.Env)

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
	return r, func() {
		l.Sync()
	}
}
