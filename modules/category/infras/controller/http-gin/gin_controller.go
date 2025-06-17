package categoryhttpgin

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.CreateCommand) (*categorymodel.Category, error)
}

type IGetDetailQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.GetDetailQuery) (*categorymodel.Category, error)
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

type CategoryHTTPController struct {
	createCmdHdl    ICreateCommandHandler
	getDetailQryHdl IGetDetailQueryHandler
	listQryHdl      IListQueryHandler
	updateCmdHdl    IUpdateByIDCommandHandler
	deleteCmdHdl    IDeleteByIDCommandHandler
}

func NewCategoryHTTPController(
	createNewCmdHdl ICreateCommandHandler,
	getDetailQryHdl IGetDetailQueryHandler,
	listQryHdl IListQueryHandler,
	updateCmdHdl IUpdateByIDCommandHandler,
	deleteCmdHdl IDeleteByIDCommandHandler,
) *CategoryHTTPController {
	return &CategoryHTTPController{
		createCmdHdl:    createNewCmdHdl,
		getDetailQryHdl: getDetailQryHdl,
		listQryHdl:      listQryHdl,
		updateCmdHdl:    updateCmdHdl,
		deleteCmdHdl:    deleteCmdHdl,
	}
}

func (ctl *CategoryHTTPController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategory)
	g.GET("", ctl.ListCategories)
	g.GET("/:id", ctl.GetCategoryByID)
	g.PATCH("/:id", ctl.UpdateCategoryByID)
	g.DELETE("/:id", ctl.DeleteCategoryByIDAPI)
}
