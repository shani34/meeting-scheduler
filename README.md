# Meeting Scheduler API

A RESTful API service that helps geographically distributed teams find common meeting times that work for everyone.

## Features

- Create, update, and delete events
- Manage participant availability
- Find optimal meeting time slots based on participant availability
- Support for multiple time zones
- RESTful API design
- Automated tests
- Containerized deployment support
- Infrastructure as Code (IaC) support

## Project Structure

```
meeting-scheduler/
├── api/
│   ├── handlers/    # HTTP request handlers
│   ├── models/      # Data models and DTOs
│   └── services/    # Business logic
├── cmd/
│   └── server/      # Application entry point
├── internal/
│   ├── config/      # Configuration management
│   └── database/    # Database operations
├── tests/           # Integration and unit tests
└── deployments/     # Infrastructure as Code
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/shani34/meeting-scheduler.git
cd meeting-scheduler
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/server/main.go
```

### Running Tests

```bash
go test ./...
```

## API Documentation

The API documentation is available in OpenAPI/Swagger format at `/swagger/index.html` when the server is running.

## Deployment

The application can be deployed using Docker:

```bash
docker build -t meeting-scheduler .
docker run -p 8080:8080 meeting-scheduler
```

For production deployment, refer to the `deployments/` directory for Infrastructure as Code templates.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 