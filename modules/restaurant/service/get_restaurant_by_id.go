package restaurantservice

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IGetByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*restaurantmodel.Restaurant, error)
}

type GetByIDQueryHandler struct {
	restaurantRepo IGetByIDRepo
}

type GetByIDQuery struct {
	ID uuid.UUID
}

func NewGetDetailQueryHandler(restaurantRepo IGetByIDRepo) *GetByIDQueryHandler {
	return &GetByIDQueryHandler{
		restaurantRepo: restaurantRepo,
	}
}

func (hdl *GetByIDQueryHandler) Execute(ctx context.Context, query *GetByIDQuery) (*restaurantmodel.Restaurant, error) {
	restaurant, err := hdl.restaurantRepo.FindByID(ctx, query.ID)
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return nil, datatype.ErrNotFound
		}
		
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if restaurant.Status == datatype.StatusDeleted {
		return nil, datatype.ErrDeleted.WithError(restaurantmodel.ErrRestaurantAlreadyDeleted.Error())
	}
	
	return restaurant, nil
}
