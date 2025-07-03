package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
)

func (repo *CategoryRepository) FindByIDs(ctx context.Context, ids []*uuid.UUID) ([]categorymodel.Category, error) {
	var categories []categorymodel.Category
	
	err := repo.db.WithContext(ctx).
		Where("id IN (?)", ids).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}
