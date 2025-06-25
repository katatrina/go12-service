package service

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IDeleteByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.Category, error)
	Delete(ctx context.Context, id uuid.UUID, isHard bool) error
}

type DeleteByIDCommandHandler struct {
	catRepo IDeleteByIDRepo
}

func NewDeleteByIDCommandHandler(catRepo IDeleteByIDRepo) *DeleteByIDCommandHandler {
	return &DeleteByIDCommandHandler{
		catRepo: catRepo,
	}
}

type DeleteByIDCommand struct {
	ID uuid.UUID
}

func (hdl *DeleteByIDCommandHandler) Execute(ctx context.Context, cmd *DeleteByIDCommand) error {
	category, err := hdl.catRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	if category.Status == datatype.StatusDeleted {
		return nil
	}
	
	if err = hdl.catRepo.Delete(ctx, cmd.ID, false); err != nil {
		return err
	}
	
	return nil
}
