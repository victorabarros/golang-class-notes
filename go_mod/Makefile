.DEFAULT_GOAL := help
.PHONY: help

APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /go/src/github.com/victorabarros/${APP_NAME}
IMAGE_NAME = ${APP_NAME}_image
PWD=$(shell pwd)

clean-up:
	@echo "\e[1m\033[32m\nRemoving go.mod and go.sum\e[0m"
	rm -rf go.*
	@echo "\e[1m\033[32m\nRemoving docker\e[0m"
	docker rmi -f ${IMAGE_NAME}


build-image:
	@echo "\e[1m\033[32m\nBuilding docker image\e[0m"
	docker build --rm . -t ${IMAGE_NAME}

run:
	@echo "\e[1m\033[32m\nStarting container\e[0m"
	docker run --rm --name ${APP_NAME} ${IMAGE_NAME}

debug:
	@echo "\e[1m\033[32m\nDebug mode\e[0m"
	docker run -it -v ${PWD}:${APP_DIR} -w ${APP_DIR} golang bash

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
