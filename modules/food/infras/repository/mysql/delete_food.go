package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/pkg/errors"
)

func (repo *FoodRepository) Delete(ctx context.Context, id uuid.UUID) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.WithContext(ctx).Table("foods").Where("id = ?", id).Update("status", foodmodel.FoodStatusDeleted).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}