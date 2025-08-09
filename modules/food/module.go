package foodmodule

import (
	"github.com/gin-gonic/gin"
	httpcontroller "github.com/katatrina/go12-service/modules/food/infras/controller/http"
	grpcclient "github.com/katatrina/go12-service/modules/food/infras/repository/grpc-client"
	mysqlrepository "github.com/katatrina/go12-service/modules/food/infras/repository/mysql"
	foodservice "github.com/katatrina/go12-service/modules/food/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

func InitializeFoodController(appCtx sharedinfras.IAppContext) *httpcontroller.FoodController {
	dbCtx := appCtx.DbContext()
	config := appCtx.GetConfig()
	
	foodRepository := mysqlrepository.NewFoodRepository(dbCtx)
	
	// Initialize gRPC clients
	categoryRPC := grpcclient.NewCategoryRPCClient(config.Grpc.CategoryServiceURL)
	restaurantRPC := grpcclient.NewRestaurantGRPCAdapter(config.Grpc.RestaurantServiceURL)
	
	createCommandHandler := foodservice.NewCreateCommandHandler(foodRepository, categoryRPC, restaurantRPC)
	getByIDCommandHandler := foodservice.NewGetByIDCommandHandler(foodRepository, categoryRPC, restaurantRPC)
	listCommandHandler := foodservice.NewListCommandHandler(foodRepository)
	updateCommandHandler := foodservice.NewUpdateCommandHandler(foodRepository)
	deleteCommandHandler := foodservice.NewDeleteCommandHandler(foodRepository)
	
	foodController := httpcontroller.NewFoodController(
		createCommandHandler,
		getByIDCommandHandler,
		listCommandHandler,
		updateCommandHandler,
		deleteCommandHandler,
	)
	
	return foodController
}

func SetupFoodModule(appCtx sharedinfras.IAppContext, g *gin.RouterGroup) {
	foodCtl := InitializeFoodController(appCtx)
	
	foodCtl.SetupRoutes(g, appCtx.MiddlewareProvider())
}