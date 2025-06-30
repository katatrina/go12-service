package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IGetByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
}

type GetByIDQueryHandler struct {
	catRepo IGetByIDRepo
}

func NewGetDetailQueryHandler(catRepo IGetByIDRepo) *GetByIDQueryHandler {
	return &GetByIDQueryHandler{
		catRepo: catRepo,
	}
}

type GetByIDQuery struct {
	ID uuid.UUID
}

func (hdl *GetByIDQueryHandler) Execute(ctx context.Context, query *GetByIDQuery) (*categorymodel.Category, error) {
	category, err := hdl.catRepo.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	
	if category.Status == datatype.StatusDeleted {
		return nil, categorymodel.ErrCategoryAlreadyDeleted
	}
	
	return category, nil
}
