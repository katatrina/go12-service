package restaurantrepository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/pkg/errors"
)

func (repo *RestaurantRepository) Update(ctx context.Context, id uuid.UUID, dto *restaurantmodel.UpdateRestaurantDTO) error {
	db := repo.dbCtx.GetMainConnection()
	
	err := db.WithContext(ctx).Where("id = ?", id).Updates(dto).Error
	if err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
