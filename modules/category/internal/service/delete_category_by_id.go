package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

func (s *CategoryService) DeleteCategoryByID(ctx context.Context, id uuid.UUID) error {
	category, err := s.catRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	
	if category.Status == categorymodel.StatusDeleted {
		return categorymodel.ErrCategoryDeleted
	}
	
	if err = s.catRepo.Delete(ctx, id, false); err != nil {
		return err
	}
	
	return nil
}
