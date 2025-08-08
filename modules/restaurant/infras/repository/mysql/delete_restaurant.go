package restaurantrepository

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
)

func (repo *RestaurantRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	db := repo.dbCtx.GetMainConnection()
	
	if isHard {
		err := db.WithContext(ctx).Where("id = ?", id).Delete(restaurantmodel.Restaurant{}).Error
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		// Soft delete
		err := db.WithContext(ctx).Model(restaurantmodel.Restaurant{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"status":     datatype.StatusDeleted,
				"updated_at": time.Now().UTC(),
			}).Error
		
		if err != nil {
			return errors.WithStack(err)
		}
	}
	
	return nil
}
