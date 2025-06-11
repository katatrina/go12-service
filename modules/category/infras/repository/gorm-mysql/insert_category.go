package categorygormmysql

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

func (repo *CategoryRepository) Insert(ctx context.Context, data *categorymodel.categorymodel) error {
	if err := repo.db.Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
