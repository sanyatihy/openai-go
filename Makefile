.PHONY: test
test:
	go test -v ./...

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:lint
lint:
	golint ./...

.PHONY:vet
vet:
	go vet ./...
	shadow ./...

.PHONY:build-examples
build-examples:
	go build -o bin/ ./examples/*
