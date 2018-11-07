MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

PROJECT_SHA ?= $(shell git rev-parse HEAD)
PROJECT_VERSION ?= $(lastword $(shell git tag --sort version:refname --merged $(shell git rev-parse --abbrev-ref HEAD)))
PROJECT_RELEASE ?= dev

ci: init lint test codecov

init:
	dep ensure
	dep status

.PHONY:codegen
codegen:
	elegen folder -d specs -o codegen || exit 1
	mv custom_validations.go custom_validations.go.keep
	mv custom_validations_test.go custom_validations_test.go.keep
	rm -rf ./*.go
	mv custom_validations.go.keep custom_validations.go
	mv custom_validations_test.go.keep custom_validations_test.go
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
	@ echo 'mode: atomic' > unit_coverage.cov
	@ for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
		if [ -f profile.out ]; then tail -q -n +2 profile.out >> unit_coverage.cov; rm -f profile.out; fi; \
	done;

coverage_aggregate:
	@ mkdir -p artifacts
	@ for f in `find . -maxdepth 1 -name '*.cov' -type f`; do \
		filename="$${f##*/}" && \
		go tool cover -html=$$f -o artifacts/$${filename%.*}.html; \
	done;

codecov: coverage_aggregate
	bash <(curl -s https://codecov.io/bash)
