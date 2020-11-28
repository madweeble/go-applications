#!/bin/sh

PROJECTNAME=$(shell basename "$(PWD)")

BUILD_DIR := ./.build

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

clean: ## Clean build files
	@echo "  >  Cleaning build cache"
	@go clean
	@-rm -rf $(BUILD_DIR) ./test.out ./cover.out ./coverage.out ./coverage.html

build: ## Build the executable
	@echo "  >  Building ${PROJECTNAME}"
	@mkdir -p $(BUILD_DIR)
	@if [ ! -e go.mod ]; then go mod init ${PROJECTNAME}; fi
	@env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/authors cmd/author_service/Authors.go

test: ## Run all tests and output to test.out
	@echo "  >  Executing unit tests"
	@go test -v ./cmd/... > ./test.out

cover: ## Generate test coverage report coverage.html
	@echo "  >  Executing test coverage report"
	@go test -coverpkg=./... -coverprofile=cover.out ./...
	@go tool cover -html=cover.out -o coverage.html

run: build ## Run the executable
	@./build/authors $(ARGS)
