package mediamodule

import (
	mediahttpgin "github.com/katatrina/go12-service/modules/media/infras/controller/http-gin"
	gormmysql "github.com/katatrina/go12-service/modules/media/infras/repository/gorm-mysql"
	mediaservice "github.com/katatrina/go12-service/modules/media/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	
	"github.com/gin-gonic/gin"
)

func InitializeMediaController(appCtx sharedinfras.IAppContext) *mediahttpgin.MediaHTTPController {
	dbCtx := appCtx.DbContext()
	
	// Setup repositories and services
	mediaRepository := gormmysql.NewMediaRepository(dbCtx)
	
	// Set up command handlers
	createCommandHandler := mediaservice.NewCreateCommandHandler(mediaRepository)
	
	// Create HTTP controller
	mediaHTTPController := mediahttpgin.NewMediaHTTPController(createCommandHandler, appCtx.Uploader())
	return mediaHTTPController
}

func SetupMediaModule(appCtx sharedinfras.IAppContext, g *gin.RouterGroup) {
	mediaCtl := InitializeMediaController(appCtx)
	
	mediaCtl.SetupRoutes(g, appCtx.MiddlewareProvider().Auth())
}
