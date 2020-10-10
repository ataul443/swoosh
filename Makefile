# Lint
.PHONY: lint
lint:
	golangci-lint run ./...
