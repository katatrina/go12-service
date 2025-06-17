package categoryhttpgin

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICategoryService interface {
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
}

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

type CategoryHTTPController struct {
	catService      ICategoryService
	createCmdHdl    ICreateCommandHandler
	getDetailQryHdl IGetDetailQueryHandler
	listQryHdl      IListQueryHandler
	updateCmdHdl    IUpdateByIDCommandHandler
}

func NewCategoryHTTPController(
	catService ICategoryService,
	getDetailQryHdl IGetDetailQueryHandler,
	createNewCmdHdl ICreateCommandHandler,
	listQryHdl IListQueryHandler,
	updateCmdHdl IUpdateByIDCommandHandler,
) *CategoryHTTPController {
	return &CategoryHTTPController{
		catService:      catService,
		getDetailQryHdl: getDetailQryHdl,
		createCmdHdl:    createNewCmdHdl,
		listQryHdl:      listQryHdl,
		updateCmdHdl:    updateCmdHdl,
	}
}

func (ctl *CategoryHTTPController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategory)
	g.GET("", ctl.ListCategories)
	g.GET("/:id", ctl.GetCategoryByID)
	g.PATCH("/:id", ctl.UpdateCategoryByID)
	g.DELETE("/:id", ctl.DeleteCategoryByIDAPI)
}
