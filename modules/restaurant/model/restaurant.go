package restaurantmodel

import (
	"encoding/json"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type Restaurant struct {
	ID               uuid.UUID        `json:"id" gorm:"column:id"`
	OwnerID          uuid.UUID        `json:"owner_id" gorm:"column:owner_id"`
	CategoryID       *uuid.UUID       `json:"category_id" gorm:"column:category_id"`
	Name             string           `json:"name" gorm:"column:name"`
	Addr             string           `json:"addr" gorm:"column:addr"`
	CityID           *uuid.UUID       `json:"city_id" gorm:"column:city_id"`
	Lat              *float64         `json:"lat" gorm:"column:lat"`
	Lng              *float64         `json:"lng" gorm:"column:lng"`
	Cover            *json.RawMessage `json:"cover" gorm:"column:cover"`
	Logo             *json.RawMessage `json:"logo" gorm:"column:logo"`
	ShippingFeePerKm float64          `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km"`
	Status           datatype.Status  `json:"status" gorm:"column:status"`
	CreatedAt        time.Time        `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time        `json:"updated_at" gorm:"column:updated_at"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
