package model

import (
	"encoding/json"
	"strings"

	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateRestaurantDTO struct {
	OwnerID    string   `json:"ownerID"`
	Name       string   `json:"name"`
	Addr       string   `json:"addr"`
	CityID     *string  `json:"cityID"`
	CategoryID *string  `json:"categoryID"`
	Lat        *float64 `json:"lat"`
	Lng        *float64 `json:"lng"`
	// Cover            interface{} `json:"cover"`
	// Logo             interface{} `json:"logo"`
	ShippingFeePerKm float64 `json:"shippingFeePerKm"`
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
