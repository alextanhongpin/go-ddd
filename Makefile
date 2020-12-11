ENV ?= development

-include .env
-include .env.$(ENV)
export

# Can be a space separated list
include Makefile.db

start:
	@go run cmd/server/main.go
