package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IUpdateByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	Update(ctx context.Context, id uuid.UUID, dto *categorymodel.UpdateCategoryDTO) error
}

type UpdateByIDCommandHandler struct {
	catRepo IUpdateByIDRepo
}

type UpdateByIDCommand struct {
	ID  uuid.UUID
	DTO *categorymodel.UpdateCategoryDTO
}

func NewUpdateByIDCommandHandler(catRepo IUpdateByIDRepo) *UpdateByIDCommandHandler {
	return &UpdateByIDCommandHandler{
		catRepo: catRepo,
	}
}

func (hdl *UpdateByIDCommandHandler) Execute(ctx context.Context, cmd *UpdateByIDCommand) error {
	if err := cmd.DTO.Validate(); err != nil {
		return err
	}
	
	category, err := hdl.catRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	
	if category.Status == datatype.StatusDeleted {
		return categorymodel.ErrCategoryAlreadyDeleted
	}
	
	if err = hdl.catRepo.Update(ctx, cmd.ID, cmd.DTO); err != nil {
		return err
	}
	
	return nil
}
