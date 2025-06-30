package mysqlrepository

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	"gorm.io/gorm"
)

func (repo *CategoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error) {
	var category categorymodel.Category
	
	if err := repo.db.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, categorymodel.ErrCategoryNotFound
		}
		return nil, err
	}
	
	return &category, nil
}
