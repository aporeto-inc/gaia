MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

PROJECT_SHA ?= $(shell git rev-parse HEAD)
PROJECT_VERSION ?= $(lastword $(shell git tag --sort version:refname --merged $(shell git rev-parse --abbrev-ref HEAD)))
PROJECT_RELEASE ?= dev

ci: init lint test

init:
	dep ensure -v

.PHONY:codegen
codegen:
	elegen folder -d specs -o codegen || exit 1
	mv custom_validations.go custom_validations.go.keep
	rm -rf ./*.go
	mv custom_validations.go.keep custom_validations.go
	mv codegen/elemental/*.go ./
	rm -rf codegen

lint:
	golangci-lint run \
		--disable-all \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=golint \
		--enable=structcheck \
		--enable=varcheck \
		--enable=ineffassign \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=unparam \
		./...

.PHONY: test
test:
	@for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
		if [ -f profile.out ]; then cat profile.out >> coverage.txt; rm -f profile.out; fi; \
	done;
