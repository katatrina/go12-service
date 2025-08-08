package foodservice

import (
	"context"
	
	foodmodel "github.com/katatrina/go12-service/modules/food/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type ListCommand struct {
	DTO *foodmodel.FoodListDTO
}

type ListCommandHandler struct {
	foodRepo IListRepo
}

type IListRepo interface {
	List(ctx context.Context, filter *foodmodel.FoodFilterDTO, offset, limit int) ([]*foodmodel.Food, error)
	Count(ctx context.Context, filter *foodmodel.FoodFilterDTO) (int64, error)
}

func NewListCommandHandler(foodRepo IListRepo) *ListCommandHandler {
	return &ListCommandHandler{foodRepo: foodRepo}
}

func (hdl *ListCommandHandler) Execute(ctx context.Context, cmd *ListCommand) (*foodmodel.FoodListResponseDTO, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, datatype.ErrBadRequest.WithError(err.Error())
	}
	
	offset := (cmd.DTO.Page - 1) * cmd.DTO.Limit
	
	filter := &foodmodel.FoodFilterDTO{
		RestaurantID: cmd.DTO.RestaurantID,
		CategoryID:   cmd.DTO.CategoryID,
		MinPrice:     cmd.DTO.MinPrice,
		MaxPrice:     cmd.DTO.MaxPrice,
		Search:       cmd.DTO.Search,
	}
	
	foods, err := hdl.foodRepo.List(ctx, filter, offset, cmd.DTO.Limit)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	total, err := hdl.foodRepo.Count(ctx, filter)
	if err != nil {
		return nil, datatype.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}
	
	response := make([]*foodmodel.FoodResponseDTO, len(foods))
	for i, food := range foods {
		response[i] = foodmodel.NewFoodResponseDTO(food)
	}
	
	return &foodmodel.FoodListResponseDTO{
		Data:  response,
		Page:  cmd.DTO.Page,
		Limit: cmd.DTO.Limit,
		Total: total,
	}, nil
}