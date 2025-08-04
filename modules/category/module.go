package categorymodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpcontroller "github.com/katatrina/go12-service/modules/category/infras/controller/http"
	mysqlrepository "github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	categoryservice "github.com/katatrina/go12-service/modules/category/service"
	"gorm.io/gorm"
)

var CategorySet = wire.NewSet(
	mysqlrepository.NewCategoryRepository,
	
	httpcontroller.NewCategoryController,
	categoryservice.NewCreateCommandHandler,
	categoryservice.NewGetDetailQueryHandler,
	categoryservice.NewListCategoriesQueryHandler,
	categoryservice.NewUpdateByIDCommandHandler,
	categoryservice.NewDeleteByIDCommandHandler,
	
	wire.Bind(new(httpcontroller.ICreateCommandHandler), new(*categoryservice.CreateCommandHandler)),
	wire.Bind(new(httpcontroller.IGetByIDQueryHandler), new(*categoryservice.GetByIDQueryHandler)),
	wire.Bind(new(httpcontroller.IListQueryHandler), new(*categoryservice.ListCategoriesQueryHandler)),
	wire.Bind(new(httpcontroller.IUpdateByIDCommandHandler), new(*categoryservice.UpdateByIDCommandHandler)),
	wire.Bind(new(httpcontroller.IDeleteByIDCommandHandler), new(*categoryservice.DeleteByIDCommandHandler)),
	wire.Bind(new(httpcontroller.ICategoryRPCRepo), new(*mysqlrepository.CategoryRepository)),
	
	wire.Bind(new(categoryservice.ICreateRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(categoryservice.IGetByIDRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(categoryservice.IListRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(categoryservice.IUpdateByIDRepo), new(*mysqlrepository.CategoryRepository)),
	wire.Bind(new(categoryservice.IDeleteByIDRepo), new(*mysqlrepository.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	categoryGroup := g.Group("/categories")
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(categoryGroup)
	catCtl.SetupRoutesForRPC(g)
}

func InitializeCategoryController(db *gorm.DB) *httpcontroller.CategoryController {
	categoryRepository := mysqlrepository.NewCategoryRepository(db)
	createCommandHandler := categoryservice.NewCreateCommandHandler(categoryRepository)
	getByIDQueryHandler := categoryservice.NewGetDetailQueryHandler(categoryRepository)
	listCategoriesQueryHandler := categoryservice.NewListCategoriesQueryHandler(categoryRepository)
	updateByIDCommandHandler := categoryservice.NewUpdateByIDCommandHandler(categoryRepository)
	deleteByIDCommandHandler := categoryservice.NewDeleteByIDCommandHandler(categoryRepository)
	categoryController := httpcontroller.NewCategoryController(createCommandHandler, getByIDQueryHandler, listCategoriesQueryHandler, updateByIDCommandHandler, deleteByIDCommandHandler, categoryRepository)
	return categoryController
}
