package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

func (repo *CategoryRepository) Update(ctx context.Context, id uuid.UUID, dto *model.UpdateCategoryDTO) error {
	db := repo.db.Begin()

	if err := repo.db.Model(&model.Category{}).Where("id = ?", id).Updates(dto).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
