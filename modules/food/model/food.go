package foodmodel

import (
	"time"
	
	"github.com/google/uuid"
)

type FoodStatus string

const (
	FoodStatusPending  FoodStatus = "pending"
	FoodStatusActive   FoodStatus = "active"
	FoodStatusInactive FoodStatus = "inactive"
	FoodStatusDeleted  FoodStatus = "deleted"
)

type Food struct {
	ID           uuid.UUID  `json:"id" gorm:"column:id;"`
	RestaurantID uuid.UUID  `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryID   *uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	Name         string     `json:"name" gorm:"column:name;"`
	Description  *string    `json:"description" gorm:"column:description;"`
	Price        float64    `json:"price" gorm:"column:price;"`
	// Images       *string    `json:"images" gorm:"column:images;"` // TODO: Will be added later after discussing with mentor
	Status       FoodStatus `json:"status" gorm:"column:status;"`
	CreatedAt    time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodFilterDTO struct {
	RestaurantID *string  `json:"restaurant_id,omitempty" form:"restaurant_id"`
	CategoryID   *string  `json:"category_id,omitempty" form:"category_id"`
	MinPrice     *float64 `json:"min_price,omitempty" form:"min_price"`
	MaxPrice     *float64 `json:"max_price,omitempty" form:"max_price"`
	Search       *string  `json:"search,omitempty" form:"search"`
}