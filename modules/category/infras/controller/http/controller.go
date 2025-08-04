package httpcontroller

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	createCmdHdl    ICreateCommandHandler
	getQryHdl       IGetByIDQueryHandler
	listQryHdl      IListQueryHandler
	updateCmdHdl    IUpdateByIDCommandHandler
	deleteCmdHdl    IDeleteByIDCommandHandler
	categoryRPCRepo ICategoryRPCRepo
}

type ICategoryRPCRepo interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]categorymodel.Category, error)
}

func NewCategoryController(
	createNewCmdHdl ICreateCommandHandler,
	getDetailQryHdl IGetByIDQueryHandler,
	listQryHdl IListQueryHandler,
	updateCmdHdl IUpdateByIDCommandHandler,
	deleteCmdHdl IDeleteByIDCommandHandler,
	categoryRPCRepo ICategoryRPCRepo,
) *CategoryController {
	return &CategoryController{
		createCmdHdl:    createNewCmdHdl,
		getQryHdl:       getDetailQryHdl,
		listQryHdl:      listQryHdl,
		updateCmdHdl:    updateCmdHdl,
		deleteCmdHdl:    deleteCmdHdl,
		categoryRPCRepo: categoryRPCRepo,
	}
}

func (ctl *CategoryController) SetupRoutes(g *gin.RouterGroup) {
	{
		g.POST("", ctl.CreateCategory)
		g.GET("", ctl.ListCategories)
		g.GET("/:id", ctl.GetCategoryByID)
		g.PATCH("/:id", ctl.UpdateCategoryByID)
		g.DELETE("/:id", ctl.DeleteCategoryByID)
	}
}

func (ctl *CategoryController) SetupRoutesForRPC(g *gin.RouterGroup) {
	{
		g.POST("/rpc/categories/find-by-ids", ctl.RPCGetByIDS)
	}
}
