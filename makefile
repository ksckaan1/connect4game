run-api:
	@go run ./cmd/api

run-cli:
	@go run ./cmd/cli

test:
	@go clean -testcache
	@go test ./...