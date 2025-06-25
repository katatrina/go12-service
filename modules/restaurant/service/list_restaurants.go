package service

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/restaurant/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

type IListRepo interface {
	ListRestaurants(
		ctx context.Context,
		pagingDTO *sharedmodel.PagingDTO,
		filterDTO *model.FilterRestaurantDTO,
	) ([]model.Restaurant, error)
}

type ListRestaurantsQueryHandler struct {
	restaurantRepo IListRepo
}

func NewListRestaurantsQueryHandler(restaurantRepo IListRepo) *ListRestaurantsQueryHandler {
	return &ListRestaurantsQueryHandler{
		restaurantRepo: restaurantRepo,
	}
}

type ListQuery struct {
	model.FilterRestaurantDTO
	sharedmodel.PagingDTO
}

func (hdl *ListRestaurantsQueryHandler) Execute(
	ctx context.Context,
	query *ListQuery,
) ([]model.Restaurant, error) {
	restaurants, err := hdl.restaurantRepo.ListRestaurants(ctx, &query.PagingDTO, &query.FilterRestaurantDTO)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}
