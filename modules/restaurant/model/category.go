package restaurantmodel

import (
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type Category struct {
	ID     uuid.UUID       `json:"id"`
	Name   string          `json:"name"`
	Status datatype.Status `json:"status"`
}

func (Category) TableName() string {
	return "categories"
}
