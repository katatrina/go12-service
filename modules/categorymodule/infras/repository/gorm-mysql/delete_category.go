package categorygormmysql

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/categorymodule/internal/model"
)

func (repo *CategoryRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	if isHard {
		if err := repo.db.Model(&categorymodel.Category{}).Where("id = ?", id).Delete(nil).Error; err != nil {
			return err
		}
		
		return nil
	}
	
	if err := repo.db.Model(&categorymodel.Category{}).Where("id = ?", id).Update("status", categorymodel.StatusDeleted).Error; err != nil {
		return err
	}
	
	return nil
}
