package foodservice

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/gen/proto/category"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
	"go.opentelemetry.io/otel"
)

type CreateCommand struct {
	DTO *foodmodel.CreateFoodDTO
}

type CreateCommandHandler struct {
	foodRepo      ICreateRepo
	categoryRPC   ICreateCategoryRPC
	restaurantRPC ICreateRestaurantRPC
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *foodmodel.Food) error
}

type ICreateCategoryRPC interface {
	GetCategoriesByIDs(ctx context.Context, ids []string) ([]*category.CategoryDTO, error)
}

type ICreateRestaurantRPC interface {
	GetRestaurantByIDForCreate(ctx context.Context, restaurantID string) (*CreateRestaurantDTO, error)
}

type CreateRestaurantDTO struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	CategoryID *string `json:"category_id"`
	Status     string  `json:"status"`
}

func NewCreateCommandHandler(foodRepo ICreateRepo, categoryRPC ICreateCategoryRPC, restaurantRPC ICreateRestaurantRPC) *CreateCommandHandler {
	return &CreateCommandHandler{
		foodRepo:      foodRepo,
		categoryRPC:   categoryRPC,
		restaurantRPC: restaurantRPC,
	}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*foodmodel.Food, error) {
	ctx, span := otel.Tracer("go12-service").Start(ctx, "food-service.create")
	defer span.End()
	
	if err := cmd.DTO.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	foodID, err := uuid.NewV7()
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	restaurantUUID, err := uuid.Parse(cmd.DTO.RestaurantID)
	if err != nil {
		return nil, datatype.ErrBadRequest.WithError("invalid restaurant_id format")
	}
	
	// Validate restaurant exists via gRPC
	restCtx, restSpan := otel.Tracer("go12-service").Start(ctx, "restaurant-grpc.validate-exists")
	_, err = hdl.restaurantRPC.GetRestaurantByIDForCreate(restCtx, cmd.DTO.RestaurantID)
	restSpan.End()
	if err != nil {
		return nil, datatype.ErrBadRequest.WithError("restaurant_id does not exist or is not accessible")
	}
	
	var categoryUUID *uuid.UUID
	if cmd.DTO.CategoryID != nil && *cmd.DTO.CategoryID != "" {
		if parsedUUID, err := uuid.Parse(*cmd.DTO.CategoryID); err == nil {
			categoryUUID = &parsedUUID
			
			// Validate category exists via gRPC
			catCtx, catSpan := otel.Tracer("go12-service").Start(ctx, "category-grpc.validate-exists")
			categories, err := hdl.categoryRPC.GetCategoriesByIDs(catCtx, []string{*cmd.DTO.CategoryID})
			catSpan.End()
			if err != nil || len(categories) == 0 {
				return nil, datatype.ErrBadRequest.WithError("category_id does not exist or is not accessible")
			}
		} else {
			return nil, datatype.ErrBadRequest.WithError("invalid category_id format")
		}
	}
	
	food := foodmodel.Food{
		ID:           foodID,
		RestaurantID: restaurantUUID,
		CategoryID:   categoryUUID,
		Name:         cmd.DTO.Name,
		Description:  cmd.DTO.Description,
		Price:        cmd.DTO.Price,
		Status:       foodmodel.FoodStatusActive,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	
	// Database insert with tracing
	insertCtx, insertSpan := otel.Tracer("go12-service").Start(ctx, "food-repo.insert")
	if err = hdl.foodRepo.Insert(insertCtx, &food); err != nil {
		insertSpan.End()
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	insertSpan.End()
	
	return &food, nil
}
