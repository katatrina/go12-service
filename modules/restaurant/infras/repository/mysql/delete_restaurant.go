package repository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (repo *RestaurantRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	if isHard {
		err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Restaurant{}).Error
		if err != nil {
			return err
		}
		return nil
	}
	// Soft delete
	err := repo.db.WithContext(ctx).Model(&model.Restaurant{}).
		Where("id = ?", id).
		Update("status", datatype.StatusDeleted).Error
	if err != nil {
		return err
	}
	return nil
}
