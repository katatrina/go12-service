package gormmysql

import (
	"context"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (repo *RestaurantLikeRepository) DeleteLike(
	ctx context.Context,
	restaurantID, userID uuid.UUID,
) error {
	db := repo.dbCtx.GetMainConnection()
	
	if err := db.Where("restaurant_id = ? AND user_id = ?", restaurantID.String(), userID.String()).
		Delete(&restaurantlikemodel.RestaurantLike{}).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
