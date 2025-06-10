package categoryservice

import (
	"context"
	
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

type ListCategoriesDTO struct {
	sharedmodel.PagingDTO
	categorymodel.FilterCategoryDTO
}

func (s *CategoryService) ListCategories(ctx context.Context,
	dto *ListCategoriesDTO,
) ([]categorymodel.Category, error) {
	categories, err := s.catRepo.ListCategories(ctx, &dto.PagingDTO, &dto.FilterCategoryDTO)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}
