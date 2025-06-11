package categoryhttpgin

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice2 "github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICategoryService interface {
	CreateNewCategory(ctx context.Context, data *categorymodel.categorymodel) error
	GetCategoryDetails(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	ListCategories(ctx context.Context,
		dto *categoryservice2.ListCategoriesDTO,
	) ([]categorymodel.Category, error)
	UpdateCategoryByID(ctx context.Context, cmd *categoryservice2.UpdateCategoryCommandDTO) error
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
}

type CategoryHTTPController struct {
	catService ICategoryService
}

func NewCategoryHTTPController(catService ICategoryService) *CategoryHTTPController {
	return &CategoryHTTPController{
		catService: catService,
	}
}

func (ctl *CategoryHTTPController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategoryAPI)
	g.GET("", ctl.ListCategoriesAPI)
	g.GET(":id", ctl.GetCategoryByIDAPI)
	g.PATCH(":id", ctl.UpdateCategoryByIDAPI)
	g.DELETE(":id", ctl.DeleteCategoryByIDAPI)
}
