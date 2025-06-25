package repository

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/category/model"
)

func (repo *CategoryRepository) Insert(ctx context.Context, data *model.Category) error {
	if err := repo.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	
	return nil
}
