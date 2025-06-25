package restaurantmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpcontroller "github.com/katatrina/go12-service/modules/restaurant/infras/controller/http"
	mysqlrepository "github.com/katatrina/go12-service/modules/restaurant/infras/repository/mysql"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	"gorm.io/gorm"
)

var RestaurantSet = wire.NewSet(
	mysqlrepository.NewRestaurantRepository,
	
	httpcontroller.NewRestaurantController,
	service.NewCreateCommandHandler,
	service.NewGetDetailQueryHandler,
	service.NewListRestaurantsQueryHandler,
	service.NewUpdateByIDCommandHandler,
	service.NewDeleteByIDCommandHandler,
	
	wire.Bind(new(httpcontroller.ICreateCommandHandler), new(*service.CreateCommandHandler)),
	wire.Bind(new(httpcontroller.IGetByIDQueryHandler), new(*service.GetByIDQueryHandler)),
	wire.Bind(new(httpcontroller.IListQueryHandler), new(*service.ListRestaurantsQueryHandler)),
	wire.Bind(new(httpcontroller.IUpdateByIDCommandHandler), new(*service.UpdateByIDCommandHandler)),
	wire.Bind(new(httpcontroller.IDeleteByIDCommandHandler), new(*service.DeleteByIDCommandHandler)),
	
	wire.Bind(new(service.ICreateRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(service.IGetByIDRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(service.IListRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(service.IUpdateByIDRepo), new(*mysqlrepository.RestaurantRepository)),
	wire.Bind(new(service.IDeleteByIDRepo), new(*mysqlrepository.RestaurantRepository)),
)

func SetupRestaurantModule(db *gorm.DB, g *gin.RouterGroup) {
	restCtl := InitializeRestaurantController(db)
	
	restCtl.SetupRoutes(g)
}
