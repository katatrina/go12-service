package sharedrpc

import (
	"context"
	"log"

	"github.com/katatrina/go12-service/gen/proto/user"
	"github.com/katatrina/go12-service/shared/datatype"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGRPCClient struct {
	client user.UserClient
	conn   *grpc.ClientConn
}

func NewUserGRPCClient(serviceURL string) *UserGRPCClient {
	conn, err := grpc.NewClient(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to User gRPC service: %v", err)
	}

	client := user.NewUserClient(conn)

	return &UserGRPCClient{
		client: client,
		conn:   conn,
	}
}

type grpcRequester struct {
	UserID    uuid.UUID         `json:"id"`
	RoleValue datatype.UserRole `json:"role"`
}

func (d grpcRequester) Subject() uuid.UUID {
	return d.UserID
}

func (d grpcRequester) GetRole() datatype.UserRole {
	return d.RoleValue
}

func (c *UserGRPCClient) Introspect(token string) (datatype.Requester, error) {
	log.Printf("Calling User gRPC service for token introspection")

	req := &user.IntrospectTokenRequest{Token: token}
	
	resp, err := c.client.IntrospectToken(context.Background(), req)
	if err != nil {
		log.Printf("Error calling User gRPC service: %v", err)
		return nil, err
	}

	userID, err := uuid.Parse(resp.Data.Id)
	if err != nil {
		return nil, err
	}

	role := datatype.UserRole(resp.Data.Role)

	return &grpcRequester{
		UserID:    userID,
		RoleValue: role,
	}, nil
}

func (c *UserGRPCClient) Close() error {
	return c.conn.Close()
}