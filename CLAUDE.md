# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building and Running
- **Build**: `go build -o app .` (creates executable binary)
- **Run locally**: `go run main.go` (starts REST API on port 8080, gRPC services on ports 6000-6003)
- **Docker build**: `docker build -t go12-service:1.0.0 .`

### Database Management
- **Create MySQL container**: `make mysql-create`
- **Start MySQL**: `make mysql-start`
- **Stop MySQL**: `make mysql-stop`
- **Connect to MySQL**: `make mysql-connect`
- **Clean up MySQL**: `make mysql-clean`

### Protocol Buffers
- **Generate gRPC code**: `buf generate` (uses buf.gen.yaml configuration)
- **⚠️ IMPORTANT**: Must run `buf generate` after cloning or modifying .proto files
- **Generated files**: Located in `gen/` directory (ignored by Git)

## Architecture Overview

This is a microservice built with Go that provides a food delivery platform with the following core modules:

### Module Structure (Hexagonal Architecture + Modular Monolith)
Each module follows a consistent structure:
```
modules/{domain}/
├── model/          # Domain entities, DTOs, errors
├── service/        # Business logic layer
└── infras/
    ├── controller/ # HTTP/gRPC handlers
    └── repository/ # Data access layer
```

### Core Modules
- **Category**: CRUD operations for restaurant categories (HTTP + gRPC)
- **Restaurant**: Restaurant management with category relationships
- **RestaurantLike**: Like/unlike functionality for restaurants
- **User**: Authentication and user management
- **Media**: File upload to AWS S3
- **Food**: Food item management with enhanced responses (HTTP + gRPC)
- **User**: Authentication and user management with gRPC introspection
- **Restaurant**: Restaurant management with gRPC services

### Infrastructure Components
- **Database**: MySQL with GORM ORM
- **Message Broker**: NATS for event-driven architecture
- **File Storage**: AWS S3 for media uploads
- **Tracing**: OpenTelemetry with Jaeger
- **Authentication**: JWT-based auth with middleware
- **Protocol**: REST API (port 8080) + Full gRPC Architecture (ports 6000-6003)
- **Configuration**: Viper-based config management (.env file + environment variables)
- **Inter-Service Communication**: 100% gRPC (no HTTP calls between services)

### Key Files
- `cmd/root.go`: Main application setup, server configuration, tracing initialization
- `shared/infras/app_context.go`: Dependency injection container
- `shared/datatype/config.go`: Environment-based configuration
- `middleware/`: Authentication, recovery, role-based access control

### Environment Variables
Configuration can be set via `.env` file or environment variables:

#### Core Configuration
- `DB_DSN`: MySQL connection string
- `PORT`: HTTP server port (default: 8080)
- `JWT_SECRET_KEY`: JWT signing key

#### gRPC Configuration  
- `GRPC_PORT`: Category gRPC port (default: 6000)
- `FOOD_SERVICE_GRPC_PORT`: Food gRPC port (default: 6001)
- `USER_SERVICE_GRPC_PORT`: User gRPC port (default: 6002)
- `RESTAURANT_SERVICE_GRPC_PORT`: Restaurant gRPC port (default: 6003)
- `CATEGORY_SERVICE_GRPC_URL`: Category gRPC endpoint (default: localhost:6000)
- `FOOD_SERVICE_GRPC_URL`: Food gRPC endpoint (default: localhost:6001)
- `USER_SERVICE_GRPC_URL`: User gRPC endpoint (default: localhost:6002)
- `RESTAURANT_SERVICE_GRPC_URL`: Restaurant gRPC endpoint (default: localhost:6003)

#### Service URLs (HTTP fallback)
- `USER_SERVICE_URL`: User service HTTP endpoint
- `CATEGORY_SERVICE_URL`: Category service HTTP endpoint
- `FOOD_SERVICE_URL`: Food service HTTP endpoint

#### External Services
- `AWS_*`: S3 configuration (ACCESS_KEY, SECRET_KEY, BUCKET_NAME, DOMAIN, REGION)
- `NATS_URL`: Message broker connection (default: nats://127.0.0.1:4222)

#### Setup Instructions
1. Copy `.env.example` to `.env`: `cp .env.example .env`
2. Update `.env` with your actual configuration values
3. **Generate gRPC code**: `buf generate` (required after cloning)
4. The application will automatically load config from `.env` file with environment variable fallback

### Development Notes
- **Generated Code**: gRPC files in `gen/` are auto-generated from `.proto` files and ignored by Git
- **Code Generation Workflow**: Always run `buf generate` after modifying Protocol Buffer definitions
- The application uses dependency injection via the `appContext` pattern
- Each module is self-contained with its own repository, service, and controller layers
- **Pure gRPC Architecture**: 
  - All services provide gRPC endpoints for inter-service communication
  - Food service calls Category + Restaurant services via gRPC
  - User authentication via gRPC introspection (JWT token validation)
  - Zero HTTP calls between services - 100% gRPC communication
- **Enhanced Responses**: Food API returns enriched data with category and restaurant information
- **JWT Authentication**: Token introspection handled by User gRPC service
- NATS is used for event-driven updates (like count updates)
- OpenTelemetry tracing is configured for observability
- **True Microservices Pattern**: All modules communicate via gRPC only (no HTTP inter-service calls)
- **Dependency Removal**: Completely removed HTTP resty client library

### gRPC Services
- **Category Service**: `localhost:6000`
  - `GetCategoriesByIDs`: Batch category lookup
- **Food Service**: `localhost:6001`
  - `GetFoodsByIDs`: Batch food lookup
  - `GetFoodsByRestaurantID`: Get foods by restaurant
  - `GetFoodsByCategoryID`: Get foods by category
- **User Service**: `localhost:6002`
  - `IntrospectToken`: JWT token validation for authentication
  - `GetUsersByIDs`: Batch user lookup
- **Restaurant Service**: `localhost:6003`
  - `GetRestaurantsByIDs`: Batch restaurant lookup
  - `GetRestaurantsByCategoryID`: Get restaurants by category

### API Endpoints
- **Foods**: `GET /v1/foods/:id` returns enhanced response with category and restaurant info (via gRPC calls)
- **Categories**: Available via both HTTP (`/v1/categories`) and gRPC
- **Restaurants**: Available via both HTTP (`/v1/restaurants`) and gRPC
- **Users**: Authentication and profile management with gRPC token validation
- **Media**: File upload to S3
- **Authentication**: JWT middleware uses User gRPC service for token introspection

### Architecture Highlights
- **Zero HTTP Inter-Service Calls**: All internal communication via gRPC
- **Enhanced Data Loading**: Services aggregate data from multiple gRPC endpoints
- **Scalable Authentication**: JWT validation distributed via User gRPC service
- **Protocol Buffers**: Strongly-typed service contracts for all gRPC services
- **Clean Dependencies**: Removed all HTTP client libraries for inter-service communication