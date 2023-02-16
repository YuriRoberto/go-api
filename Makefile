.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

default: help
run-all-tests: ## Run all tests without verbose
	go test ./...
	
run-all-tests-verbose: ## Run all tests verbosely
	go test -v ./...

get-deps: ## Get dependencies
	go mod tidy

run: get-deps ## Run api
	go run main.go