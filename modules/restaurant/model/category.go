package restaurantmodel

import (
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type Category struct {
	ID     uuid.UUID       `json:"id" gorm:"column:id"`
	Name   string          `json:"name" gorm:"column:name"`
	Status datatype.Status `json:"status" gorm:"column:status"`
}

func (*Category) TableName() string {
	return "categories"
}
