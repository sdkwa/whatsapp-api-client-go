# Makefile for WhatsApp API Client Go

.PHONY: build test clean install examples lint fmt vet

# Build all examples
build: examples

# Run tests
test:
	go test -v -race -coverprofile=coverage.out ./...

# Test with coverage report
test-coverage: test
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install dependencies
install:
	go mod download
	go mod tidy

# Build all examples
examples:
	mkdir -p bin
	go build -o bin/send_message ./examples/send_message
	go build -o bin/get_qr ./examples/get_qr
	go build -o bin/webhook_server ./examples/webhook_server
	go build -o bin/create_group ./examples/create_group
	go build -o bin/instance_management ./examples/instance_management

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run all checks
check: fmt vet lint test

# Development setup
dev-setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run example - send message
run-send-message:
	go run ./examples/send_message

# Run example - get QR
run-get-qr:
	go run ./examples/get_qr

# Run example - webhook server
run-webhook-server:
	go run ./examples/webhook_server

# Run example - create group
run-create-group:
	go run ./examples/create_group

# Run example - instance management
run-instance-management:
	go run ./examples/instance_management
