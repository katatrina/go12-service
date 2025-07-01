package rpcclient

import (
	"context"
	"fmt"
	
	"github.com/google/uuid"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"resty.dev/v3"
)

type CategoryRPCClient struct {
	catServiceURL string
}

func NewCategoryRPCClient(catServiceURL string) *CategoryRPCClient {
	return &CategoryRPCClient{
		catServiceURL: catServiceURL,
	}
}

func (c *CategoryRPCClient) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Category, error) {
	client := resty.New()
	
	type ResponseDTO struct {
		Data []restaurantmodel.Category `json:"data"`
	}
	
	var response ResponseDTO
	
	url := fmt.Sprintf("%s/find-by-ids", c.catServiceURL)
	
	_, err := client.R().
		SetBody(map[string]interface{}{
			"ids": ids,
		}).
		SetResult(&response).
		Post(url)
	if err != nil {
		return nil, err
	}
	
	return response.Data, nil
}
