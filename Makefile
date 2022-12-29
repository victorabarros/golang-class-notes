.DEFAULT_GOAL := help
.PHONY: help

APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /go/src/github.com/victorabarros/${APP_NAME}
PWD=$(shell pwd)

clean-up:
	@docker rm -f ${APP_NAME}

debug:
	@echo "Debug mode"
	# @docker rm -f go_notes
	docker run -it -v ${PWD}:${APP_DIR} -w ${APP_DIR} \
		-p 8091:8091 --rm --name ${APP_NAME} golang bash

linter:
	# https://github.com/golangci/golangci-lint
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run -v

fmt:
	docker run -v $(pwd):/go/src/github.com/victorabarros/fmt/ \
		-w /go/src/github.com/victorabarros/fmt/ golang:1.14 gofmt -e -l -s -w .

