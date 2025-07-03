package restaurantmodel

import (
	"encoding/json"
	"strings"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateRestaurantDTO struct {
	OwnerID    uuid.UUID  `json:"ownerID"`
	Name       string     `json:"name"`
	Addr       string     `json:"addr"`
	CityID     *uuid.UUID `json:"cityID"`
	CategoryID *uuid.UUID `json:"categoryID"`
}

type UpdateRestaurantDTO struct {
	Name             *string          `json:"name"`
	Addr             *string          `json:"addr"`
	CityID           *string          `json:"cityID"`
	CategoryID       *string          `json:"categoryID"`
	Lat              *float64         `json:"lat"`
	Lng              *float64         `json:"lng"`
	Cover            *json.RawMessage `json:"cover"`
	Logo             *json.RawMessage `json:"logo"`
	ShippingFeePerKm *float64         `json:"shippingFeePerKm"`
	Status           *string          `json:"status"`
}

type FilterRestaurantDTO struct {
	Status     *string `json:"status" form:"status"`
	CityID     *string `json:"cityID" form:"cityID"`
	CategoryID *string `json:"categoryID" form:"categoryID"`
}

func (*CreateRestaurantDTO) TableName() string {
	return "restaurants"
}

func (dto *CreateRestaurantDTO) Validate() error {
	// Validate name
	dto.Name = strings.TrimSpace(dto.Name)
	if dto.Name == "" {
		return ErrNameRequired
	}
	if len(dto.Name) > 50 {
		return ErrInvalidNameLength
	}
	
	// Validate address
	dto.Addr = strings.TrimSpace(dto.Addr)
	if dto.Addr == "" {
		return ErrAddrRequired
	}
	
	return nil
}

func (dto *UpdateRestaurantDTO) Validate() error {
	if dto.Name != nil {
		*dto.Name = strings.TrimSpace(*dto.Name)
		if *dto.Name == "" {
			return ErrNameRequired
		}
		if len(*dto.Name) > 50 {
			return ErrInvalidNameLength
		}
	}
	if dto.Addr != nil {
		*dto.Addr = strings.TrimSpace(*dto.Addr)
		if *dto.Addr == "" {
			return ErrAddrRequired
		}
	}
	if dto.Status != nil {
		*dto.Status = strings.TrimSpace(*dto.Status)
		status := datatype.Status(strings.ToLower(*dto.Status))
		if !status.Valid() {
			return ErrStatusInvalid
		}
	}
	return nil
}

func (dto *FilterRestaurantDTO) Validate() error {
	if dto.Status != nil {
		*dto.Status = strings.TrimSpace(*dto.Status)
		status := datatype.Status(strings.ToLower(*dto.Status))
		if !status.Valid() {
			return ErrStatusInvalid
		}
	}
	return nil
}
