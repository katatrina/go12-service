package gormmysql

import (
	"context"
	
	restaurantlikemodel "github.com/katatrina/go12-service/modules/restaurantlike/model"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *RestaurantLikeRepository) FindLike(
	ctx context.Context,
	restaurantId, userId uuid.UUID,
) (*restaurantlikemodel.RestaurantLike, error) {
	db := repo.dbCtx.GetMainConnection()
	var like restaurantlikemodel.RestaurantLike
	
	if err := db.Where("restaurant_id = ? AND user_id = ?", restaurantId.String(), userId.String()).
		First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, datatype.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	
	return &like, nil
}
