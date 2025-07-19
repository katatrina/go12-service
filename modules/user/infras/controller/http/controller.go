package userhttpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	usermodel "github.com/katatrina/go12-service/modules/user/model"
	userservice "github.com/katatrina/go12-service/modules/user/service"
)

type IRegisterCommandHandler interface {
	Execute(ctx context.Context, cmd *userservice.RegisterCommand) (*usermodel.User, error)
}

type IAuthenticateCommandHandler interface {
	Execute(ctx context.Context, cmd *userservice.AuthenticateCommand) (*userservice.AuthenticateResponse, error)
}

type IQueryUserRepository interface {
	FindByID(ctx context.Context, userID uuid.UUID) (*usermodel.User, error)
}

type IIntrospectCommandHandler interface {
	Execute(ctx context.Context, cmd *userservice.IntrospectCommand) (*usermodel.User, error)
}

type UserHTTPController struct {
	registerCmdHandler     IRegisterCommandHandler
	authenticateCmdHandler IAuthenticateCommandHandler
	introspectCmdHandler   IIntrospectCommandHandler
}

func NewUserHTTPController(
	registerCmdHandler IRegisterCommandHandler,
	authenticateCmdHandler IAuthenticateCommandHandler,
	introspectCmdHandler IIntrospectCommandHandler,
) *UserHTTPController {
	return &UserHTTPController{
		registerCmdHandler:     registerCmdHandler,
		authenticateCmdHandler: authenticateCmdHandler,
		introspectCmdHandler:   introspectCmdHandler,
	}
}

func (ctrl *UserHTTPController) SetupRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.POST("/register", ctrl.Register)
	router.POST("/authenticate", ctrl.Authenticate)
	
	router.GET("/profile", authMiddleware, ctrl.GetProfile)
	router.POST("/rpc/users/introspect-token", ctrl.IntrospectTokenRpc)
}

func (ctrl *UserHTTPController) SetupAdminRoutes(
	router *gin.RouterGroup,
	authMiddleware gin.HandlerFunc,
	adminRoleMiddleware gin.HandlerFunc) {
	router.GET("/admin/users", authMiddleware, adminRoleMiddleware, func(c *gin.Context) {
		// Admin-specific user management logic can be added here
		c.JSON(200, gin.H{"message": "Admin user management endpoint"})
	})
}
