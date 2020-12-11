ENV ?= development

-include .env
-include .env.$(ENV)
export

# Can be a space separated list
include Makefile.db Makefile.cli

start:
	@go run cmd/server/*.go
