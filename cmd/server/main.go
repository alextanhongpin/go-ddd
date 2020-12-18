// package server runs the main golang api server.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init usecase.
	//uc := usercrud.New(nil)
	//ctl := NewUserController(uc.Service)
	//_ = ctl
	//u, err := uc.Service.FindOne(context.Background(), "1")
	//if err != nil {
	//log.Fatal(err)
	//}
	cfg := NewConfig()
	r, stop := newRouter(cfg)
	defer stop()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})
	newServer(cfg, r)
}
