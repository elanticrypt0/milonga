# Milonga

A powerful, flexible, and zero-configuration Go framework for building modern backend and fullstack applications. Download and start coding right away!

## Features

- 🚀 **Zero Configuration**: Download and run with a simple `go run .` - start coding immediately!
- 🏗️ **Flexible Architecture**: MVC by default, but adaptable to any architecture pattern
- 🔒 **Built-in Authentication**: Includes "vigilante" module with JWT and OTP support
- 📦 **Database Integration**: GORM integration with multi-database configuration support
- ⚡ **High Performance**: Built on top of Fiber framework for excellent documentation and performance
- 🛠️ **Powerful CLI**: Generate configurations, migrations, seeds, and CRUD models
- 🐳 **Docker Ready**: Includes Docker Compose and hot-reload with Air configuration

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
go run .

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

## Project Vision

Milonga aims to be a simple yet powerful framework that allows developers to focus on building their applications rather than dealing with configuration. The goal is to provide a complete solution that works out of the box while remaining flexible enough to accommodate various project requirements.

The project is actively being developed with plans to include:
- Comprehensive documentation
- Web interface
- Additional API features
- Extended examples and use cases

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

¡Es argentino, papá! 🇦🇷