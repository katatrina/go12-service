package restaurantservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateCommand struct {
	DTO *restaurantmodel.CreateRestaurantDTO
}

type CreateCommandHandler struct {
	restRepo ICreateRepo
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *restaurantmodel.CreateRestaurantDTO) error
}

func NewCreateCommandHandler(restRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{restRepo: restRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*uuid.UUID, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}
	
	restaurantID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	
	restaurant := restaurantmodel.CreateRestaurantDTO{
		ID:         restaurantID,
		OwnerID:    cmd.DTO.OwnerID,
		Name:       cmd.DTO.Name,
		Addr:       cmd.DTO.Addr,
		CityID:     cmd.DTO.CityID,
		CategoryID: cmd.DTO.CategoryID,
		Status:     datatype.StatusActive,
	}
	
	if err = hdl.restRepo.Insert(ctx, &restaurant); err != nil {
		return nil, err
	}
	
	return &restaurantID, nil
}
