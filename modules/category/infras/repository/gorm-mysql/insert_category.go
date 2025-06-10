package categorygormmysql

import (
	"context"
	
	categorymodel "github.com/katatrina/go12-service/modules/category/internal/model"
)

func (repo *CategoryRepository) Insert(ctx context.Context, data *categorymodel.Category) error {
	if err := repo.db.Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
