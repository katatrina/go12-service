package categorygrpcctl

import (
	"context"
	"log"
	
	"github.com/katatrina/go12-service/gen/proto/category"
	categorymodel "github.com/katatrina/go12-service/modules/category/model"
	
	"github.com/google/uuid"
)

type CategoryRepository interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]categorymodel.Category, error)
}

type CategoryGrpcServer struct {
	category.UnimplementedCategoryServer
	repo CategoryRepository
}

func NewCategoryGrpcServer(repo CategoryRepository) *CategoryGrpcServer {
	return &CategoryGrpcServer{
		repo: repo,
	}
}

func (s *CategoryGrpcServer) GetCategoriesByIDs(
	ctx context.Context,
	req *category.GetCatIDsRequest,
) (*category.CatIDsResp, error) {
	log.Println("GetCategoriesByIDs by gRPC")
	
	uuidIds := make([]uuid.UUID, len(req.Ids))
	
	for i, id := range req.Ids {
		uuidIds[i] = uuid.MustParse(id)
	}
	
	cats, err := s.repo.FindByIDs(ctx, uuidIds)
	
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	result := make([]*category.CategoryDTO, len(cats))
	
	for i, cat := range cats {
		result[i] = &category.CategoryDTO{
			Id:     cat.ID.String(),
			Name:   cat.Name,
			Status: string(cat.Status),
		}
	}
	
	return &category.CatIDsResp{Data: result}, nil
}
