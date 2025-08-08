package restaurantgrpcctl

import (
	"context"
	"log"

	"github.com/katatrina/go12-service/gen/proto/restaurant"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"

	"github.com/google/uuid"
)

type RestaurantRepository interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Restaurant, error)
	FindByCategoryID(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]restaurantmodel.Restaurant, error)
}

type RestaurantGrpcServer struct {
	restaurant.UnimplementedRestaurantServer
	repo RestaurantRepository
}

func NewRestaurantGrpcServer(repo RestaurantRepository) *RestaurantGrpcServer {
	return &RestaurantGrpcServer{
		repo: repo,
	}
}

func (s *RestaurantGrpcServer) GetRestaurantsByIDs(
	ctx context.Context,
	req *restaurant.GetRestaurantIDsRequest,
) (*restaurant.RestaurantIDsResp, error) {
	log.Println("GetRestaurantsByIDs by gRPC")

	uuidIds := make([]uuid.UUID, 0, len(req.Ids))

	for _, id := range req.Ids {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			log.Printf("Invalid UUID: %s, error: %v", id, err)
			continue
		}
		uuidIds = append(uuidIds, parsedId)
	}

	restaurants, err := s.repo.FindByIDs(ctx, uuidIds)
	if err != nil {
		log.Printf("Error finding restaurants: %v", err)
		return nil, err
	}

	result := make([]*restaurant.RestaurantDTO, 0, len(restaurants))

	for _, r := range restaurants {
		categoryID := ""
		if r.CategoryID != nil {
			categoryID = r.CategoryID.String()
		}

		result = append(result, &restaurant.RestaurantDTO{
			Id:         r.ID.String(),
			Name:       r.Name,
			Address:    r.Addr,
			CategoryId: categoryID,
			Phone:      "", // Add phone field to model if needed
			Status:     string(r.Status),
			LikedCount: 0, // Add liked_count field to model if needed
			CreatedAt:  r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:  r.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &restaurant.RestaurantIDsResp{Data: result}, nil
}

func (s *RestaurantGrpcServer) GetRestaurantsByCategoryID(
	ctx context.Context,
	req *restaurant.GetRestaurantsByCategoryRequest,
) (*restaurant.RestaurantIDsResp, error) {
	log.Println("GetRestaurantsByCategoryID by gRPC")

	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		log.Printf("Invalid category UUID: %s, error: %v", req.CategoryId, err)
		return nil, err
	}

	limit := 10
	if req.Limit != nil {
		limit = int(*req.Limit)
	}

	offset := 0
	if req.Offset != nil {
		offset = int(*req.Offset)
	}

	restaurants, err := s.repo.FindByCategoryID(ctx, categoryID, limit, offset)
	if err != nil {
		log.Printf("Error finding restaurants by category: %v", err)
		return nil, err
	}

	result := make([]*restaurant.RestaurantDTO, 0, len(restaurants))

	for _, r := range restaurants {
		categoryIDStr := ""
		if r.CategoryID != nil {
			categoryIDStr = r.CategoryID.String()
		}

		result = append(result, &restaurant.RestaurantDTO{
			Id:         r.ID.String(),
			Name:       r.Name,
			Address:    r.Addr,
			CategoryId: categoryIDStr,
			Phone:      "", // Add phone field to model if needed
			Status:     string(r.Status),
			LikedCount: 0, // Add liked_count field to model if needed
			CreatedAt:  r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:  r.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &restaurant.RestaurantIDsResp{Data: result}, nil
}