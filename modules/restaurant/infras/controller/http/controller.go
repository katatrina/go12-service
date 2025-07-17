package httpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *restaurantservice.CreateCommand) (*restaurantmodel.Restaurant, error)
}

type IGetByIDQueryHandler interface {
	Execute(ctx context.Context, query *restaurantservice.GetByIDQuery) (*restaurantmodel.Restaurant, error)
}

type IListQueryHandler interface {
	Execute(ctx context.Context, query *restaurantservice.ListQuery) ([]restaurantservice.ListRestaurantItemDTO, error)
}

type IUpdateByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *restaurantservice.UpdateByIDCommand) error
}

type IDeleteByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *restaurantservice.DeleteByIDCommand) error
}

type RestaurantController struct {
	createCmdHdl ICreateCommandHandler
	getQryHdl    IGetByIDQueryHandler
	listQryHdl   IListQueryHandler
	updateCmdHdl IUpdateByIDCommandHandler
	deleteCmdHdl IDeleteByIDCommandHandler
}

func NewRestaurantController(
	createNewCmdHdl ICreateCommandHandler,
	getDetailQryHdl IGetByIDQueryHandler,
	listQryHdl IListQueryHandler,
	updateCmdHdl IUpdateByIDCommandHandler,
	deleteCmdHdl IDeleteByIDCommandHandler,
) *RestaurantController {
	return &RestaurantController{
		createCmdHdl: createNewCmdHdl,
		getQryHdl:    getDetailQryHdl,
		listQryHdl:   listQryHdl,
		updateCmdHdl: updateCmdHdl,
		deleteCmdHdl: deleteCmdHdl,
	}
}

func (ctl *RestaurantController) SetupRoutes(g *gin.RouterGroup, mldProvider sharedinfras.IMiddlewareProvider) {
	restaurantGroup := g.Group("/restaurants")
	{
		restaurantGroup.POST("", mldProvider.Auth(), ctl.CreateRestaurant)
		restaurantGroup.GET("", ctl.ListRestaurants)
		restaurantGroup.GET("/:id", ctl.GetRestaurantByID)
		restaurantGroup.PATCH("/:id", ctl.UpdateRestaurantByID)
		restaurantGroup.DELETE("/:id", ctl.DeleteRestaurantByID)
	}
}
