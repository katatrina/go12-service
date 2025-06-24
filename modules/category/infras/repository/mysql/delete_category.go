package repository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (repo *CategoryRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	if isHard {
		result := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Category{})
		if result.Error != nil {
			return result.Error
		}
		
		if result.RowsAffected == 0 {
			return model.ErrCategoryNotFound
		}
		
		return nil
	}
	
	// Soft delete
	result := repo.db.WithContext(ctx).Model(&model.Category{}).
		Where("id = ?", id).
		Update("status", datatype.StatusDeleted)
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return model.ErrCategoryNotFound
	}
	
	return nil
}
