package restaurantlikehttpgin

import (
	"context"
	
	restaurantlikeservice "github.com/katatrina/go12-service/modules/restaurantlike/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	
	"github.com/gin-gonic/gin"
)

type ILikeRestaurantCommandHandler interface {
	Execute(ctx context.Context, cmd *restaurantlikeservice.LikeRestaurantCommand) error
}

type IUnlikeRestaurantCommandHandler interface {
	Execute(ctx context.Context, cmd *restaurantlikeservice.UnlikeRestaurantCommand) error
}

type RestaurantLikeHTTPController struct {
	likeHandler   ILikeRestaurantCommandHandler
	unlikeHandler IUnlikeRestaurantCommandHandler
}

func NewRestaurantLikeHTTPController(
	likeHandler ILikeRestaurantCommandHandler,
	unlikeHandler IUnlikeRestaurantCommandHandler,
) *RestaurantLikeHTTPController {
	return &RestaurantLikeHTTPController{
		likeHandler:   likeHandler,
		unlikeHandler: unlikeHandler,
	}
}

func (ctrl *RestaurantLikeHTTPController) SetupRoutes(router *gin.RouterGroup, mldProvider sharedinfras.IMiddlewareProvider) {
	restaurants := router.Group("/restaurants")
	{
		restaurants.POST("/:id/like", mldProvider.Auth(), ctrl.LikeRestaurantAPI)
		restaurants.DELETE("/:id/unlike", mldProvider.Auth(), ctrl.UnlikeRestaurantAPI)
	}
}
