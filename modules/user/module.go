package usermodule

import (
	"os"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/middleware"
	userhttpcontroller "github.com/katatrina/go12-service/modules/user/infras/controller/http"
	userrepository "github.com/katatrina/go12-service/modules/user/infras/repository/mysql"
	userservice "github.com/katatrina/go12-service/modules/user/service"
	sharedcomponent "github.com/katatrina/go12-service/shared/component"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

func InitializeUserHTTPController(appCtx sharedinfras.IAppContext) *userhttpcontroller.UserHTTPController {
	dbCtx := appCtx.DbContext()
	userRepo := userrepository.NewUserRepository(dbCtx)
	jwtComp := sharedcomponent.NewJWTComp(os.Getenv("JWT_SECRET_KEY"), 3600*24*7*time.Second) // 7 days expiration
	
	// Command & Query Handlers
	registerCmdHandler := userservice.NewRegisterCommandHandler(userRepo)
	authenticateCmdHandler := userservice.NewAuthenticateCommandHandler(userRepo, jwtComp)
	introspectCmdHandler := userservice.NewIntrospectCommandHandler(jwtComp, userRepo)
	
	userCtrl := userhttpcontroller.NewUserHTTPController(registerCmdHandler, authenticateCmdHandler, introspectCmdHandler)
	
	return userCtrl
}

func SetupUserModule(appCtx sharedinfras.IAppContext, g *gin.RouterGroup) {
	userCtrl := InitializeUserHTTPController(appCtx)
	
	jwtComp := sharedcomponent.NewJWTComp(os.Getenv("JWT_SECRET_KEY"), 3600*24*7*time.Second)
	
	userRepo := userrepository.NewUserRepository(appCtx.DbContext())
	introspectCmdHandler := userservice.NewIntrospectCommandHandler(jwtComp, userRepo)
	introspectCmdHdlWrapper := userservice.NewIntrospectCmdHdlWrapper(introspectCmdHandler)
	
	userCtrl.SetupRoutes(g, middleware.Auth(introspectCmdHdlWrapper))
	userCtrl.SetupAdminRoutes(g, middleware.Auth(introspectCmdHdlWrapper), middleware.CheckRoles(datatype.RoleAdmin))
}
