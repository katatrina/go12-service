package categoryhttpgin

import (
	"context"
	
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
)

type ICategoryService interface {
	CreateNewCategory(ctx context.Context, data *categorymodel.Category) error
}

func NewCategoryHTTPController(catService ICategoryService) *CategoryHTTPController {
	return &CategoryHTTPController{
		catService: catService,
	}
}

type CategoryHTTPController struct {
	catService ICategoryService
}
