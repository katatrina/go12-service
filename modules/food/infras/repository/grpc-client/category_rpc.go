package grpcclient

import (
	"context"
	"log"

	"github.com/katatrina/go12-service/gen/proto/category"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryRPCClient struct {
	client category.CategoryClient
	conn   *grpc.ClientConn
}

func NewCategoryRPCClient(serviceURL string) *CategoryRPCClient {
	conn, err := grpc.NewClient(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Category gRPC service: %v", err)
	}

	client := category.NewCategoryClient(conn)

	return &CategoryRPCClient{
		client: client,
		conn:   conn,
	}
}

func (c *CategoryRPCClient) GetCategoriesByIDs(ctx context.Context, ids []string) ([]*category.CategoryDTO, error) {
	log.Printf("Calling Category gRPC service for IDs: %v", ids)

	req := &category.GetCatIDsRequest{Ids: ids}
	
	resp, err := c.client.GetCategoriesByIDs(ctx, req)
	if err != nil {
		log.Printf("Error calling Category gRPC service: %v", err)
		return nil, err
	}

	return resp.Data, nil
}

func (c *CategoryRPCClient) Close() error {
	return c.conn.Close()
}