package restaurantlikemodel

import (
	"time"
	
	"github.com/google/uuid"
)

type RestaurantLike struct {
	RestaurantID uuid.UUID  `json:"restaurantID" gorm:"column:restaurant_id;primaryKey"`
	UserID       uuid.UUID  `json:"userID" gorm:"column:user_id;primaryKey"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;"`
}

func (rl RestaurantLike) ToData() map[string]interface{} {
	return map[string]interface{}{
		"restaurantID": rl.RestaurantID,
		"userID":       rl.UserID,
		"createdAt":    rl.CreatedAt,
	}
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
