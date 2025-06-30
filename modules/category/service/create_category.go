package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CreateCommand struct {
	DTO *categorymodel.CreateCategoryDTO
}

type CreateCommandHandler struct {
	catRepo ICreateRepo
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *categorymodel.Category) error
}

func NewCreateCommandHandler(catRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{catRepo: catRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*categorymodel.Category, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}
	
	categoryID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	
	category := categorymodel.Category{
		ID:          categoryID,
		Name:        cmd.DTO.Name,
		Description: cmd.DTO.Description,
		Status:      datatype.StatusActive,
	}
	
	if err = hdl.catRepo.Insert(ctx, &category); err != nil {
		return nil, err
	}
	
	return &category, nil
}
