
CGO_ENABLED=0

.PHONY: build
build: 
	go build -v ./cmd

.DEFAULT_GOAL := build