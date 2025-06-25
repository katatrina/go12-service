package repository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/model"
)

func (repo *CategoryRepository) Update(ctx context.Context, id uuid.UUID, dto *model.UpdateCategoryDTO) error {
	return repo.db.WithContext(ctx).
		Model(&model.Category{}).
		Where("id = ?", id).
		Updates(dto).Error
}
