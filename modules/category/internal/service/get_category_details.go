package categoryservice

import (
	"context"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/modules/category/internal/model"
)

type GetDetailQuery struct {
	ID uuid.UUID
	// mores...
}

type GetDetailQueryHandler struct {
	catRepo ICategoryQueryRepo
}

func NewGetDetailQueryHandler(catRepo ICategoryQueryRepo) *GetDetailQueryHandler {
	return &GetDetailQueryHandler{
		catRepo: catRepo,
	}
}

func (hdl *GetDetailQueryHandler) Execute(ctx context.Context, query *GetDetailQuery) (*categorymodel.Category, error) {
	category, err := hdl.catRepo.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	
	if category.Status == categorymodel.StatusDeleted {
		return nil, categorymodel.ErrCategoryNotFound
	}
	
	return category, nil
}
