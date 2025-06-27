package repository

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/restaurant/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func (repo *RestaurantRepository) ListRestaurants(
	ctx context.Context,
	pagingDTO *sharedmodel.PagingDTO,
	filterDTO *model.FilterRestaurantDTO,
) ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	
	db := repo.db.WithContext(ctx).Model(&model.Restaurant{})
	
	if filterDTO.Status != nil {
		db = db.Where("status = ?", *filterDTO.Status)
	}
	if filterDTO.CityID != nil {
		db = db.Where("city_id = ?", *filterDTO.CityID)
	}
	if filterDTO.CategoryID != nil {
		db = db.Where("category_id = ?", *filterDTO.CategoryID)
	}
	
	if err := db.Count(&pagingDTO.Total).Error; err != nil {
		return nil, err
	}
	
	db = db.Preload("Category")
	
	offset := (pagingDTO.Page - 1) * pagingDTO.Limit
	if err := db.Offset(offset).Limit(pagingDTO.Limit).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}
