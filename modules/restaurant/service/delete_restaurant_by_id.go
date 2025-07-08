package restaurantservice

import (
	"context"
	"errors"
	
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
		if errors.Is(err, datatype.ErrNotFound) {
			return datatype.ErrNotFound
		}
		
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if restaurant.Status == datatype.StatusDeleted {
		return datatype.ErrNotFound.WithError(restaurantmodel.ErrRestaurantAlreadyDeleted.Error())
	}
	
	if err = hdl.restaurantRepo.Delete(ctx, cmd.ID, false); err != nil {
		return datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	return nil
}
