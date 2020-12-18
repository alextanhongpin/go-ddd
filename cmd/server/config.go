package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int    `envconfig:"PORT" default:"8080"`
	Env  string `envconfig:"ENV" default:"development"`
}

func NewConfig() Config {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("parseConfigError: %v", err)
	}
	return cfg
}
