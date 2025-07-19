package sharedrpc

import (
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
	"github.com/pkg/errors"
	"resty.dev/v3"
)

type IntrospectRPCClient struct {
	userServiceURL string
}

func NewIntrospectRPCClient(userServiceURL string) *IntrospectRPCClient {
	return &IntrospectRPCClient{
		userServiceURL: userServiceURL,
	}
}

type dataRequester struct {
	UserID    uuid.UUID         `json:"id"`
	RoleValue datatype.UserRole `json:"role"`
}

func (d dataRequester) Subject() uuid.UUID {
	return d.UserID
}

func (d dataRequester) GetRole() datatype.UserRole {
	return d.RoleValue
}

func (c *IntrospectRPCClient) Introspect(token string) (datatype.Requester, error) {
	client := resty.New()
	
	type ResponseDTO struct {
		Data struct {
			UserID uuid.UUID         `json:"id"`
			Role   datatype.UserRole `json:"role"`
		} `json:"data"`
	}
	
	var response ResponseDTO
	
	url := c.userServiceURL + "/introspect-token"
	
	_, err := client.R().
		SetBody(map[string]interface{}{
			"token": token,
		}).
		SetResult(&response).
		Post(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	
	return &dataRequester{
		UserID:    response.Data.UserID,
		RoleValue: response.Data.Role,
	}, nil
}
