.PHONY:  tools lint


# Tools required
TOOLS = github.com/golangci/golangci-lint/cmd/golangci-lint

# Lint
lint: tools
	PATH=$(PATH) golangci-lint run

# test
test:
	PATH=$(PATH) go test ./...