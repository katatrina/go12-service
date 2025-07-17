package mysqlrepository

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/pkg/errors"
)

func (repo *RestaurantRepository) Insert(ctx context.Context, data *restaurantmodel.Restaurant) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
