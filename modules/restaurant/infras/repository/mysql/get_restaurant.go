package mysqlrepository

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"gorm.io/gorm"
)

func (repo *RestaurantRepository) FindByID(ctx context.Context, id uuid.UUID) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	if err := repo.db.WithContext(ctx).First(&restaurant, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, restaurantmodel.ErrRestaurantNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}
