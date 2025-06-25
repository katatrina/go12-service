package controller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/modules/restaurant/service"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *service.CreateCommand) (*model.Restaurant, error)
}

type IGetByIDQueryHandler interface {
	Execute(ctx context.Context, query *service.GetByIDQuery) (*model.Restaurant, error)
}

type IListQueryHandler interface {
	Execute(ctx context.Context, query *service.ListQuery) ([]model.Restaurant, error)
}

type IUpdateByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *service.UpdateByIDCommand) error
}

type IDeleteByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *service.DeleteByIDCommand) error
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

func (ctl *RestaurantController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateRestaurant)
	g.GET("", ctl.ListRestaurants)
	g.GET("/:id", ctl.GetRestaurantByID)
	g.PATCH("/:id", ctl.UpdateRestaurantByID)
	g.DELETE("/:id", ctl.DeleteRestaurantByID)
}
