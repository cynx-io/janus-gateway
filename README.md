# Janus API Gateway

A lightweight API gateway for microservices, providing HTTP REST endpoints that communicate with gRPC services.

## Features

- HTTP REST API endpoints
- gRPC service communication
- JWT-based authentication
- CORS support
- Configurable service endpoints
- Clean and consistent response format

## Project Structure

```
janus/
├── api/
│   └── proto/
│       ├── src/           # Proto source files
│       │   └── hermes/    # Hermes service proto files
│       └── gen/           # Generated proto files
│           └── hermes/    # Generated Hermes service files
├── internal/
│   └── gateway/
│       ├── handlers/      # HTTP request handlers
│       └── middleware/    # Middleware components
├── config.json           # Application configuration
├── main.go              # Application entry point
├── Makefile            # Build and development commands
└── README.md           # This file
```

## Configuration

The application is configured through `config.json`:

```json
{
  "app": {
    "address": "0.0.0.0",
    "port": 8080,
    "name": "janus",
    "debug": true,
    "key": "your-app-key"
  },
  "grpc": {
    "hermes": "localhost:50051"
  },
  "jwt": {
    "secret": "your-jwt-secret",
    "expires_in": 24
  },
  "cors": {
    "enabled": true,
    "origins": ["http://localhost:3000"],
    "domain": "localhost"
  }
}
```

## Setup

1. Install dependencies:
   ```bash
   make deps
   ```

2. Generate proto files:
   ```bash
   make proto
   ```

3. Build the application:
   ```bash
   make build
   ```

4. Run the application:
   ```bash
   make run
   ```

## Available Commands

- `make deps` - Install required dependencies
- `make proto` - Generate proto files for all microservices
- `make proto-clean` - Clean generated proto files
- `make proto-gen` - Generate proto files without cleaning
- `make build` - Build the application
- `make run` - Run the application
- `make help` - Show available commands

## Adding a New Microservice

1. Create a new directory in `api/proto/src/` for your service
2. Add your proto files
3. Run `make proto` to generate the necessary Go files

Example:
```bash
mkdir -p api/proto/src/another-service
# Add your proto files
make proto
```

## Response Format

All API responses follow a consistent format:

```json
{
  "success": true,
  "data": {
    // Response data
  },
  "error": null
}
```

Error responses:
```json
{
  "success": false,
  "data": null,
  "error": {
    "code": "ERROR_CODE",
    "message": "Error description"
  }
}
```

## Development

The project uses several development tools:

- `golangci-lint` for code linting
- `fieldalignment` for struct field optimization
- `wire` for dependency injection
- `gorm` for database operations

Run `make tidy` to format code and run linters.

## License

MIT