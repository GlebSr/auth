.PHONY: build
build:
	go build -v ./cmd/authApp
.PHONY: test
test:
	go test -v -race -timeout 30s ./...