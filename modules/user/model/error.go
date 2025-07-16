package usermodel

import (
	"errors"
)

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidFirstName   = errors.New("invalid first name")
	ErrInvalidLastName    = errors.New("invalid last name")
	
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrUserBannedOrDeleted    = errors.New("user is banned or deleted")
)
