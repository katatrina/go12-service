package mysqlrepository

import (
	"context"

	foodmodel "github.com/katatrina/go12-service/modules/food/model"

	"github.com/google/uuid"
)

func (repo *FoodRepository) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]foodmodel.Food, error) {
	db := repo.dbCtx.GetMainConnection()
	
	var foods []foodmodel.Food
	
	if err := db.WithContext(ctx).Where("id IN ? AND status != ?", ids, foodmodel.FoodStatusDeleted).Find(&foods).Error; err != nil {
		return nil, err
	}
	
	return foods, nil
}