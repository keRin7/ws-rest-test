
CGO_ENABLED=0

.PHONY: build
build: 
	go build -v ./cmd

.PHONY: build-linux
build-linux: 
	go build -o run -v ./cmd

.DEFAULT_GOAL := build