package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (repo *CategoryRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	if isHard {
		err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&categorymodel.Category{}).Error
		if err != nil {
			return err
		}
		
		return nil
	}
	
	// Soft delete
	err := repo.db.WithContext(ctx).Model(&categorymodel.Category{}).
		Where("id = ?", id).
		Update("status", datatype.StatusDeleted).Error
	
	if err != nil {
		return err
	}
	
	return nil
}
