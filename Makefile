.PHONY: install
install:
	go install ./...

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: format
format:
	go fmt ./...
	gofmt -s -w ./