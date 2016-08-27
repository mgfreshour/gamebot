
SOURCEDIR=.
BINARY=$(SOURCEDIR)/bin/gamebot
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
PACKAGES := $(shell go list ./... | grep -v /vendor/)

all: clean fmt test build

.DEFAULT_GOAL: all

deps:
	go get -u github.com/govend/govend
	govend -v
build:
	go build -o ${BINARY} bot.go
lint:
	gometalinter --exclude=vendor --exclude=_test --vendor --deadline 20s --disable=gocyclo --disable=gotype ./...
fmt:
	go fmt ${PACKAGES}
test:
	go test ${PACKAGES}
clean:
	rm -f ${BINARY}
