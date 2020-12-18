ENV ?= development

-include .env
-include .env.$(ENV)
export

# Can be a space separated list
include Makefile.db Makefile.cli

GIT_COMMIT := $(shell git rev-parse --short HEAD)

start:
	@go run cmd/server/*.go
