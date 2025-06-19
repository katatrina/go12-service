package categorymodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/katatrina/go12-service/modules/category/infras/controller/controller"
	"github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	"github.com/katatrina/go12-service/modules/category/internal/service"
	"gorm.io/gorm"
)

var CategorySet = wire.NewSet(
	repository.NewCategoryRepository,
	
	controller.NewCategoryHTTPController,
	service.NewCreateCommandHandler,
	service.NewGetDetailQueryHandler,
	service.NewListCategoriesQueryHandler,
	service.NewUpdateByIDCommandHandler,
	service.NewDeleteByIDCommandHandler,
	
	wire.Bind(new(controller.ICreateCommandHandler), new(*service.CreateCommandHandler)),
	wire.Bind(new(controller.IGetByIDQueryHandler), new(*service.GetByIDQueryHandler)),
	wire.Bind(new(controller.IListQueryHandler), new(*service.ListCategoriesQueryHandler)),
	wire.Bind(new(controller.IUpdateByIDCommandHandler), new(*service.UpdateByIDCommandHandler)),
	wire.Bind(new(controller.IDeleteByIDCommandHandler), new(*service.DeleteByIDCommandHandler)),
	
	wire.Bind(new(service.ICreateRepo), new(*repository.CategoryRepository)),
	wire.Bind(new(service.IGetByIDRepo), new(*repository.CategoryRepository)),
	wire.Bind(new(service.IListRepo), new(*repository.CategoryRepository)),
	wire.Bind(new(service.IUpdateByIDRepo), new(*repository.CategoryRepository)),
	wire.Bind(new(service.IDeleteByIDRepo), new(*repository.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(g)
}
