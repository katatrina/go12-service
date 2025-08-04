package restaurantmodule

import (
	"github.com/gin-gonic/gin"
	httpcontroller "github.com/katatrina/go12-service/modules/restaurant/infras/controller/http"
	rpcclient "github.com/katatrina/go12-service/modules/restaurant/infras/repository/grpc-client"
	mysqlrepository "github.com/katatrina/go12-service/modules/restaurant/infras/repository/mysql"
	restaurantservice "github.com/katatrina/go12-service/modules/restaurant/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

func InitializeRestaurantController(appCtx sharedinfras.IAppContext) *httpcontroller.RestaurantController {
	dbCtx := appCtx.DbContext()
	restaurantRepository := mysqlrepository.NewRestaurantRepository(dbCtx)
	createCommandHandler := restaurantservice.NewCreateCommandHandler(restaurantRepository)
	getByIDQueryHandler := restaurantservice.NewGetDetailQueryHandler(restaurantRepository)
	// categoryRPCClient := rpcclient.NewCategoryRPCClient(appCtx.GetConfig().CategoryServiceURL)
	categoryGRPCClient := rpcclient.NewCategoryRPCClient("0.0.0.0:6000")
	listRestaurantsQueryHandler := restaurantservice.NewListRestaurantsQueryHandler(restaurantRepository, categoryGRPCClient)
	updateByIDCommandHandler := restaurantservice.NewUpdateByIDCommandHandler(restaurantRepository)
	deleteByIDCommandHandler := restaurantservice.NewDeleteByIDCommandHandler(restaurantRepository)
	restaurantController := httpcontroller.NewRestaurantController(createCommandHandler, getByIDQueryHandler, listRestaurantsQueryHandler, updateByIDCommandHandler, deleteByIDCommandHandler)
	return restaurantController
}

func SetupRestaurantModule(appCtx sharedinfras.IAppContext, g *gin.RouterGroup) {
	restCtl := InitializeRestaurantController(appCtx)
	
	restCtl.SetupRoutes(g, appCtx.MiddlewareProvider())
}
