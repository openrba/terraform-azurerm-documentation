include .env

PWD := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

DOCKER := docker run --rm
MOUNT := -v $(PWD):/data
MDBOOK := -it hrektts/mdbook mdbook
PORTS := -p 3000:3000 -p 3001:3001

.EXPORT_ALL_VARIABLES:

env:
	@echo $(GITHUB_TOKEN)
	@echo $(ORG)

go.mod:
	go mod init 'docs'

generate: go.mod
	go run .

build:
	@$(DOCKER) $(MOUNT) $(MDBOOK) $@

serve: build
	@$(DOCKER) $(PORTS) $(MOUNT) $(MDBOOK) $@ -p 3000 -n 0.0.0.0

.ONESHELL: