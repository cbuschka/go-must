PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	cd $(PROJECT_DIR) && \
	go build ./...
