package userhttpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	userservice "github.com/katatrina/go12-service/modules/user/service"
)

type IRegisterCommandHandler interface {
	Execute(ctx context.Context, cmd *userservice.RegisterCommand) (*usermodel.User, error)
}

type IAuthenticateCommandHandler interface {
	Execute(ctx context.Context, cmd *userservice.AuthenticateCommand) (*userservice.AuthenticateResponse, error)
}

type UserHTTPController struct {
	registerCmdHandler     IRegisterCommandHandler
	authenticateCmdHandler IAuthenticateCommandHandler
}

func NewUserHTTPController(
	registerCmdHandler IRegisterCommandHandler,
	authenticateCmdHandler IAuthenticateCommandHandler) *UserHTTPController {
	return &UserHTTPController{
		registerCmdHandler:     registerCmdHandler,
		authenticateCmdHandler: authenticateCmdHandler,
	}
}

func (h *UserHTTPController) SetupRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.POST("/register", h.Register)
	router.POST("/authenticate", h.Authenticate)
	
	router.GET("/profile", authMiddleware, h.GetProfile)
}
