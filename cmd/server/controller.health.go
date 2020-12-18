package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	DeployedAt time.Time
	Host       string
	GitCommit  string
}

func NewHealthController(cfg Config) *HealthController {
	return &HealthController{
		DeployedAt: cfg.DeployedAt,
		Host:       cfg.Host,
		GitCommit:  cfg.GitCommit,
	}
}

func (ctl *HealthController) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uptime":     time.Since(ctl.DeployedAt).String(),
		"host":       ctl.Host,
		"git_commit": ctl.GitCommit,
	})
}
