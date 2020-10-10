# Lint
.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: tes
test:
	go test ./...
