package restaurantrepository

import (
	"context"

	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"

	"github.com/google/uuid"
)

func (repo *RestaurantRepository) FindByCategoryID(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]restaurantmodel.Restaurant, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var restaurants []restaurantmodel.Restaurant
	
	query := db.WithContext(ctx).Where("category_id = ? AND status != ?", categoryID, datatype.StatusDeleted)
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	if offset > 0 {
		query = query.Offset(offset)
	}
	
	if err := query.Order("created_at DESC").Find(&restaurants).Error; err != nil {
		return nil, err
	}
	
	return restaurants, nil
}