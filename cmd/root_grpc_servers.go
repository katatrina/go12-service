package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/katatrina/go12-service/gen/proto/restaurant"
	"github.com/katatrina/go12-service/gen/proto/user"
	restaurantgrpcctl "github.com/katatrina/go12-service/modules/restaurant/infras/controller/grpcctl"
	restaurantgormmysql "github.com/katatrina/go12-service/modules/restaurant/infras/repository/mysql"
	usergrpcctl "github.com/katatrina/go12-service/modules/user/infras/controller/grpcctl"
	usergormmysql "github.com/katatrina/go12-service/modules/user/infras/repository/mysql"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedcomponent "github.com/katatrina/go12-service/shared/component"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func startUserGRPCServer(config *datatype.Config, db *gorm.DB) {
	go func() {
		// Create a listener on TCP port for User service
		lis3, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Grpc.UserServicePort))
		if err != nil {
			log.Fatalln("Failed to listen for User gRPC:", err)
		}
		
		// Create a gRPC server object for User service
		s3 := grpc.NewServer()
		// Attach the User service to the server (need JWT component for introspection)
		jwtComp := sharedcomponent.NewJWTComp(config.JWTSecretKey)
		user.RegisterUserServer(s3, usergrpcctl.NewUserGrpcServer(usergormmysql.NewUserRepository(sharedinfras.NewDbContext(db)), jwtComp))
		// Serve User gRPC Server
		
		log.Printf("Serving User gRPC on 0.0.0.0:%s", config.Grpc.UserServicePort)
		log.Fatal(s3.Serve(lis3))
	}()
}

func startRestaurantGRPCServer(config *datatype.Config, db *gorm.DB) {
	go func() {
		// Create a listener on TCP port for Restaurant service
		lis4, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Grpc.RestaurantServicePort))
		if err != nil {
			log.Fatalln("Failed to listen for Restaurant gRPC:", err)
		}
		
		// Create a gRPC server object for Restaurant service
		s4 := grpc.NewServer()
		// Attach the Restaurant service to the server
		restaurant.RegisterRestaurantServer(s4, restaurantgrpcctl.NewRestaurantGrpcServer(restaurantgormmysql.NewRestaurantRepository(sharedinfras.NewDbContext(db))))
		// Serve Restaurant gRPC Server
		
		log.Printf("Serving Restaurant gRPC on 0.0.0.0:%s", config.Grpc.RestaurantServicePort)
		log.Fatal(s4.Serve(lis4))
	}()
}