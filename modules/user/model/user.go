package usermodel

import (
	"encoding/json"
	"time"
	
	"github.com/google/uuid"
)

type UserType string

const (
	TypeEmailPassword UserType = "email_password"
	TypeFacebook      UserType = "facebook"
	TypeGoogle        UserType = "google"
)

type UserRole string

const (
	RoleUser    UserRole = "user"
	RoleAdmin   UserRole = "admin"
	RoleShipper UserRole = "shipper"
)

type UserStatus string

const (
	StatusPending  UserStatus = "pending"
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
	StatusBanned   UserStatus = "banned"
	StatusDeleted  UserStatus = "deleted"
)

type User struct {
	ID        uuid.UUID        `json:"id" gorm:"column:id"`
	Email     string           `json:"email" gorm:"column:email"`
	Avatar    *json.RawMessage `json:"avatar" gorm:"column:avatar"`
	FirstName string           `json:"first_name" gorm:"column:first_name"`
	LastName  string           `json:"last_name" gorm:"column:last_name"`
	FbID      *string          `json:"fb_id" gorm:"column:fb_id"`
	GgID      *string          `json:"gg_id" gorm:"column:gg_id"`
	Password  string           `json:"-" gorm:"column:password"`
	Salt      string           `json:"-" gorm:"column:salt"`
	Phone     *string          `json:"phone" gorm:"column:phone"`
	Type      UserType         `json:"type" gorm:"column:type"`
	Role      UserRole         `json:"role" gorm:"column:role"`
	Status    UserStatus       `json:"status" gorm:"column:status"`
	CreatedAt time.Time        `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time        `json:"updated_at" gorm:"column:updated_at"`
}

func (u User) Subject() uuid.UUID {
	return u.ID
}

func (u User) GetRole() string {
	return string(u.Role)
}

func (User) TableName() string {
	return "users"
}
