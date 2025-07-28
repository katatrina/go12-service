package gormmysql

import (
	"context"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (repo *RestaurantLikeRepository) DeleteLike(
	ctx context.Context,
	restaurantId, userId uuid.UUID,
) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.Where("restaurant_id = ? AND user_id = ?", restaurantId.String(), userId.String()).
		Delete(&restaurantlikemodel.RestaurantLike{}).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
