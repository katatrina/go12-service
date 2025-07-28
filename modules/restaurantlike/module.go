package restaurantlikemodule

import (
	restaurantlikehttpgin "github.com/katatrina/go12-service/modules/restaurantlike/infras/controller/http-gin"
	gormmysql "github.com/katatrina/go12-service/modules/restaurantlike/infras/repository/gorm-mysql"
	restaurantlikeservice "github.com/katatrina/go12-service/modules/restaurantlike/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	
	"github.com/gin-gonic/gin"
)

func InitializeRestaurantLikeController(appCtx sharedinfras.IAppContext) *restaurantlikehttpgin.RestaurantLikeHTTPController {
	dbCtx := appCtx.DbContext()
	
	restaurantLikeRepository := gormmysql.NewRestaurantLikeRepository(dbCtx)
	likeCommandHandler := restaurantlikeservice.NewLikeRestaurantCommandHandler(restaurantLikeRepository, appCtx.MsgBroker())
	unlikeCommandHandler := restaurantlikeservice.NewUnlikeRestaurantCommandHandler(restaurantLikeRepository)
	
	restaurantLikeHTTPController := restaurantlikehttpgin.NewRestaurantLikeHTTPController(
		likeCommandHandler,
		unlikeCommandHandler,
	)
	
	return restaurantLikeHTTPController
}

func SetupRestaurantLikeModule(appCtx sharedinfras.IAppContext, g *gin.RouterGroup) {
	restLikeCtl := InitializeRestaurantLikeController(appCtx)
	
	restLikeCtl.SetupRoutes(g, appCtx.MiddlewareProvider())
}
