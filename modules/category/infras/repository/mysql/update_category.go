package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
)

func (repo *CategoryRepository) Update(ctx context.Context, id uuid.UUID, dto *categorymodel.UpdateCategoryDTO) error {
	return repo.db.WithContext(ctx).
		Model(&categorymodel.Category{}).
		Where("id = ?", id).
		Updates(dto).Error
}
