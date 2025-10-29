# Makefile for Go Labs Artisan CLI

.PHONY: help build install artisan clean

# Default target
help:
	@echo "Available commands:"
	@echo "  make artisan     - Run artisan command (alias for ./artisan)"
	@echo "  make build       - Build the artisan binary"
	@echo "  make install     - Install artisan to /usr/local/bin"
	@echo "  make clean       - Clean build artifacts"

# Run artisan command
artisan:
	@./artisan $(ARGS)

# Build the artisan binary
build:
	@echo "Building artisan binary..."
	@go build -o bin/artisan cmd/artisan/main.go
	@echo "Binary built at bin/artisan"

# Install artisan to system PATH
install: build
	@echo "Installing artisan to /usr/local/bin..."
	@sudo cp bin/artisan /usr/local/bin/
	@echo "Artisan installed successfully!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@echo "Clean complete!"

# Example usage:
# make artisan ARGS="make:model User -m"
# make build
# make install
