envfile ?= .env
-include $(envfile)
ifneq ("$(wildcard $(envfile))","")
	export $(shell sed 's/=.*//' $(envfile))
endif

GOLANGCI_VERSION:=1.50.1
PROJECT_NAME:=flight-tracker
GOPATH_BIN:=$(shell go env GOPATH)/bin

.PHONY: init
init:
	@cp .env.dist .env
	@cp .env.test.dist .env.test

.PHONY: install
install:
	# Install protobuf compilation plugins.
	go install \
		github.com/swaggo/swag/cmd/swag@latest

	# Install golangci-lint for go code linting.
	curl -sSfL \
		"https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | \
		sh -s -- -b ${GOPATH_BIN} v${GOLANGCI_VERSION}

.PHONY: lint
lint:
	@echo ">>> Performing golang code linting.."
	golangci-lint run --config=.golangci.yml

.PHONY: build-server
build-server:
	@echo ">>> Building ${PROJECT_NAME} API server..."
	go build -o bin/server cmd/prometheus-api-server/main.go

.PHONY: run-server
run-server:
	@echo ">>> Running ${PROJECT_NAME} API server..."
	@go run ./cmd/${PROJECT_NAME}/main.go

.PHONY: doc
doc:
	@echo ">>> Generate Swagger API Documentation..."
	swag init --generalInfo cmd/${PROJECT_NAME}/main.go