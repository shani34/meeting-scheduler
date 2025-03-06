.PHONY: build run test migrate clean

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Run database migrations
migrate:
	psql -U postgres -d meeting_scheduler -f migrations/001_initial_schema.sql

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Install dependencies
deps:
	go mod tidy

# Run linter
lint:
	golangci-lint run

# Generate API documentation
docs:
	swag init -g cmd/server/main.go 