package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

type CreateNewCommand struct {
	Dto categorymodel.Category
}

type CreateNewCommandHandler struct {
	catRepo ICategoryCommandRepo
}

func NewCreateNewCommandHandler(catRepo ICategoryCommandRepo) *CreateNewCommandHandler {
	return &CreateNewCommandHandler{catRepo: catRepo}
}

func (hdl *CreateNewCommandHandler) Execute(ctx context.Context, cmd *CreateNewCommand) (*uuid.UUID, error) {
	if err := cmd.Dto.Validate(); err != nil {
		return nil, err
	}
	
	cmd.Dto.ID, _ = uuid.NewV7()
	
	if err := hdl.catRepo.Insert(ctx, &cmd.Dto); err != nil {
		return nil, err
	}
	
	return &cmd.Dto.ID, nil
}
