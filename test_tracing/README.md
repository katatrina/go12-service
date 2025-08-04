# Simple GIN HTTP Server with OpenTelemetry and Jaeger

This is a simple demo of a GIN HTTP server with OpenTelemetry and Jaeger tracing.

## Prerequisites

- Go 1.20 or later
- Docker and Docker Compose

## Getting Started

1. Start Jaeger:

```bash
make jaeger-up
```

2. Install dependencies:

```bash
make setup
```

3. Run the server:

```bash
make run
```

4. Test the endpoints:

```bash
make test
```

5. View traces in Jaeger UI:

Open [http://localhost:16686](http://localhost:16686) in your browser.

## Endpoints

- `GET /ping` - Simple ping endpoint
- `GET /user/:id` - Get user by ID

## Shutdown

To stop Jaeger:

```bash
make jaeger-down
```
