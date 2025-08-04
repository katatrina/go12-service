package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	serviceName = "gin-demo-api"
)

// initTracer initializes the OpenTelemetry tracer with Jaeger HTTP collector
func initTracer() (*sdktrace.TracerProvider, error) {
	log.Println("Initializing tracer with Jaeger HTTP collector")

	// Create Jaeger exporter with HTTP collector endpoint
	exporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint("http://localhost:14268/api/traces"),
		),
	)

	if err != nil {
		log.Printf("Failed to create Jaeger exporter: %v", err)
		return nil, err
	}
	log.Println("Jaeger exporter created successfully")

	// Create resource with service information using semantic conventions
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String("development"),
		),
	)
	if err != nil {
		log.Printf("Failed to create resource: %v", err)
		return nil, err
	}
	log.Println("Resource created successfully")

	// Create trace provider with the exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	// Set global trace provider
	otel.SetTracerProvider(tp)

	// Set global propagator for distributed tracing context
	otel.SetTextMapPropagator(propagation.TraceContext{})

	log.Println("Tracer initialized successfully")
	return tp, nil
}

// handlePing handles the /ping endpoint
func handlePing(c *gin.Context) {
	// Get the current span from context
	span := trace.SpanFromContext(c.Request.Context())
	span.AddEvent("Handling ping request")

	// Create a child span for processing
	ctx, childSpan := otel.Tracer(serviceName).Start(c.Request.Context(), "process-ping")
	defer childSpan.End()

	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	childSpan.AddEvent("Ping processing completed")

	// Create another child span for database simulation
	_, dbSpan := otel.Tracer(serviceName).Start(ctx, "database-query")
	// Simulate database query
	time.Sleep(50 * time.Millisecond)
	dbSpan.AddEvent("Database query completed")
	dbSpan.End()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// handleUser handles the /user/:id endpoint
func handleUser(c *gin.Context) {
	// Get the current span from context
	span := trace.SpanFromContext(c.Request.Context())
	userID := c.Param("id")
	span.AddEvent("Handling user request for ID: " + userID)

	// Create a child span for user processing
	ctx, childSpan := otel.Tracer(serviceName).Start(c.Request.Context(), "process-user")
	defer childSpan.End()

	// Simulate some work
	time.Sleep(150 * time.Millisecond)
	childSpan.AddEvent("User processing completed")

	// Create another child span for database simulation
	_, dbSpan := otel.Tracer(serviceName).Start(ctx, "database-query")
	// Simulate database query
	time.Sleep(75 * time.Millisecond)
	dbSpan.AddEvent("Database query completed")
	dbSpan.End()

	c.JSON(http.StatusOK, gin.H{
		"id":    userID,
		"name":  "User " + userID,
		"email": "user" + userID + "@example.com",
		"time":  time.Now().Format(time.RFC3339),
	})
}

func main() {
	// Initialize the tracer
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Create a gin router
	r := gin.Default()

	// Add OpenTelemetry middleware
	r.Use(otelgin.Middleware(serviceName))

	// Define routes
	r.GET("/ping", handlePing)
	r.GET("/user/:id", handleUser)

	// Start the server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
