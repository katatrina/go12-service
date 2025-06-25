package service

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IUpdateByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.Restaurant, error)
	Update(ctx context.Context, id uuid.UUID, dto *model.UpdateRestaurantDTO) error
}

type UpdateByIDCommandHandler struct {
	restaurantRepo IUpdateByIDRepo
}

type UpdateByIDCommand struct {
	ID  uuid.UUID
	DTO *model.UpdateRestaurantDTO
}

func NewUpdateByIDCommandHandler(restaurantRepo IUpdateByIDRepo) *UpdateByIDCommandHandler {
	return &UpdateByIDCommandHandler{
		restaurantRepo: restaurantRepo,
	}
}

func (hdl *UpdateByIDCommandHandler) Execute(ctx context.Context, cmd *UpdateByIDCommand) error {
	if err := cmd.DTO.Validate(); err != nil {
		return err
	}
	
	restaurant, err := hdl.restaurantRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	if restaurant.Status == datatype.StatusDeleted {
		return model.ErrRestaurantAlreadyDeleted
	}
	
	if err = hdl.restaurantRepo.Update(ctx, cmd.ID, cmd.DTO); err != nil {
		return err
	}
	
	return nil
}
