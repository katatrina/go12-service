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
	UserID    uuid.UUID `json:"id"`
	RoleValue string    `json:"role"`
}

func (r dataRequester) Subject() uuid.UUID {
	return r.UserID
}

func (r dataRequester) GetRole() string {
	return r.RoleValue
}

func (c *IntrospectRPCClient) Introspect(token string) (datatype.Requester, error) {
	client := resty.New()
	
	type ResponseDTO struct {
		Data struct {
			UserID uuid.UUID `json:"id"`
			Role   string    `json:"role"`
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
