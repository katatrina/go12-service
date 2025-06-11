package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

type UpdateCategoryCommandDTO struct {
	ID uuid.UUID `json:"-"`
	categorymodel.UpdateCategoryDTO
}

func (s *CategoryService) UpdateCategoryByID(ctx context.Context, cmd *UpdateCategoryCommandDTO) error {
	category, err := s.catRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	if category.Status == categorymodel.StatusDeleted {
		return categorymodel.ErrCategoryDeleted
	}
	
	if err = s.catRepo.Update(ctx, cmd.ID, &cmd.UpdateCategoryDTO); err != nil {
		return err
	}
	
	return nil
}
