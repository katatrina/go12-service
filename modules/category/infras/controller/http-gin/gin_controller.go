package categoryhttpgin

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICategoryService interface {
	// CreateNewCategory(ctx context.Context, data *categorymodel.Category) error
	// GetCategoryDetails(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	ListCategories(ctx context.Context,
		dto *categoryservice.ListCategoriesDTO,
	) ([]categorymodel.Category, error)
	UpdateCategoryByID(ctx context.Context, cmd *categoryservice.UpdateCategoryCommandDTO) error
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
}

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.CreateCommand) (*categorymodel.Category, error)
}

type IDetailQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.GetDetailQuery) (*categorymodel.Category, error)
}

type CategoryHTTPController struct {
	catService      ICategoryService
	createCmdHdl    ICreateCommandHandler
	getDetailQryHdl IDetailQueryHandler
}

func NewCategoryHTTPController(
	catService ICategoryService,
	getDetailQryHdl IDetailQueryHandler,
	createNewCmdHdl ICreateCommandHandler,
) *CategoryHTTPController {
	return &CategoryHTTPController{
		catService:      catService,
		getDetailQryHdl: getDetailQryHdl,
		createCmdHdl:    createNewCmdHdl,
	}
}

func (ctl *CategoryHTTPController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategory)
	g.GET("", ctl.ListCategoriesAPI)
	g.GET("/:id", ctl.GetCategoryByIDAPI)
	g.PATCH("/:id", ctl.UpdateCategoryByIDAPI)
	g.DELETE("/:id", ctl.DeleteCategoryByIDAPI)
}
