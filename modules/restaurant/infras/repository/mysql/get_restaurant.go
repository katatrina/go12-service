package mysqlrepository

import (
	"context"
	stderrors "errors"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *RestaurantRepository) FindByID(ctx context.Context, id uuid.UUID) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	
	if err := repo.db.WithContext(ctx).First(&restaurant, "id1 = ?", id).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, datatype.ErrRecordNotFound
		}
		
		return nil, errors.WithStack(err)
	}
	
	return &restaurant, nil
}
