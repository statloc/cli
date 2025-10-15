.DEFAULT_GOAL := check

check: lint test build

lint:
	go tool golangci-lint run

build:
	go mod download
	go build -o bin/statloc cmd/main.go

test:
	go test ./...

cov:
	go test -coverprofile=.coverage ./...

clean:
	rm -f .coverage
	go clean -testcache
