.PHONY: build

tidy:
	go mod tidy

generate:
	buf generate

fmt:
	go fmt ./...

vet:
	go vet -v ./...

lint:
	buf lint
	golangci-lint -v run

build:
	go build

test:
	go test -race -covermode=atomic -coverprofile=coverage.tx -v ./...
	go tool cover -func=coverage.tx -o=coverage.out

test-html:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o=coverage.html

all: generate tidy fmt vet lint test