package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
	"go.opentelemetry.io/otel"
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
	dbSpanCtx, dbSpan := otel.Tracer("go12-service").Start(ctx, "list-restaurants")
	
	restaurants, err := hdl.restRepo.List(dbSpanCtx, &query.PagingDTO, &query.FilterRestaurantDTO)
	
	dbSpan.AddEvent("Get list of restaurants")
	dbSpan.End()
	
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	// Collect unique category IDs (skip nil values)
	categoryIDSet := make(map[uuid.UUID]bool)
	for i := range restaurants {
		if restaurants[i].CategoryID != nil {
			categoryIDSet[*restaurants[i].CategoryID] = true
		}
	}
	
	// Convert to slice for gRPC call
	categoryIDs := make([]uuid.UUID, 0, len(categoryIDSet))
	for id := range categoryIDSet {
		categoryIDs = append(categoryIDs, id)
	}
	
	_, catRepoSpan := otel.Tracer("go12-service").Start(ctx, "cat-repo")
	categories, err := hdl.catRepo.FindByIDs(ctx, categoryIDs)
	catRepoSpan.End()
	if err != nil {
		// return restaurants, nil (If category is not important, we can return restaurants without categories)
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	mapCatIDToCategory := make(map[uuid.UUID]restaurantmodel.Category, len(restaurants))
	
	for i := range categories {
		mapCatIDToCategory[categories[i].ID] = categories[i] // Better to get by index
	}
	
	result := make([]ListRestaurantItemDTO, len(restaurants))
	for i := range restaurants {
		if restaurants[i].CategoryID == nil {
			result[i] = ListRestaurantItemDTO{
				Restaurant: restaurants[i],
			}
			continue
		}
		
		if category, ok := mapCatIDToCategory[*restaurants[i].CategoryID]; ok {
			result[i] = ListRestaurantItemDTO{
				Restaurant: restaurants[i],
				Category:   &category,
			}
			continue
		}
		
		result[i] = ListRestaurantItemDTO{
			Restaurant: restaurants[i],
		}
	}
	
	return result, nil
}
