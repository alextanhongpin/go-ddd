// package server runs the main golang api server.

package main

import (
	"github.com/alextanhongpin/go-ddd/infra/server"
	"github.com/alextanhongpin/go-ddd/interface/controller"
	"github.com/alextanhongpin/go-ddd/interface/router"
)

func main() {
	cfg := NewConfig()
	r, stop := router.New(cfg.Env)
	defer stop()

	{
		ctl := controller.NewHealth(cfg.DeployedAt, cfg.Host, cfg.GitCommit)
		r.GET("/health", ctl.GetHealth)
	}

	server.New(cfg.Port, r)
}
