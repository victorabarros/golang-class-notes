# environment-related variables
export GOPATH ?= $(shell go env GOPATH)
export GOLANGCILINT := $(GOPATH)/bin/golangci-lint
export GOLANGCILINT_VERSION := v1.50.1
export PROTOC_VERSION ?= 21.8
export PROTOC_GEN_GO ?= v1.28
export PROTOC_GEN_GO_GRPC ?= v1.2
PROTOC := $(CURDIR)/protoc/bin/protoc

# app-related variables
APP_NAME ?= articles-api
VERSION ?= dev

build: ## Build executales
	go build -mod=vendor -ldflags "-X main.Version=$(VERSION)" -o $(APP_NAME) ./cmd/articles

clean:  ## Clean
	rm -fr $(APP_NAME) tmp protoc &> /dev/null

fmt:  ## Execute go fmt
	go fmt ./...

lint: $(GOLANGCILINT)  ## Run static checks
	$(CURDIR)/scripts/lint.sh

protoc: $(PROTOC) ## Protobuf compile
	$(PROTOC) --proto_path=internal/transport/proto --go_out=. --go-grpc_out=. internal/transport/proto/*.proto

test: ## Run tests
	go test -count=1 ./... -v

vendor: ## Fetch dependencies
	go mod tidy
	go get ./...
	go mod verify
	go mod vendor

reinstall-golangci-lint: ## Remove any existing versions of golangci-lint and install it again
	$(CURDIR)/scripts/reinstall-golangci-lint.sh

run: ## Run executable
	./$(APP_NAME)

$(GOLANGCILINT):
	$(CURDIR)/scripts/install-golangci-lint.sh

$(PROTOC):
	$(CURDIR)/scripts/install-protoc.sh

help: ## Display help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.MAIN: clean build run
.PHONY: build clean lint protoc run test vendor
