package service

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
	"github.com/katatrina/go12-service/shared/datatype"
)

type IGetByIDRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.Category, error)
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

func (hdl *GetByIDQueryHandler) Execute(ctx context.Context, query *GetByIDQuery) (*model.Category, error) {
	category, err := hdl.catRepo.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	
	if category.Status == datatype.StatusDeleted {
		return nil, model.ErrCategoryDeleted
	}
	
	return category, nil
}
