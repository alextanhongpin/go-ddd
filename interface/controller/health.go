package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Health struct {
	startAt   time.Time
	host      string
	gitCommit string
}

func NewHealth() *Health {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &Health{
		startAt:   time.Now(),
		host:      host,
		gitCommit: os.Getenv("GIT_COMMIT"),
	}
}

func (ctl *Health) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uptime":     time.Since(ctl.startAt).String(),
		"host":       ctl.host,
		"git_commit": ctl.gitCommit,
	})
}
