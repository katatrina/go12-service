package mysqlrepository

import (
	"context"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (repo *FoodRepository) List(ctx context.Context, filter *foodmodel.FoodFilterDTO, offset, limit int) ([]*foodmodel.Food, error) {
	_, span := otel.Tracer("go12-service").Start(ctx, "food-repo-mysql.list")
	defer span.End()
	
	db := repo.dbCtx.GetMainConnection()
	
	var foods []*foodmodel.Food
	
	query := db.WithContext(ctx).Where("status <> ?", foodmodel.FoodStatusDeleted)
	
	if filter != nil {
		if filter.RestaurantID != nil {
			if restaurantUUID, err := uuid.Parse(*filter.RestaurantID); err == nil {
				query = query.Where("restaurant_id = ?", restaurantUUID)
			}
		}
		
		if filter.CategoryID != nil {
			if categoryUUID, err := uuid.Parse(*filter.CategoryID); err == nil {
				query = query.Where("category_id = ?", categoryUUID)
			}
		}
		
		if filter.MinPrice != nil {
			query = query.Where("price >= ?", *filter.MinPrice)
		}
		
		if filter.MaxPrice != nil {
			query = query.Where("price <= ?", *filter.MaxPrice)
		}
		
		if filter.Search != nil && *filter.Search != "" {
			query = query.Where("name LIKE ? OR description LIKE ?", "%"+*filter.Search+"%", "%"+*filter.Search+"%")
		}
	}
	
	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&foods).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return foods, nil
}

func (repo *FoodRepository) Count(ctx context.Context, filter *foodmodel.FoodFilterDTO) (int64, error) {
	_, span := otel.Tracer("go12-service").Start(ctx, "food-repo-mysql.count")
	defer span.End()
	
	db := repo.dbCtx.GetMainConnection()
	
	var count int64
	
	query := db.WithContext(ctx).Model(&foodmodel.Food{}).Where("status <> ?", foodmodel.FoodStatusDeleted)
	
	if filter != nil {
		if filter.RestaurantID != nil {
			if restaurantUUID, err := uuid.Parse(*filter.RestaurantID); err == nil {
				query = query.Where("restaurant_id = ?", restaurantUUID)
			}
		}
		
		if filter.CategoryID != nil {
			if categoryUUID, err := uuid.Parse(*filter.CategoryID); err == nil {
				query = query.Where("category_id = ?", categoryUUID)
			}
		}
		
		if filter.MinPrice != nil {
			query = query.Where("price >= ?", *filter.MinPrice)
		}
		
		if filter.MaxPrice != nil {
			query = query.Where("price <= ?", *filter.MaxPrice)
		}
		
		if filter.Search != nil && *filter.Search != "" {
			query = query.Where("name LIKE ? OR description LIKE ?", "%"+*filter.Search+"%", "%"+*filter.Search+"%")
		}
	}
	
	if err := query.Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	
	return count, nil
}