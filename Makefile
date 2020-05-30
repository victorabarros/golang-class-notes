.DEFAULT_GOAL := help
.PHONY: help

APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /go/src/github.com/victorabarros/${APP_NAME}
PWD=$(shell pwd)

debug:
	@echo "\e[1m\033[32m\nDebug mode\e[0m"
	@docker rm -f go_notes
	docker run -it -v ${PWD}:${APP_DIR} -w ${APP_DIR} \
		-p 8091:8091 --name go_notes golang bash

linter:
	# https://github.com/golangci/golangci-lint
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run -v

fmt:
	docker run -v $(pwd):/go/src/github.com/victorabarros/fmt/ \
		-w /go/src/github.com/victorabarros/fmt/ golang:1.14 gofmt -e -l -s -w .

