package mysqlrepository

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (repo *RestaurantRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	if isHard {
		err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(restaurantmodel.Restaurant{}).Error
		if err != nil {
			return err
		}
		
		return nil
	}
	
	// Soft delete
	err := repo.db.WithContext(ctx).Model(restaurantmodel.Restaurant{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     datatype.StatusDeleted,
			"updated_at": time.Now().UTC(),
		}).Error
	if err != nil {
		return err
	}
	
	return nil
}
