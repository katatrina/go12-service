package mysqlrepository

import (
	"context"
	
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/pkg/errors"
)

func (repo *FoodRepository) Insert(ctx context.Context, data *foodmodel.Food) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}