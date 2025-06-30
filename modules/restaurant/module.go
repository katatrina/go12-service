package restaurantmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpcontroller "github.com/katatrina/go12-service/modules/restaurant/infras/controller/http"
	mysqlrepository "github.com/katatrina/go12-service/modules/restaurant/infras/repository/mysql"
	rpcclient "github.com/katatrina/go12-service/modules/restaurant/infras/repository/rpc"
	restaurantservice "github.com/katatrina/go12-service/modules/restaurant/service"
	"gorm.io/gorm"
)

var RestaurantSet = wire.NewSet(
	mysqlrepository.NewRestaurantRepository,
	
	httpcontroller.NewRestaurantController,
	restaurantservice.NewCreateCommandHandler,
	restaurantservice.NewGetDetailQueryHandler,
	restaurantservice.NewListRestaurantsQueryHandler,
	restaurantservice.NewUpdateByIDCommandHandler,
	restaurantservice.NewDeleteByIDCommandHandler,
	rpcclient.NewCategoryRPCClient,
	
	wire.Bind(new(httpcontroller.ICreateCommandHandler), new(*restaurantservice.CreateCommandHandler)),
	wire.Bind(new(httpcontroller.IGetByIDQueryHandler), new(*restaurantservice.GetByIDQueryHandler)),
	wire.Bind(new(httpcontroller.IListQueryHandler), new(*restaurantservice.ListRestaurantsQueryHandler)),
	wire.Bind(new(httpcontroller.IUpdateByIDCommandHandler), new(*restaurantservice.UpdateByIDCommandHandler)),
	wire.Bind(new(httpcontroller.IDeleteByIDCommandHandler), new(*restaurantservice.DeleteByIDCommandHandler)),
	
	wire.Bind(new(restaurantservice.ICreateRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(restaurantservice.IGetByIDRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(restaurantservice.IListRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(restaurantservice.IUpdateByIDRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(restaurantservice.IDeleteByIDRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(restaurantservice.ICategoryRepo), new(*rpcclient.CategoryRPCClient)),
)

func SetupRestaurantModule(db *gorm.DB, g *gin.RouterGroup) {
	restCtl := InitializeRestaurantController(db)
	
	restCtl.SetupRoutes(g)
}
