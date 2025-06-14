package categorygormmysql

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/categorymodule/internal/model"
)

func (repo *CategoryRepository) Insert(ctx context.Context, data *categorymodel.Category) error {
	if err := repo.db.Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
