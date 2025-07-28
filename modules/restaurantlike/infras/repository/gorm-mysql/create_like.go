package gormmysql

import (
	"context"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	
	"github.com/pkg/errors"
)

func (repo *RestaurantLikeRepository) CreateLike(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
