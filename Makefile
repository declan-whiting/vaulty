.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...
test: 
	go test -v -failfast ./tests/
build: vet
	go build -o bin/vaulty ./cmd/vaulty
clean: 
	go clean ./...