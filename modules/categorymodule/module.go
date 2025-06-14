package categorymodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	categoryhttpgin "github.com/katatrina/go12-service/modules/categorymodule/infras/controller/http-gin"
	categorygormmysql "github.com/katatrina/go12-service/modules/categorymodule/infras/repository/gorm-mysql"
	"github.com/katatrina/go12-service/modules/categorymodule/internal/service"
	"gorm.io/gorm"
)

var CategorySet = wire.NewSet(
	categorygormmysql.NewCategoryRepository,
	wire.Bind(new(categoryservice.ICategoryRepository), new(*categorygormmysql.CategoryRepository)),
	categoryservice.NewCategoryService,
	wire.Bind(new(categoryhttpgin.ICategoryService), new(*categoryservice.CategoryService)),
	categoryhttpgin.NewCategoryHTTPController,
	categoryservice.NewGetDetailQueryHandler,
	wire.Bind(new(categoryhttpgin.IDetailQueryHandler), new(*categoryservice.GetDetailQueryHandler)),
	wire.Bind(new(categoryservice.ICategoryQueryRepo), new(*categorygormmysql.CategoryRepository)),
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)
	
	catCtl.SetupRoutes(g)
}
