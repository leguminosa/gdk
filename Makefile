mockgen: # Generate mock files for all files with public interfaces.
	@./scripts/mockgen.sh

test: # Run unit tests for the whole repository.
	@go test -timeout 30s -short -count=1 -race -cover -coverprofile coverage.out -v ./...
	@go tool cover -func coverage.out