# Go12 Service - Food Delivery Platform

A food delivery platform built with Go, featuring a modular monolith architecture with gRPC inter-service communication and event-driven updates.

## 🏗️ Architecture Overview

This service implements a **Hexagonal Architecture + Modular Monolith** pattern with:

- **Modular Design**: Each domain module is self-contained with clear boundaries
- **gRPC Communication**: Inter-service communication via gRPC (ports 6000-6003)
- **REST API**: HTTP endpoints for client applications (port 8080)
- **Event-Driven**: NATS message broker for asynchronous processing
- **Observability**: OpenTelemetry tracing with Jaeger integration

### Module Structure

```
modules/{domain}/
├── model/          # Domain entities, DTOs, errors
├── service/        # Business logic layer
└── infras/
    ├── controller/ # HTTP/gRPC handlers
    └── repository/ # Data access layer
```

## 🛠️ Technology Stack

- **Language**: Go 1.24.0
- **Web Framework**: Gin
- **Database**: MySQL with GORM ORM
- **Message Broker**: NATS
- **Protocol**: gRPC + Protocol Buffers
- **Authentication**: JWT
- **File Storage**: AWS S3
- **Tracing**: OpenTelemetry + Jaeger
- **Configuration**: Viper
- **Containerization**: Docker

## 🔧 Development Setup

### Prerequisites

- Go 1.24.0+
- Docker & Docker Compose
- MySQL 8.4.5+
- **Buf CLI** for Protocol Buffers: [Installation Guide](https://buf.build/docs/installation)

### Quick Start

1. **Clone and setup environment**:
   ```bash
   git clone <repository-url>
   cd go12-service
   cp .env.example .env
   # Edit .env with your configuration
   ```

2. **Start MySQL database**:
   ```bash
   make mysql-create
   make mysql-start
   ```

3. **Generate gRPC code** (required):
   ```bash
   # Generate Go code from .proto files
   buf generate
   ```

4. **Build and run**:
   ```bash
   go build -o app .
   ./app
   ```

   Or run directly:
   ```bash
   go run main.go
   ```

### Database Management

```bash
# Create MySQL container
make mysql-create

# Start/stop MySQL
make mysql-start
make mysql-stop

# Connect to MySQL
make mysql-connect

# View logs
make mysql-logs

# Clean up (removes container and data)
make mysql-clean
```

## ⚙️ Configuration

The service uses environment-based configuration with `.env` file support:

### Core Configuration
```env
# Database
DB_DSN=root:secret@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True

# Server
PORT=8080

# JWT
JWT_SECRET_KEY=your-jwt-secret-key
```

### gRPC Configuration
```env
# gRPC Service URLs
CATEGORY_GRPC_URL=localhost:6000
FOOD_GRPC_URL=localhost:6001
USER_GRPC_URL=localhost:6002
RESTAURANT_GRPC_URL=localhost:6003
```

### External Services
```env
# AWS S3
AWS_ACCESS_KEY=your-access-key
AWS_SECRET_KEY=your-secret-key
AWS_BUCKET_NAME=your-bucket
AWS_DOMAIN=https://your-domain.cloudfront.net
AWS_REGION=ap-southeast-1

# Message Broker
NATS_URL=nats://127.0.0.1:4222
```

## 🐳 Docker Deployment

```bash
# Build image
docker build -t go12-service:1.0.0 .

# Run container (see run_container.md for complete setup)
docker run -p 8080:8080 go12-service:1.0.0
```

## 🚀 Running the Application

The application will start:
- **REST API**: `http://localhost:8080`
- **gRPC Services**: ports 6000-6003
- **Health Check**: `GET /ping`

## 📋 Available Services

- **REST API** (port 8080): Client-facing HTTP endpoints
- **gRPC Services** (ports 6000-6003): Inter-service communication
- **Message Broker**: NATS for event processing
- **Observability**: Jaeger UI at `http://localhost:16686`

## 🛠️ Development Commands

```bash
# Generate gRPC code (required after cloning or changing .proto files)
buf generate

# Build and run
go build -o app .
./app

# Or run directly  
go run main.go
```

## ⚠️ Important Notes

- **gRPC Code Generation**: After cloning the repository, you MUST run `buf generate` before building
- **Protocol Buffer Files**: Located in `proto/` directory, generated Go files go to `gen/` (ignored by git)
- **When to Regenerate**: Run `buf generate` whenever you modify `.proto` files