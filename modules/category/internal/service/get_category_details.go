package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func (s *CategoryService) GetCategoryDetails(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error) {
	category, err := s.catRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if category.Status == categorymodel.StatusDeleted {
		return nil, sharedmodel.ErrRecordNotFound
	}
	
	return category, nil
}
