package httpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	foodservice "github.com/katatrina/go12-service/modules/food/service"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *foodservice.CreateCommand) (*foodmodel.Food, error)
}

type IGetByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *foodservice.GetByIDCommand) (*foodmodel.FoodResponseDTO, error)
}

type IListCommandHandler interface {
	Execute(ctx context.Context, cmd *foodservice.ListCommand) (*foodmodel.FoodListResponseDTO, error)
}

type IUpdateCommandHandler interface {
	Execute(ctx context.Context, cmd *foodservice.UpdateCommand) (*foodmodel.Food, error)
}

type IDeleteCommandHandler interface {
	Execute(ctx context.Context, cmd *foodservice.DeleteCommand) error
}

type FoodController struct {
	createCmdHdl ICreateCommandHandler
	getCmdHdl    IGetByIDCommandHandler
	listCmdHdl   IListCommandHandler
	updateCmdHdl IUpdateCommandHandler
	deleteCmdHdl IDeleteCommandHandler
}

func NewFoodController(
	createCmdHdl ICreateCommandHandler,
	getCmdHdl IGetByIDCommandHandler,
	listCmdHdl IListCommandHandler,
	updateCmdHdl IUpdateCommandHandler,
	deleteCmdHdl IDeleteCommandHandler,
) *FoodController {
	return &FoodController{
		createCmdHdl: createCmdHdl,
		getCmdHdl:    getCmdHdl,
		listCmdHdl:   listCmdHdl,
		updateCmdHdl: updateCmdHdl,
		deleteCmdHdl: deleteCmdHdl,
	}
}

func (ctl *FoodController) SetupRoutes(g *gin.RouterGroup, mldProvider sharedinfras.IMiddlewareProvider) {
	foodGroup := g.Group("/foods")
	{
		foodGroup.POST("", mldProvider.Auth(), ctl.CreateFood)
		foodGroup.GET("", ctl.ListFoods)
		foodGroup.GET("/:id", ctl.GetFoodByID)
		foodGroup.PATCH("/:id", mldProvider.Auth(), ctl.UpdateFoodByID)
		foodGroup.DELETE("/:id", mldProvider.Auth(), ctl.DeleteFoodByID)
	}
}