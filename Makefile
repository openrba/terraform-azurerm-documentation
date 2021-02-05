include .env

PWD := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

MDBOOK := docker run --rm -v $(PWD):/data -it hrektts/mdbook mdbook
PORTS := -p 3000:3000 -p 3001:3001 -p 3000 -n 0.0.0.0

.EXPORT_ALL_VARIABLES:

env:
	@echo $(GITHUB_TOKEN)
	@echo $(ORG)

go.mod:
	go mod init 'docs'

test:
	go test

generate: go.mod
	go run .

build:
	@$(MDBOOK) $@

serve: build
	@$(MDBOOK) $@ $(PORTS)

help:
	@echo All Make Targets:
	@echo - test
	@echo - generate
	@echo - build
	@echo - serve

.PHONY: test generate build serve

.ONESHELL: