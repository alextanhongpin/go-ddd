// package server runs the main golang api server.

package main

func main() {
	cfg := NewConfig()
	r, stop := newRouter(cfg)
	defer stop()

	{
		ctl := NewHealthController(cfg)
		r.GET("/health", ctl.GetHealth)
	}

	newServer(cfg, r)
}
