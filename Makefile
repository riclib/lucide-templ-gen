.PHONY: build install test clean example lint fmt

# Build the CLI tool
build:
	@echo "Building lucide-gen..."
	@go build -o bin/lucide-gen ./cmd/lucide-gen
	@echo "Built bin/lucide-gen"

# Install the CLI tool globally
install:
	@echo "Installing lucide-gen..."
	@go install ./cmd/lucide-gen
	@echo "Installed lucide-gen to GOPATH/bin"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@rm -rf examples/*/icons/
	@rm -rf examples/*/components/
	@rm -f examples/*/*.go
	@echo "Cleaned build artifacts"

# Run example
example:
	@echo "Running basic example..."
	@cd examples/basic && go run main.go

# Lint code
lint:
	@echo "Running linter..."
	@golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Generate documentation
docs:
	@echo "Generating documentation..."
	@godoc -http=:6060 &
	@echo "Documentation server started at http://localhost:6060"

# Release build (cross-platform)
release:
	@echo "Building release binaries..."
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build -o dist/lucide-gen-linux-amd64 ./cmd/lucide-gen
	@GOOS=darwin GOARCH=amd64 go build -o dist/lucide-gen-darwin-amd64 ./cmd/lucide-gen
	@GOOS=darwin GOARCH=arm64 go build -o dist/lucide-gen-darwin-arm64 ./cmd/lucide-gen
	@GOOS=windows GOARCH=amd64 go build -o dist/lucide-gen-windows-amd64.exe ./cmd/lucide-gen
	@echo "Release binaries built in dist/"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	@go mod tidy
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development environment ready"