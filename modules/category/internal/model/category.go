package categorymodel

import (
	"errors"
	"strings"
	"time"
	
	"github.com/google/uuid"
)

type Category struct {
	Id          uuid.UUID  `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	Icon        []byte     `json:"icon" gorm:"column:icon"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	
	if c.Name == "" {
		return errors.New("name is required")
	}
	
	return nil
}
