.PHONY: test bench lint fmt vet coverage clean help

# Default target
help:
	@echo "GoTree Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  make test      - Run all tests"
	@echo "  make bench     - Run benchmarks"
	@echo "  make lint      - Run linter"
	@echo "  make fmt       - Format code"
	@echo "  make vet       - Run go vet"
	@echo "  make coverage  - Generate coverage report"
	@echo "  make clean     - Clean build artifacts"

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem -run=^$$

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean:
	rm -f coverage.out coverage.html
	go clean -testcache
