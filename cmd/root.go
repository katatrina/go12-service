package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	
	"github.com/katatrina/go12-service/gen/proto/category"
	"github.com/katatrina/go12-service/middleware"
	categorymodule "github.com/katatrina/go12-service/modules/category"
	categorygrpcctl "github.com/katatrina/go12-service/modules/category/infras/controller/grpcctl"
	categorygormmysql "github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	mediamodule "github.com/katatrina/go12-service/modules/media"
	restaurantmodule "github.com/katatrina/go12-service/modules/restaurant"
	restaurantlikemodule "github.com/katatrina/go12-service/modules/restaurantlike"
	usermodule "github.com/katatrina/go12-service/modules/user"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start rest api service",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		
		if port == "" {
			port = "8080"
		}
		
		dsn := os.Getenv("DB_DSN")
		dbMaster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
		if err != nil {
			log.Fatal("failed to connect database", err)
		}
		
		db := dbMaster.Debug()
		
		fmt.Println("Connected to database", db)
		
		r := gin.Default()
		
		r.Use(middleware.RecoverMiddleware())
		r.Use(otelgin.Middleware(serviceName))
		
		r.GET("/ping", func(c *gin.Context) {
			
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			
		})
		
		r.Static("/uploads", "./uploads")
		
		// CRUDL - Create Read Update Delete List
		// Version Prefix: /v1
		
		v1 := r.Group("/v1")
		
		appCtx := sharedinfras.NewAppContext(db)
		
		categorymodule.SetupCategoryModule(db, v1)
		restaurantmodule.SetupRestaurantModule(appCtx, v1)
		restaurantlikemodule.SetupRestaurantLikeModule(appCtx, v1)
		usermodule.SetupUserModule(appCtx, v1)
		mediamodule.SetupMediaModule(appCtx, v1)
		
		// Run gRPC server
		go func() {
			// Create a listener on TCP port
			lis, err := net.Listen("tcp", ":6000")
			if err != nil {
				log.Fatalln("Failed to listen:", err)
			}
			
			// Create a gRPC server object
			s := grpc.NewServer()
			// Attach the Greeter service to the server
			category.RegisterCategoryServer(s, categorygrpcctl.NewCategoryGrpcServer(categorygormmysql.NewCategoryRepository(db)))
			// Serve gRPC Server
			
			log.Println("Serving gRPC on 0.0.0.0:6000")
			log.Fatal(s.Serve(lis))
		}()
		
		// Init OTel Tracer
		
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
		
		r.Run(fmt.Sprintf(":%s", port))
	},
}

func Execute() {
	setupConsumerCmd()
	
	rootCmd.AddCommand(consumerCmd)
	rootCmd.AddCommand(testNatsCmd)
	
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to execute command", err)
	}
}

const serviceName = "go12-service"

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
	
	// sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.01))
	
	// Set global trace provider
	otel.SetTracerProvider(tp)
	
	// Set global propagator for distributed tracing context
	otel.SetTextMapPropagator(propagation.TraceContext{})
	
	log.Println("Tracer initialized successfully")
	return tp, nil
}
