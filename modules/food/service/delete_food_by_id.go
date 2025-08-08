package foodservice

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type DeleteCommand struct {
	ID string `json:"id"`
}

type DeleteCommandHandler struct {
	foodRepo IDeleteRepo
}

type IDeleteRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*foodmodel.Food, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewDeleteCommandHandler(foodRepo IDeleteRepo) *DeleteCommandHandler {
	return &DeleteCommandHandler{foodRepo: foodRepo}
}

func (cmd *DeleteCommand) Validate() error {
	if cmd.ID == "" {
		return errors.New("food id is required")
	}
	
	if _, err := uuid.Parse(cmd.ID); err != nil {
		return errors.New("invalid food id format")
	}
	
	return nil
}

func (hdl *DeleteCommandHandler) Execute(ctx context.Context, cmd *DeleteCommand) error {
	if err := cmd.Validate(); err != nil {
		return datatype.ErrBadRequest.WithError(err.Error())
	}
	
	foodUUID, _ := uuid.Parse(cmd.ID)
	
	// Check if food exists and not already deleted
	existingFood, err := hdl.foodRepo.FindByID(ctx, foodUUID)
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return datatype.ErrNotFound.WithError(foodmodel.ErrFoodNotFound.Error())
		}
		
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if existingFood.Status == foodmodel.FoodStatusDeleted {
		return datatype.ErrBadRequest.WithError(foodmodel.ErrFoodAlreadyDeleted.Error())
	}
	
	// Soft delete the food
	if err = hdl.foodRepo.Delete(ctx, foodUUID); err != nil {
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return nil
}