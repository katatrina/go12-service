package grpcclient

import (
	"context"
	"log"

	"github.com/katatrina/go12-service/gen/proto/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantGRPCClient struct {
	client restaurant.RestaurantClient
	conn   *grpc.ClientConn
}

func NewRestaurantGRPCClient(serviceURL string) *RestaurantGRPCClient {
	conn, err := grpc.NewClient(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Restaurant gRPC service: %v", err)
	}

	client := restaurant.NewRestaurantClient(conn)

	return &RestaurantGRPCClient{
		client: client,
		conn:   conn,
	}
}

func (c *RestaurantGRPCClient) GetRestaurantsByIDs(ctx context.Context, ids []string) ([]*restaurant.RestaurantDTO, error) {
	log.Printf("Calling Restaurant gRPC service for IDs: %v", ids)

	req := &restaurant.GetRestaurantIDsRequest{Ids: ids}
	
	resp, err := c.client.GetRestaurantsByIDs(ctx, req)
	if err != nil {
		log.Printf("Error calling Restaurant gRPC service: %v", err)
		return nil, err
	}

	return resp.Data, nil
}

func (c *RestaurantGRPCClient) Close() error {
	return c.conn.Close()
}