.PHONY: help
help: # List all available make commands.
	@printf "\033[33m%s\033[0m\n" "Usage: make [target]"
	@grep -E '^[A-Za-z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m%s\n", $$1, $$2}'

.PHONY: mockgen
mockgen: # Generate mock files for all files with public interfaces.
	@./scripts/mockgen.sh

.PHONY: test
test: # Run unit tests for the whole repository.
	@go test -timeout 30s -short -count=1 -race -cover -coverprofile coverage.out -v ./...
	@go tool cover -func coverage.out