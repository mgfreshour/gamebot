

default: fmt test lint build

deps:
	go get -u github.com/govend/govend
	govend -v
build:
	go build bot.go
lint:
	gometalinter --deadline 10s --disable=gocyclo $(go list ./... | grep -v /vendor/)
fmt:
	go fmt $(go list ./... | grep -v /vendor/)
test:
	go test $(go list ./... | grep -v /vendor/)
