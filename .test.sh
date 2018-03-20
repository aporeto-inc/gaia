#!/bin/bash
set -e

# --enable golint is disabled because Dimitri likes underscores...
gometalinter \
    --exclude bindata.go \
    --exclude vendor \
    --vendor \
    --disable-all \
    --enable vet \
    --enable vetshadow \
    --enable ineffassign \
    --enable goconst \
    --enable errcheck \
    --enable varcheck \
    --enable structcheck \
    --enable gosimple \
    --enable misspell \
    --enable deadcode \
    --enable staticcheck \
    --enable golint \
    --deadline 5m \
    --tests ./...

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then cat profile.out >> coverage.txt; rm -f profile.out; fi
done
