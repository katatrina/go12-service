package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IGetDetailRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
}

type GetDetailQueryHandler struct {
	catRepo IGetDetailRepo
}

func NewGetDetailQueryHandler(catRepo IGetDetailRepo) *GetDetailQueryHandler {
	return &GetDetailQueryHandler{
		catRepo: catRepo,
	}
}

type GetDetailQuery struct {
	ID uuid.UUID
}

func (hdl *GetDetailQueryHandler) Execute(ctx context.Context, query *GetDetailQuery) (*categorymodel.Category, error) {
	category, err := hdl.catRepo.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	
	if category.Status == datatype.StatusDeleted {
		return nil, categorymodel.ErrCategoryDeleted
	}
	
	return category, nil
}
