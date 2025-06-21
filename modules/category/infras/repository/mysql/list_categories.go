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
	
	// Build base query
	baseQuery := repo.db.WithContext(ctx).Model(&model.Category{})
	
	// Apply filters if any
	if filterDTO.Status != nil {
		baseQuery = baseQuery.Where("status = ?", *filterDTO.Status)
	} else {
		// Default to active status if no filter is provided
		baseQuery = baseQuery.Where("status = ?", datatype.StatusActive)
	}
	
	// Count total
	if err := baseQuery.Count(&pagingDTO.Total).Error; err != nil {
		return nil, err
	}
	
	if pagingDTO.Total == 0 {
		return categories, nil
	}
	
	// Get paginated results
	offset := (pagingDTO.Page - 1) * pagingDTO.Limit
	if err := baseQuery.
		Order("created_at DESC").
		Offset(offset).
		Limit(pagingDTO.Limit).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	
	return categories, nil
}
