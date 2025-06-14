package categorygormmysql

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/categorymodule/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func (repo *CategoryRepository) ListCategories(
	ctx context.Context,
	pagingDTO *sharedmodel.PagingDTO,
	filterDTO *categorymodel.FilterCategoryDTO,
) ([]categorymodel.Category, error) {
	var categories []categorymodel.Category
	
	query := repo.db.WithContext(ctx).
		Where("status in (?)", []string{categorymodel.StatusActive})
	
	// Lấy tổng số records sau khi đã filterDTO
	if err := query.Table((&categorymodel.Category{}).TableName()).Count(&pagingDTO.Total).Error; err != nil {
		return nil, err
	}
	
	// Nếu không có records nào, trả về slice rỗng
	if pagingDTO.Total == 0 {
		return categories, nil
	}
	
	// Tính toán offset và limit
	offset := (pagingDTO.Page - 1) * pagingDTO.Limit
	
	// Lấy data với pagingDTO và sắp xếp
	if err := query.
		Order("created_at desc").
		Offset(offset).
		Limit(pagingDTO.Limit).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	
	return categories, nil
}
