package httpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/service"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.CreateCommand) (*categorymodel.Category, error)
}

type IGetByIDQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.GetByIDQuery) (*categorymodel.Category, error)
}

type IListQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.ListQuery) ([]categorymodel.Category, error)
}

type IUpdateByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.UpdateByIDCommand) error
}

type IDeleteByIDCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.DeleteByIDCommand) error
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
	categoryGroup := g.Group("/categories")
	{
		categoryGroup.POST("", ctl.CreateCategory)
		categoryGroup.GET("", ctl.ListCategories)
		categoryGroup.GET("/:id", ctl.GetCategoryByID)
		categoryGroup.PATCH("/:id", ctl.UpdateCategoryByID)
		categoryGroup.DELETE("/:id", ctl.DeleteCategoryByID)
	}
}
