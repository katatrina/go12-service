package restaurantlikemodel

import (
	"github.com/google/uuid"
)

type RestaurantLikeCreateDTO struct {
	RestaurantID uuid.UUID `json:"restaurant_id" gorm:"column:restaurant_id;primaryKey"`
	UserID       uuid.UUID `json:"user_id" gorm:"column:user_id;primaryKey"`
}

func (RestaurantLikeCreateDTO) TableName() string {
	return RestaurantLike{}.TableName()
}
