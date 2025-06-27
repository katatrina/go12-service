package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IDeleteByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id uuid.UUID, isHard bool) error
}

type DeleteByIDCommandHandler struct {
	restaurantRepo IDeleteByIDRepo
}

func NewDeleteByIDCommandHandler(restaurantRepo IDeleteByIDRepo) *DeleteByIDCommandHandler {
	return &DeleteByIDCommandHandler{
		restaurantRepo: restaurantRepo,
	}
}

type DeleteByIDCommand struct {
	ID uuid.UUID
}

func (hdl *DeleteByIDCommandHandler) Execute(ctx context.Context, cmd *DeleteByIDCommand) error {
	restaurant, err := hdl.restaurantRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	if restaurant.Status == datatype.StatusDeleted {
		return nil
	}
	
	if err = hdl.restaurantRepo.Delete(ctx, cmd.ID, false); err != nil {
		return err
	}
	
	return nil
}
