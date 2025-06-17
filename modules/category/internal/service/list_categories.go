package categoryservice

import (
	"context"
	
	"github.com/katatrina/go12-service/modules/category/internal/model"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

type IListRepo interface {
	ListCategories(
		ctx context.Context,
		pagingDTO *sharedmodel.PagingDTO,
		filterDTO *categorymodel.FilterCategoryDTO,
	) ([]categorymodel.Category, error)
}

type ListQuery struct {
	categorymodel.FilterCategoryDTO
	sharedmodel.PagingDTO
}

type ListCategoriesQueryHandler struct {
	catRepo IListRepo
}

func NewListCategoriesQueryHandler(catRepo IListRepo) *ListCategoriesQueryHandler {
	return &ListCategoriesQueryHandler{
		catRepo: catRepo,
	}
}

func (hdl *ListCategoriesQueryHandler) Execute(
	ctx context.Context,
	query *ListQuery,
) ([]categorymodel.Category, error) {
	categories, err := hdl.catRepo.ListCategories(ctx, &query.PagingDTO, &query.FilterCategoryDTO)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}
