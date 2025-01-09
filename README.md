# Milonga

A powerful and flexible Go framework for building modern web applications.

## Features

- 🚀 High-performance web server using Fiber
- 🛠 Built-in CLI tool for code generation
- 📦 Docker support out of the box
- 🔄 Hot reload for development
- 🗄️ GORM integration for database operations
- 🔒 Built-in security features

## Quick Start

### Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose (optional)
- Bun (for web UI development)

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/milonga.git
cd milonga

# Install dependencies
go mod download
```

### Development

```bash
# Run the application
go run main.go

# Or with hot reload
air
```

### Using Docker

```bash
# Build the image
docker-compose build

# Start the service
docker-compose up

# Or in detached mode
docker-compose up -d

# View logs
docker-compose logs -f
```

## CLI Tool

Milonga comes with a powerful CLI tool for code generation.

### Generate CRUD Models

```bash
# Generate a new model with CRUD operations
go run main.go generate model User

# This will create:
# - api/models/user.go
# - api/handlers/user_handler.go
# - api/routes/user_routes.go
```

## Project Structure

```
.
├── api/
│   ├── handlers/    # Request handlers
│   ├── models/      # Data models
│   └── routes/      # Route definitions
├── cmd/
│   └── cli/         # CLI commands
├── config/          # Configuration files
├── public/          # Static files
└── docker-compose.yaml
```

## Configuration

Required folders for build:

- config
  - app_config.toml
  - db_config.toml
- public

## Web User Interface

To build the web UI:

> Note: The API must be running for the build to work

```bash
bun run build
```

The output will be placed in the `/public` directory.

## API Access

Default port is 8921 (configurable)

- API: [http://localhost:8921](http://localhost:8921)
- Public files: [http://localhost:8921/public](http://localhost:8921/public)
- HTMX example: [http://localhost:8921/public/examplex.html](http://localhost:8921/public/examplex.html)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.