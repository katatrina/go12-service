package categorymodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpcontroller "github.com/katatrina/go12-service/modules/category/infras/controller/http"
	mysqlrepository "github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	"github.com/katatrina/go12-service/modules/category/internal/service"
	"gorm.io/gorm"
)

var CategorySet = wire.NewSet(
	mysqlrepository.NewCategoryRepository,
	
	httpcontroller.NewCategoryController,
	service.NewCreateCommandHandler,
	service.NewGetDetailQueryHandler,
	service.NewListCategoriesQueryHandler,
	service.NewUpdateByIDCommandHandler,
	service.NewDeleteByIDCommandHandler,
	
	wire.Bind(new(httpcontroller.ICreateCommandHandler), new(*service.CreateCommandHandler)),
	wire.Bind(new(httpcontroller.IGetByIDQueryHandler), new(*service.GetByIDQueryHandler)),
	wire.Bind(new(httpcontroller.IListQueryHandler), new(*service.ListCategoriesQueryHandler)),
	wire.Bind(new(httpcontroller.IUpdateByIDCommandHandler), new(*service.UpdateByIDCommandHandler)),
	wire.Bind(new(httpcontroller.IDeleteByIDCommandHandler), new(*service.DeleteByIDCommandHandler)),
	
	wire.Bind(new(service.ICreateRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(service.IGetByIDRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(service.IListRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(service.IUpdateByIDRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(service.IDeleteByIDRepo), new(*mysqlrepository.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(g)
}
