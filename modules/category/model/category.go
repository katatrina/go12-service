package categorymodel

import (
	"strings"
	"time"
	
	"github.com/google/uuid"
	"github.com/katatrina/go12-service/shared/datatype"
)

type Category struct {
	ID          uuid.UUID       `json:"id" gorm:"column:id"`
	Name        string          `json:"name" gorm:"column:name"`
	Description *string         `json:"description" gorm:"column:description"`
	Status      datatype.Status `json:"status" gorm:"column:status"`
	CreatedAt   time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

func (dto *CreateCategoryDTO) Validate() error {
	// Introspect name
	dto.Name = strings.TrimSpace(dto.Name)
	
	if dto.Name == "" {
		return ErrNameRequired
	}
	
	if len(dto.Name) > 100 {
		return ErrInvalidNameLength
	}
	
	// TODO: Introspect description (if provided)
	
	return nil
}
