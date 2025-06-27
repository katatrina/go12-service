package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

type IListRepo interface {
	List(
		ctx context.Context,
		pagingDTO *sharedmodel.PagingDTO,
		filterDTO *restaurantmodel.FilterRestaurantDTO,
	) ([]restaurantmodel.Restaurant, error)
}

type ICategoryRepo interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Category, error)
}

type ListRestaurantsQueryHandler struct {
	restRepo IListRepo
	catRepo  ICategoryRepo
}

func NewListRestaurantsQueryHandler(restRepo IListRepo, catRepo ICategoryRepo) *ListRestaurantsQueryHandler {
	return &ListRestaurantsQueryHandler{
		restRepo: restRepo,
		catRepo:  catRepo,
	}
}

type ListQuery struct {
	restaurantmodel.FilterRestaurantDTO
	sharedmodel.PagingDTO
}

type ListRestaurantItemDTO struct {
	restaurantmodel.Restaurant
	Category *restaurantmodel.Category `json:"category"`
}

func (hdl *ListRestaurantsQueryHandler) Execute(
	ctx context.Context,
	query *ListQuery,
) ([]ListRestaurantItemDTO, error) {
	restaurants, err := hdl.restRepo.List(ctx, &query.PagingDTO, &query.FilterRestaurantDTO)
	if err != nil {
		return nil, err
	}
	
	categoryIDs := make([]uuid.UUID, len(restaurants))
	
	for i := range restaurants {
		categoryIDs[i] = restaurants[i].CategoryID
	}
	
	categories, err := hdl.catRepo.FindByIDs(ctx, categoryIDs)
	if err != nil {
		// return restaurants, nil (If category is not important, we can return restaurants without categories)
		return nil, err
	}
	
	mapCatIDToCategory := make(map[uuid.UUID]restaurantmodel.Category, len(restaurants))
	
	for i := range categories {
		mapCatIDToCategory[categories[i].ID] = categories[i] // Better to get by index
	}
	
	result := make([]ListRestaurantItemDTO, len(restaurants))
	for i := range restaurants {
		cat := mapCatIDToCategory[restaurants[i].CategoryID]
		result[i] = ListRestaurantItemDTO{
			Restaurant: restaurants[i],
			Category:   &cat,
		}
	}
	
	return result, nil
}
