package model

import (
	"github.com/katatrina/go12-service/shared/datatype"
)

type Category struct {
	ID     string          `json:"id" gorm:"column:id"`
	Name   string          `json:"name" gorm:"column:name"`
	Icon   string          `json:"logo" gorm:"column:logo"`
	Status datatype.Status `json:"status" gorm:"column:status"`
}

func (*Category) TableName() string {
	return "categories"
}
