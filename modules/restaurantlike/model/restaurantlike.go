package restaurantlikemodel

import (
	"time"
	
	"github.com/google/uuid"
)

type RestaurantLike struct {
	RestaurantID uuid.UUID  `json:"restaurant_id" gorm:"column:restaurant_id;primaryKey"`
	UserID       uuid.UUID  `json:"user_id" gorm:"column:user_id;primaryKey"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (rl RestaurantLike) ToData() map[string]interface{} {
	return map[string]interface{}{
		"restaurant_id": rl.RestaurantID,
		"user_id":       rl.UserID,
		"created_at":    rl.CreatedAt,
	}
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
