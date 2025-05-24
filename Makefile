# ==============================================================================
# Help

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==============================================================================
# Quality Checks

.PHONY: vet
## vet: runs Go vet to analyze code for potential issues
vet:
	@ echo "Running go vet..."
	@ go vet ./...

.PHONY: govulncheck
## govulncheck: runs Go vulnerability check
govulncheck:
	@ go install golang.org/x/vuln/cmd/govulncheck@latest
	@ echo "Running go vuln check..."
	@ govulncheck ./...

# ==============================================================================
# Tests

.PHONY: test
## test: run unit tests
test:
	@ go test -v ./... -count=1

.PHONY: coverage
## coverage: run unit tests and generate coverage report in html format
coverage:
	@ go test -coverprofile=coverage.out ./...  && go tool cover -html=coverage.out
