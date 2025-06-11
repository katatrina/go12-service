package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

func (s *CategoryService) CreateNewCategory(ctx context.Context, data *categorymodel.categorymodel) error {
	// business logic to create a new category
	
	if err := data.Validate(); err != nil {
		return err
	}
	
	data.ID, _ = uuid.NewV7()
	
	if err := s.catRepo.Insert(ctx, data); err != nil {
		return err
	}
	
	return nil
}
