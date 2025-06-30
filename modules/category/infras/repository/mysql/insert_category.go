package mysqlrepository

import (
	"context"
	
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
)

func (repo *CategoryRepository) Insert(ctx context.Context, data *categorymodel.Category) error {
	if err := repo.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
