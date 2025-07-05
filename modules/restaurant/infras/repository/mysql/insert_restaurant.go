package mysqlrepository

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/restaurant/model"
)

func (repo *RestaurantRepository) Insert(ctx context.Context, data *restaurantmodel.Restaurant) error {
	if err := repo.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
