package foodservice

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateCommand struct {
	DTO *foodmodel.CreateFoodDTO
}

type CreateCommandHandler struct {
	foodRepo ICreateRepo
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *foodmodel.Food) error
}

func NewCreateCommandHandler(foodRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{foodRepo: foodRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*foodmodel.Food, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	foodID, err := uuid.NewV7()
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	restaurantUUID, err := uuid.Parse(cmd.DTO.RestaurantID)
	if err != nil {
		return nil, datatype.ErrBadRequest.WithError("invalid restaurant_id format")
	}
	
	var categoryUUID *uuid.UUID
	if cmd.DTO.CategoryID != nil && *cmd.DTO.CategoryID != "" {
		if parsedUUID, err := uuid.Parse(*cmd.DTO.CategoryID); err == nil {
			categoryUUID = &parsedUUID
		} else {
			return nil, datatype.ErrBadRequest.WithError("invalid category_id format")
		}
	}
	
	food := foodmodel.Food{
		ID:           foodID,
		RestaurantID: restaurantUUID,
		CategoryID:   categoryUUID,
		Name:         cmd.DTO.Name,
		Description:  cmd.DTO.Description,
		Price:        cmd.DTO.Price,
		Status:       foodmodel.FoodStatusActive,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	
	if err = hdl.foodRepo.Insert(ctx, &food); err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return &food, nil
}