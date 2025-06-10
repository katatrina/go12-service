package category

import (
	"github.com/gin-gonic/gin"
	categoryhttpgin "github.com/katatrina/go12-service/modules/category/infras/controller/http-gin"
	categorygormmysql "github.com/katatrina/go12-service/modules/category/infras/repository/gorm-mysql"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
	"gorm.io/gorm"
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catRepo := categorygormmysql.NewCategoryRepository(db)
	catService := categoryservice.NewCategoryService(catRepo)
	catCtl := categoryhttpgin.NewCategoryHTTPController(catService)
	
	catCtl.SetupRoutes(g)
}
