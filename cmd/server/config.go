package main

import (
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int       `envconfig:"PORT" default:"8080"`
	Env        string    `envconfig:"ENV" default:"development"`
	GitCommit  string    `envconfig:"GIT_COMMIT" required:"true"`
	DeployedAt time.Time `envconfig:"-"`
	Host       string    `envconfig:"-"`
}

func NewConfig() Config {
	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("hostnameError: %v", err)
	}

	cfg := Config{
		DeployedAt: time.Now(),
		Host:       host,
	}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("parseConfigError: %v", err)
	}
	return cfg
}
