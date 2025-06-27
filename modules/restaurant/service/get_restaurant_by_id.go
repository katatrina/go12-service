package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
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
		return nil, err
	}
	return restaurant, nil
}
