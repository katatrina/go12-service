.PHONY: setup run jaeger-up jaeger-down test

# Setup the project
setup:
	cd test_tracing && go mod tidy

# Run the application
run:
	cd test_tracing && go run main.go

# Start Jaeger
jaeger-up:
	cd test_tracing && docker-compose up -d

# Stop Jaeger
jaeger-down:
	cd test_tracing && docker-compose down

# Test the API endpoints
test:
	@echo "Testing ping endpoint..."
	@curl -s http://localhost:8080/ping
	@echo "\n"
	@echo "Testing user API endpoint..."
	@curl -s http://localhost:8080/user/123
	@echo "\n"

# Default target
all: setup jaeger-up run
