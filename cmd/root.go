package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	
	"github.com/katatrina/go12-service/gen/proto/category"
	"github.com/katatrina/go12-service/gen/proto/food"
	"github.com/katatrina/go12-service/middleware"
	categorymodule "github.com/katatrina/go12-service/modules/category"
	categorygrpcctl "github.com/katatrina/go12-service/modules/category/infras/controller/grpcctl"
	categorygormmysql "github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	foodmodule "github.com/katatrina/go12-service/modules/food"
	foodgrpcctl "github.com/katatrina/go12-service/modules/food/infras/controller/grpcctl"
	foodgormmysql "github.com/katatrina/go12-service/modules/food/infras/repository/mysql"
	mediamodule "github.com/katatrina/go12-service/modules/media"
	restaurantmodule "github.com/katatrina/go12-service/modules/restaurant"
	restaurantlikemodule "github.com/katatrina/go12-service/modules/restaurantlike"
	usermodule "github.com/katatrina/go12-service/modules/user"
	"github.com/katatrina/go12-service/shared/datatype"
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
		// Load configuration
		config := datatype.NewConfig()
		
		port := config.Port
		dsn := config.DBDSN
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
		foodmodule.SetupFoodModule(appCtx, v1)
		
		// Run Category gRPC server
		go func() {
			// Create a listener on TCP port
			lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Grpc.GetCategoryPort()))
			if err != nil {
				log.Fatalln("Failed to listen for Category gRPC:", err)
			}
			
			// Create a gRPC server object
			s := grpc.NewServer()
			// Attach the Category service to the server
			category.RegisterCategoryServer(s, categorygrpcctl.NewCategoryGrpcServer(categorygormmysql.NewCategoryRepository(db)))
			// Serve gRPC Server
			
			log.Printf("Serving Category gRPC on 0.0.0.0:%s", config.Grpc.GetCategoryPort())
			log.Fatal(s.Serve(lis))
		}()
		
		// Run Food gRPC server
		go func() {
			// Create a listener on TCP port for Food service
			lis2, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Grpc.GetFoodPort()))
			if err != nil {
				log.Fatalln("Failed to listen for Food gRPC:", err)
			}
			
			// Create a gRPC server object for Food service
			s2 := grpc.NewServer()
			// Attach the Food service to the server
			food.RegisterFoodServer(s2, foodgrpcctl.NewFoodGrpcServer(foodgormmysql.NewFoodRepository(sharedinfras.NewDbContext(db))))
			// Serve Food gRPC Server
			
			log.Printf("Serving Food gRPC on 0.0.0.0:%s", config.Grpc.GetFoodPort())
			log.Fatal(s2.Serve(lis2))
		}()
		
		// Start User and Restaurant gRPC servers
		startUserGRPCServer(config, db)
		startRestaurantGRPCServer(config, db)
		
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
