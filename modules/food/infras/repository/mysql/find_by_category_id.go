package mysqlrepository

import (
	"context"

	foodmodel "github.com/katatrina/go12-service/modules/food/model"

	"github.com/google/uuid"
)

func (repo *FoodRepository) FindByCategoryID(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]foodmodel.Food, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var foods []foodmodel.Food
	
	query := db.WithContext(ctx).Where("category_id = ? AND status != ?", categoryID, foodmodel.FoodStatusDeleted)
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	if offset > 0 {
		query = query.Offset(offset)
	}
	
	if err := query.Order("created_at DESC").Find(&foods).Error; err != nil {
		return nil, err
	}
	
	return foods, nil
}