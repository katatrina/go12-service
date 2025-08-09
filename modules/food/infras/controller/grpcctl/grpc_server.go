package foodgrpcctl

import (
	"context"
	"log"

	"github.com/katatrina/go12-service/gen/proto/food"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type FoodRepository interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]foodmodel.Food, error)
	FindByRestaurantID(ctx context.Context, restaurantID uuid.UUID, limit, offset int) ([]foodmodel.Food, error)
	FindByCategoryID(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]foodmodel.Food, error)
}

type FoodGrpcServer struct {
	food.UnimplementedFoodServer
	repo FoodRepository
}

func NewFoodGrpcServer(repo FoodRepository) *FoodGrpcServer {
	return &FoodGrpcServer{
		repo: repo,
	}
}

func (s *FoodGrpcServer) GetFoodsByIDs(
	ctx context.Context,
	req *food.GetFoodIDsRequest,
) (*food.FoodIDsResp, error) {
	ctx, span := otel.Tracer("go12-service").Start(ctx, "food-grpc.get-by-ids")
	defer span.End()
	
	log.Println("GetFoodsByIDs by gRPC")

	uuidIds := make([]uuid.UUID, len(req.Ids))

	for i, id := range req.Ids {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			log.Printf("Invalid UUID: %s, error: %v", id, err)
			continue
		}
		uuidIds[i] = parsedId
	}

	foods, err := s.repo.FindByIDs(ctx, uuidIds)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*food.FoodDTO, len(foods))

	for i, f := range foods {
		categoryID := ""
		if f.CategoryID != nil {
			categoryID = f.CategoryID.String()
		}

		description := ""
		if f.Description != nil {
			description = *f.Description
		}

		result[i] = &food.FoodDTO{
			Id:           f.ID.String(),
			RestaurantId: f.RestaurantID.String(),
			CategoryId:   categoryID,
			Name:         f.Name,
			Description:  description,
			Price:        f.Price,
			Status:       string(f.Status),
			CreatedAt:    f.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:    f.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return &food.FoodIDsResp{Data: result}, nil
}

func (s *FoodGrpcServer) GetFoodsByRestaurantID(
	ctx context.Context,
	req *food.GetFoodsByRestaurantRequest,
) (*food.FoodIDsResp, error) {
	ctx, span := otel.Tracer("go12-service").Start(ctx, "food-grpc.get-by-restaurant")
	defer span.End()
	
	log.Println("GetFoodsByRestaurantID by gRPC")

	restaurantID, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		log.Printf("Invalid restaurant UUID: %s, error: %v", req.RestaurantId, err)
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

	foods, err := s.repo.FindByRestaurantID(ctx, restaurantID, limit, offset)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*food.FoodDTO, len(foods))

	for i, f := range foods {
		categoryID := ""
		if f.CategoryID != nil {
			categoryID = f.CategoryID.String()
		}

		description := ""
		if f.Description != nil {
			description = *f.Description
		}

		result[i] = &food.FoodDTO{
			Id:           f.ID.String(),
			RestaurantId: f.RestaurantID.String(),
			CategoryId:   categoryID,
			Name:         f.Name,
			Description:  description,
			Price:        f.Price,
			Status:       string(f.Status),
			CreatedAt:    f.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:    f.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return &food.FoodIDsResp{Data: result}, nil
}

func (s *FoodGrpcServer) GetFoodsByCategoryID(
	ctx context.Context,
	req *food.GetFoodsByCategoryRequest,
) (*food.FoodIDsResp, error) {
	ctx, span := otel.Tracer("go12-service").Start(ctx, "food-grpc.get-by-category")
	defer span.End()
	
	log.Println("GetFoodsByCategoryID by gRPC")

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

	foods, err := s.repo.FindByCategoryID(ctx, categoryID, limit, offset)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*food.FoodDTO, len(foods))

	for i, f := range foods {
		categoryID := ""
		if f.CategoryID != nil {
			categoryID = f.CategoryID.String()
		}

		description := ""
		if f.Description != nil {
			description = *f.Description
		}

		result[i] = &food.FoodDTO{
			Id:           f.ID.String(),
			RestaurantId: f.RestaurantID.String(),
			CategoryId:   categoryID,
			Name:         f.Name,
			Description:  description,
			Price:        f.Price,
			Status:       string(f.Status),
			CreatedAt:    f.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:    f.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return &food.FoodIDsResp{Data: result}, nil
}