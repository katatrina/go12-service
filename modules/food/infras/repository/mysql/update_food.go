package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (repo *FoodRepository) Update(ctx context.Context, id uuid.UUID, data map[string]interface{}) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.WithContext(ctx).Table("foods").Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}