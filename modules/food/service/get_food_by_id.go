package foodservice

import (
	"context"
	"errors"
	"log"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/gen/proto/category"
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type GetByIDCommand struct {
	ID string `json:"id"`
}

type GetByIDCommandHandler struct {
	foodRepo      IGetByIDRepo
	categoryRPC   ICategoryRPC
	restaurantRPC IRestaurantRPC
}

type IGetByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*foodmodel.Food, error)
}

type ICategoryRPC interface {
	GetCategoriesByIDs(ctx context.Context, ids []string) ([]*category.CategoryDTO, error)
}

type IRestaurantRPC interface {
	GetRestaurantByID(ctx context.Context, restaurantID string) (*RestaurantDTO, error)
}

// RestaurantDTO - duplicate definition for interface
type RestaurantDTO struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	CategoryID *string `json:"category_id"`
	Status     string  `json:"status"`
}

func NewGetByIDCommandHandler(foodRepo IGetByIDRepo, categoryRPC ICategoryRPC, restaurantRPC IRestaurantRPC) *GetByIDCommandHandler {
	return &GetByIDCommandHandler{
		foodRepo:      foodRepo,
		categoryRPC:   categoryRPC,
		restaurantRPC: restaurantRPC,
	}
}

func (cmd *GetByIDCommand) Validate() error {
	if cmd.ID == "" {
		return errors.New("food id is required")
	}
	
	if _, err := uuid.Parse(cmd.ID); err != nil {
		return errors.New("invalid food id format")
	}
	
	return nil
}

func (hdl *GetByIDCommandHandler) Execute(ctx context.Context, cmd *GetByIDCommand) (*foodmodel.FoodResponseDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	foodUUID, _ := uuid.Parse(cmd.ID)
	
	food, err := hdl.foodRepo.FindByID(ctx, foodUUID)
	if err != nil {
		if errors.Is(err, datatype.ErrRecordNotFound) {
			return nil, datatype.ErrNotFound.WithError(foodmodel.ErrFoodNotFound.Error())
		}
		
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	if food.Status == foodmodel.FoodStatusDeleted {
		return nil, datatype.ErrNotFound.WithError(foodmodel.ErrFoodAlreadyDeleted.Error())
	}
	
	// Create response DTO
	responseDTO := foodmodel.NewFoodResponseDTO(food)
	
	// Fetch category information if category_id exists
	if food.CategoryID != nil {
		categories, err := hdl.categoryRPC.GetCategoriesByIDs(ctx, []string{food.CategoryID.String()})
		if err != nil {
			log.Printf("Warning: Failed to fetch category info: %v", err)
		} else if len(categories) > 0 {
			categoryInfo := &foodmodel.CategoryInfo{
				ID:     categories[0].Id,
				Name:   categories[0].Name,
				Status: categories[0].Status,
			}
			responseDTO.WithCategory(categoryInfo)
		}
	}
	
	// Fetch restaurant information
	restaurant, err := hdl.restaurantRPC.GetRestaurantByID(ctx, food.RestaurantID.String())
	if err != nil {
		log.Printf("Warning: Failed to fetch restaurant info: %v", err)
	} else {
		restaurantInfo := &foodmodel.RestaurantInfo{
			ID:         restaurant.ID,
			Name:       restaurant.Name,
			Address:    restaurant.Address,
			CategoryID: restaurant.CategoryID,
			Status:     restaurant.Status,
		}
		responseDTO.WithRestaurant(restaurantInfo)
	}
	
	return responseDTO, nil
}