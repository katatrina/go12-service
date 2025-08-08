package restaurantrepository

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/restaurant/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (repo *RestaurantRepository) List(
	ctx context.Context,
	pagingDTO *sharedmodel.PagingDTO,
	filterDTO *restaurantmodel.FilterRestaurantDTO,
) ([]restaurantmodel.Restaurant, error) {
	_, span := otel.Tracer("go12-service").Start(ctx, "list-restaurants-gorm")
	defer span.End()
	
	db := repo.dbCtx.GetMainConnection()
	
	var restaurants []restaurantmodel.Restaurant
	
	db = db.WithContext(ctx).Model(&restaurantmodel.Restaurant{})
	
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
		return nil, errors.WithStack(err)
	}
	
	// db = db.Preload("Category")
	
	offset := (pagingDTO.Page - 1) * pagingDTO.Limit
	if err := db.Offset(offset).Limit(pagingDTO.Limit).Find(&restaurants).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return restaurants, nil
}
