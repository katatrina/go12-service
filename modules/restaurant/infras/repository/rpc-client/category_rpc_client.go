package restaurantrpcclient

import (
	"context"
	
	"github.com/google/uuid"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"gorm.io/gorm"
)

type CategoryRPCClient struct {
	db *gorm.DB
}

func NewCategoryRPCClient(db *gorm.DB) *CategoryRPCClient {
	return &CategoryRPCClient{
		db: db,
	}
}

func (c *CategoryRPCClient) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Category, error) {
	var categories []restaurantmodel.Category
	
	if err := c.db.WithContext(ctx).Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return nil, err
	}
	
	return categories, nil
}
