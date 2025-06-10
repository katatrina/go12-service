package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

type ICategoryRepository interface {
	Insert(ctx context.Context, category *categorymodel.Category) error
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	ListCategories(ctx context.Context,
		pagingDTO *sharedmodel.PagingDTO,
		filterDTO *categorymodel.FilterCategoryDTO,
	) ([]categorymodel.Category, error)
}

// Dependency Injection by constructor/new function
func NewCategoryService(catRepo ICategoryRepository) *CategoryService {
	return &CategoryService{
		catRepo: catRepo,
	}
}

// Dependency Injection by setter method
// func (s *CategoryService) SetCategoryRepository(catRepo ICategoryRepository) {
// 	s.catRepo = catRepo
// }

type CategoryService struct {
	catRepo ICategoryRepository // private
}
