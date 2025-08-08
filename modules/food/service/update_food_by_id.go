package foodservice

import (
	"context"
	"errors"
	"time"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type UpdateCommand struct {
	ID  string                    `json:"id"`
	DTO *foodmodel.UpdateFoodDTO `json:"data"`
}

type UpdateCommandHandler struct {
	foodRepo IUpdateRepo
}

type IUpdateRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*foodmodel.Food, error)
	Update(ctx context.Context, id uuid.UUID, data map[string]interface{}) error
}

func NewUpdateCommandHandler(foodRepo IUpdateRepo) *UpdateCommandHandler {
	return &UpdateCommandHandler{foodRepo: foodRepo}
}

func (cmd *UpdateCommand) Validate() error {
	if cmd.ID == "" {
		return errors.New("food id is required")
	}
	
	if _, err := uuid.Parse(cmd.ID); err != nil {
		return errors.New("invalid food id format")
	}
	
	return cmd.DTO.Validate()
}

func (hdl *UpdateCommandHandler) Execute(ctx context.Context, cmd *UpdateCommand) (*foodmodel.Food, error) {
	if err := cmd.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	foodUUID, _ := uuid.Parse(cmd.ID)
	
	// Check if food exists and not deleted
	existingFood, err := hdl.foodRepo.FindByID(ctx, foodUUID)
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return nil, datatype.ErrNotFound.WithError(foodmodel.ErrFoodNotFound.Error())
		}
		
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if existingFood.Status == foodmodel.FoodStatusDeleted {
		return nil, datatype.ErrNotFound.WithError(foodmodel.ErrFoodAlreadyDeleted.Error())
	}
	
	// Prepare update data
	updateData := make(map[string]interface{})
	updateData["updated_at"] = time.Now().UTC()
	
	if cmd.DTO.Name != nil {
		updateData["name"] = *cmd.DTO.Name
	}
	
	if cmd.DTO.Description != nil {
		updateData["description"] = *cmd.DTO.Description
	}
	
	if cmd.DTO.Price != nil {
		updateData["price"] = *cmd.DTO.Price
	}
	
	if cmd.DTO.CategoryID != nil {
		if *cmd.DTO.CategoryID == "" {
			updateData["category_id"] = nil
		} else {
			if categoryUUID, err := uuid.Parse(*cmd.DTO.CategoryID); err == nil {
				updateData["category_id"] = categoryUUID
			} else {
				return nil, datatype.ErrBadRequest.WithError("invalid category_id format")
			}
		}
	}
	
	// Update the food
	if err = hdl.foodRepo.Update(ctx, foodUUID, updateData); err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	// Fetch and return updated food
	updatedFood, err := hdl.foodRepo.FindByID(ctx, foodUUID)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return updatedFood, nil
}