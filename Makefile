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

build/app:
	go build -a -ldflags '-extldflags "-static"' -o app ./cmd/app

build/gw:
	go build -a -ldflags '-extldflags "-static"' -o gateway ./cmd/gateway

build: build/app build/gw

test:
	go test -race -covermode=atomic -coverprofile=coverage.tx -v ./...
	go tool cover -func=coverage.tx -o=coverage.out

test-html:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o=coverage.html

ci/compose-up:
	docker-compose -f ./deployments/local/ci.docker-compose.yml up -d --build

ci/compose-down:
	docker-compose -f ./deployments/local/ci.docker-compose.yml down

test-e2e:
	npx playwright test

ci/test-e2e:
	CI=true npx playwright test

ci: ci/compose-up ci/test-e2e ci/compose-down

all: generate tidy fmt vet lint test

