package categorymodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	categoryhttpgin "github.com/katatrina/go12-service/modules/category/infras/controller/http-gin"
	categorygormmysql "github.com/katatrina/go12-service/modules/category/infras/repository/gorm-mysql"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
	"gorm.io/gorm"
)

var CategorySet = wire.NewSet(
	categorygormmysql.NewCategoryRepository,
	
	categoryservice.NewCategoryService,
	
	categoryhttpgin.NewCategoryHTTPController,
	categoryservice.NewCreateCommandHandler,
	categoryservice.NewGetDetailQueryHandler,
	categoryservice.NewListCategoriesQueryHandler,
	
	wire.Bind(new(categoryhttpgin.ICategoryService), new(*categoryservice.CategoryService)),
	wire.Bind(new(categoryhttpgin.ICreateCommandHandler), new(*categoryservice.CreateCommandHandler)),
	wire.Bind(new(categoryhttpgin.IGetDetailQueryHandler), new(*categoryservice.GetDetailQueryHandler)),
	wire.Bind(new(categoryhttpgin.IListCategoriesQueryHandler), new(*categoryservice.ListCategoriesQueryHandler)),
	
	wire.Bind(new(categoryservice.ICreateRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.IGetDetailRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.ICategoryCommandRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.IListCategoriesRepo), new(*categorygormmysql.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(g)
}
