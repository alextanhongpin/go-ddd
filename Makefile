ENV ?= development

-include .env
-include .env.$(ENV)
export

# Can be a space separated list
include Makefile.db Makefile.cli

GIT_COMMIT := $(shell git rev-parse --short HEAD)

install:
	@go get github.com/google/wire/cmd/wie
	@go get -u -a golang.org/x/tools/cmd/stringer

start:
	@go run cmd/server/*.go
