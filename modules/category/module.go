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
	categoryservice.NewUpdateByIDCommandHandler,
	
	wire.Bind(new(categoryhttpgin.ICategoryService), new(*categoryservice.CategoryService)),
	wire.Bind(new(categoryhttpgin.ICreateCommandHandler), new(*categoryservice.CreateCommandHandler)),
	wire.Bind(new(categoryhttpgin.IGetDetailQueryHandler), new(*categoryservice.GetDetailQueryHandler)),
	wire.Bind(new(categoryhttpgin.IListQueryHandler), new(*categoryservice.ListCategoriesQueryHandler)),
	wire.Bind(new(categoryhttpgin.IUpdateByIDCommandHandler), new(*categoryservice.UpdateByIDCommandHandler)),
	
	wire.Bind(new(categoryservice.ICreateRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.IGetDetailRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.IListRepo), new(*categorygormmysql.CategoryRepository)),
	wire.Bind(new(categoryservice.IUpdateByIDRepo), new(*categorygormmysql.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(g)
}
