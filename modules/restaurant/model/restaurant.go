package restaurantmodel

import (
	"encoding/json"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type Restaurant struct {
	ID               uuid.UUID        `json:"id" gorm:"column:id"`
	OwnerID          uuid.UUID        `json:"ownerID" gorm:"column:owner_id"`
	CategoryID       uuid.UUID        `json:"categoryID" gorm:"column:category_id"`
	Name             string           `json:"name" gorm:"column:name"`
	Addr             string           `json:"addr" gorm:"column:addr"`
	CityID           *uuid.UUID       `json:"cityID" gorm:"column:city_id"`
	Lat              *float64         `json:"lat" gorm:"column:lat"`
	Lng              *float64         `json:"lng" gorm:"column:lng"`
	Cover            *json.RawMessage `json:"cover" gorm:"column:cover"`
	Logo             *json.RawMessage `json:"logo" gorm:"column:logo"`
	ShippingFeePerKm float64          `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km"`
	Status           datatype.Status  `json:"status" gorm:"column:status"`
	CreatedAt        time.Time        `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt        time.Time        `json:"updatedAt" gorm:"column:updated_at"`
	// Category 	 *categorymodel.Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	// Category      *Category        		 `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

func (r *Restaurant) TableName() string {
	return "restaurants"
}
