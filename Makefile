.PHONY: test
test:
	go test -race -cover ./...

.PHONY: install
install:
	go install ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: format
format:
	go fmt ./...
	gofmt -s -w ./