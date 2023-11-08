.PHONY: build
build:
	go mod vendor
	go build -v ./cmd/main

.PHONY: test
test:
	go test -v -timeout 30s -run ./cmd/main/ ...

.DEFAULT_GOAL := build
