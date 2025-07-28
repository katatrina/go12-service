package restaurantlikemodel

import (
	"github.com/google/uuid"
)

type RestaurantLikeCreateDTO struct {
	RestaurantId uuid.UUID `json:"restaurantId" gorm:"column:restaurant_id;primaryKey"`
	UserId       uuid.UUID `json:"userId" gorm:"column:user_id;primaryKey"`
}

func (RestaurantLikeCreateDTO) TableName() string {
	return RestaurantLike{}.TableName()
}
