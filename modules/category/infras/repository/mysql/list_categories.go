package repository

import (
	"context"

	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func (repo *CategoryRepository) ListCategories(
	ctx context.Context,
	pagingDTO *sharedmodel.PagingDTO,
	filterDTO *model.FilterCategoryDTO,
) ([]model.Category, error) {
	var categories []model.Category

	// TODO: Check the logic again for best practices

	query := repo.db.WithContext(ctx).
		Where("status in (?)", []string{string(datatype.StatusActive)})

	if err := query.Table((&model.Category{}).TableName()).Count(&pagingDTO.Total).Error; err != nil {
		return nil, err
	}

	if pagingDTO.Total == 0 {
		return categories, nil
	}

	offset := (pagingDTO.Page - 1) * pagingDTO.Limit

	if err := query.
		Order("created_at desc").
		Offset(offset).
		Limit(pagingDTO.Limit).
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
