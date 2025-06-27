package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateCommand struct {
	DTO *restaurantmodel.CreateRestaurantDTO
}

type CreateCommandHandler struct {
	restaurantRepo ICreateRepo
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *restaurantmodel.Restaurant) error
}

func NewCreateCommandHandler(restaurantRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{restaurantRepo: restaurantRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*restaurantmodel.Restaurant, error) {
	// TODO: validate DTO nếu cần
	// if err := cmd.DTO.Validate(); err != nil {
	// 	return nil, err
	// }
	
	restaurantID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	
	ownerID, err := uuid.Parse(cmd.DTO.OwnerID)
	if err != nil {
		return nil, err
	}
	
	categoryID, err := uuid.Parse(cmd.DTO.CategoryID)
	if err != nil {
		return nil, err
	}
	
	var cityID *uuid.UUID
	if cmd.DTO.CityID != nil {
		parsed, err := uuid.Parse(*cmd.DTO.CityID)
		if err != nil {
			return nil, err
		}
		cityID = &parsed
	}
	
	restaurant := restaurantmodel.Restaurant{
		ID:               restaurantID,
		OwnerID:          ownerID,
		CategoryID:       categoryID,
		Name:             cmd.DTO.Name,
		Addr:             cmd.DTO.Addr,
		CityID:           cityID,
		Lat:              cmd.DTO.Lat,
		Lng:              cmd.DTO.Lng,
		ShippingFeePerKm: cmd.DTO.ShippingFeePerKm,
		Status:           datatype.StatusActive,
	}
	
	if err = hdl.restaurantRepo.Insert(ctx, &restaurant); err != nil {
		return nil, err
	}
	
	return &restaurant, nil
}
