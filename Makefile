.PHONY: test
test:
	go test -race -cover ./...

.PHONY: install
install:
	go install ./...

.PHONY: lint
lint:
	go vet ./...

.PHONY: format
format:
	go fmt ./...
	gofmt -s -w ./