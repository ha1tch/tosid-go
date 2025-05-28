.PHONY: build test clean examples benchmark lint

# Build all binaries
build:
	@echo "Building examples..."
	@go build -o bin/basic-example cmd/examples/main.go
	@go build -o bin/disaster-response cmd/examples/disaster_response.go
	@go build -o bin/exoplanetary cmd/examples/exoplanetary.go
	@go build -o bin/kmac-disassembler cmd/examples/kmac_disassembler_example.go
	@go build -o bin/space-program cmd/examples/space_program.go
	@echo "Build complete!"

# Run all tests
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests complete!"

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
benchmark:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Run examples
examples: build
	@echo "Running basic example..."
	@./bin/basic-example
	@echo "\nRunning disaster response example..."
	@./bin/disaster-response
	@echo "\nRunning space program example..."
	@./bin/space-program

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete!"

# Lint code
lint:
	@echo "Running linters..."
	@golangci-lint run
	@echo "Linting complete!"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting complete!"

# Check for dependencies
deps:
	@echo "Checking dependencies..."
	@go mod tidy
	@go mod verify
	@echo "Dependencies verified!"

# Generate documentation
docs:
	@echo "Generating documentation..."
	@godoc -http=:6060 &
	@echo "Documentation server started at http://localhost:6060"

# Docker build
docker-build:
	@echo "Building Docker image..."
	@docker build -t semantic-computing:latest .
	@echo "Docker build complete!"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/tools/cmd/godoc@latest
	@echo "Tools installed!"

# Run all checks (test, lint, fmt)
check: fmt lint test
	@echo "All checks passed!"