package mediamodel

import (
	"time"
	
	"github.com/google/uuid"
)

type MediaStatus string

const (
	MediaStatusPending  MediaStatus = "pending"
	MediaStatusActive   MediaStatus = "active"
	MediaStatusInactive MediaStatus = "inactive"
	MediaStatusDeleted  MediaStatus = "deleted"
)

type Media struct {
	ID        uuid.UUID   `json:"id" gorm:"column:id;"`
	URL       string      `json:"url" gorm:"-"`
	Filename  string      `json:"filename" gorm:"column:filename;"`
	CloudName string      `json:"cloudName" gorm:"column:cloud_name;"`
	Size      int64       `json:"size" gorm:"column:size;"`
	Ext       string      `json:"ext" gorm:"column:ext;"`
	Status    MediaStatus `json:"status" gorm:"column:status;"`
	CreatedAt time.Time   `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"column:updated_at;"`
}

func (Media) TableName() string {
	return "medias"
}
