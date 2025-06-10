package categoryservice

import (
	"context"
	
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
)

type ICategoryRepository interface {
	Insert(ctx context.Context, category *categorymodel.Category) error
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
