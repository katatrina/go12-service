package restaurantrepository

import (
	"context"

	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"

	"github.com/google/uuid"
)

func (repo *RestaurantRepository) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Restaurant, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var restaurants []restaurantmodel.Restaurant
	
	if err := db.WithContext(ctx).Where("id IN ? AND status != ?", ids, datatype.StatusDeleted).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	
	return restaurants, nil
}