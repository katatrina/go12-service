package categorygormmysql

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/categorymodule/internal/model"
)

func (repo *CategoryRepository) Update(ctx context.Context, id uuid.UUID, dto *categorymodel.UpdateCategoryDTO) error {
	db := repo.db.Begin()
	
	if err := repo.db.Model(&categorymodel.Category{}).Where("id = ?", id).Updates(dto).Error; err != nil {
		db.Rollback()
		return err
	}
	
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	
	return nil
}
