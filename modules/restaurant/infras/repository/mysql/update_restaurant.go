package repository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
)

func (repo *RestaurantRepository) Update(ctx context.Context, id uuid.UUID, dto *restaurantmodel.UpdateRestaurantDTO) error {
	return repo.db.WithContext(ctx).Model(&restaurantmodel.Restaurant{}).Where("id = ?", id).Updates(dto).Error
}
