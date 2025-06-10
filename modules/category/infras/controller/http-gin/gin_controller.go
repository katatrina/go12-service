package categoryhttpgin

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
	categoryservice "github.com/katatrina/go12-service/modules/category/internal/service"
)

type ICategoryService interface {
	CreateNewCategory(ctx context.Context, data *categorymodel.Category) error
	GetCategoryDetails(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	ListCategories(ctx context.Context,
		dto *categoryservice.ListCategoriesDTO,
	) ([]categorymodel.Category, error)
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
}
