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
	"gorm.io/gorm"
)

func InitializeUserHTTPController(db *gorm.DB) *userhttpcontroller.UserHTTPController {
	userRepo := userrepository.NewUserRepository(db)
	
	jwtComp := sharedcomponent.NewJWTComp(os.Getenv("JWT_SECRET_KEY"), 3600*24*7*time.Second) // 7 days expiration
	
	// Command & Query Handlers
	registerCmdHandler := userservice.NewRegisterCommandHandler(userRepo)
	authenticateCmdHandler := userservice.NewAuthenticateCommandHandler(userRepo, jwtComp)
	
	userCtrl := userhttpcontroller.NewUserHTTPController(registerCmdHandler, authenticateCmdHandler)
	
	return userCtrl
}

func SetupUserModule(db *gorm.DB, router *gin.RouterGroup) {
	userCtrl := InitializeUserHTTPController(db)
	
	jwtComp := sharedcomponent.NewJWTComp(os.Getenv("JWT_SECRET_KEY"), 3600*24*7*time.Second) // 7 days expiration\
	
	userCtrl.SetupRoutes(router, middleware.AuthMiddleware(jwtComp))
}
