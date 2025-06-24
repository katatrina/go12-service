package controller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *service.CreateCommand) (*model.Category, error)
}

type IGetByIDQueryHandler interface {
	Execute(ctx context.Context, query *service.GetByIDQuery) (*model.Category, error)
}

type IListQueryHandler interface {
	Execute(ctx context.Context, query *service.ListQuery) ([]model.Category, error)
}

type IUpdateByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *service.UpdateByIDCommand) error
}

type IDeleteByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *service.DeleteByIDCommand) error
}

type CategoryController struct {
	createCmdHdl ICreateCommandHandler
	getQryHdl    IGetByIDQueryHandler
	listQryHdl   IListQueryHandler
	updateCmdHdl IUpdateByIDCommandHandler
	deleteCmdHdl IDeleteByIDCommandHandler
}

func NewCategoryController(
	createNewCmdHdl ICreateCommandHandler,
	getDetailQryHdl IGetByIDQueryHandler,
	listQryHdl IListQueryHandler,
	updateCmdHdl IUpdateByIDCommandHandler,
	deleteCmdHdl IDeleteByIDCommandHandler,
) *CategoryController {
	return &CategoryController{
		createCmdHdl: createNewCmdHdl,
		getQryHdl:    getDetailQryHdl,
		listQryHdl:   listQryHdl,
		updateCmdHdl: updateCmdHdl,
		deleteCmdHdl: deleteCmdHdl,
	}
}

func (ctl *CategoryController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategory)
	g.GET("", ctl.ListCategories)
	g.GET("/:id", ctl.GetCategoryByID)
	g.PATCH("/:id", ctl.UpdateCategoryByID)
	g.DELETE("/:id", ctl.DeleteCategoryByID)
}
