package grpcclient

import (
	"context"
	"errors"
	foodservice "github.com/katatrina/go12-service/modules/food/service"
)

// RestaurantGRPCAdapter adapts RestaurantGRPCClient to foodservice.IRestaurantRPC interface
type RestaurantGRPCAdapter struct {
	client *RestaurantGRPCClient
}

func NewRestaurantGRPCAdapter(serviceURL string) *RestaurantGRPCAdapter {
	client := NewRestaurantGRPCClient(serviceURL)
	return &RestaurantGRPCAdapter{client: client}
}

func (a *RestaurantGRPCAdapter) GetRestaurantByID(ctx context.Context, restaurantID string) (*foodservice.RestaurantDTO, error) {
	restaurants, err := a.client.GetRestaurantsByIDs(ctx, []string{restaurantID})
	if err != nil {
		return nil, err
	}

	if len(restaurants) == 0 {
		return nil, errors.New("restaurant not found")
	}

	dto := restaurants[0]
	// Convert restaurant.RestaurantDTO to foodservice.RestaurantDTO
	return &foodservice.RestaurantDTO{
		ID:         dto.Id,
		Name:       dto.Name,
		Address:    dto.Address,
		CategoryID: &dto.CategoryId,
		Status:     dto.Status,
	}, nil
}