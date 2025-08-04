package categorygrpcclient

import (
	"context"
	"log"
	
	"github.com/katatrina/go12-service/gen/proto/category"
	restaurantmodel "github.com/katatrina/go12-service/modules/restaurant/model"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryRPCClient struct {
	catGRPCServerURL string
	conn             *grpc.ClientConn
	client           category.CategoryClient
}

func NewCategoryRPCClient(catGRPCServerURL string) *CategoryRPCClient {
	conn, err := grpc.NewClient(
		catGRPCServerURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	
	if err != nil {
		log.Fatal("failed to connect to grpc server", err)
	}
	
	client := category.NewCategoryClient(conn)
	
	return &CategoryRPCClient{catGRPCServerURL: catGRPCServerURL, conn: conn, client: client}
}

func (c *CategoryRPCClient) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]restaurantmodel.Category, error) {
	strIDs := make([]string, len(ids))
	
	for i, id := range ids {
		strIDs[i] = id.String()
	}
	
	resp, err := c.client.GetCategoriesByIDs(ctx, &category.GetCatIDsRequest{Ids: strIDs})
	
	if err != nil {
		return nil, err
	}
	
	result := make([]restaurantmodel.Category, len(resp.Data))
	
	for i, cat := range resp.Data {
		result[i] = restaurantmodel.Category{
			ID:     uuid.MustParse(cat.Id),
			Name:   cat.Name,
			Status: datatype.Status(cat.Status),
		}
	}
	
	return result, nil
}
