.PHONY: build test run docker-build docker-run deploy

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Run the application locally
run:
	go run cmd/server/main.go

# Build Docker image
docker-build:
	docker build -t meeting-scheduler .

# Run Docker container
docker-run:
	docker run -p 8080:8080 meeting-scheduler

# Deploy to AWS ECS
deploy:
	cd deployments/terraform && \
	terraform init && \
	terraform apply -var-file=terraform.tfvars

# Clean up build artifacts
clean:
	rm -rf bin/
	go clean

# Install dependencies
deps:
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate OpenAPI documentation
docs:
	swag init -g cmd/server/main.go

# Run all checks
check: fmt lint test 