package model

import (
	"errors"
)

var (
	ErrRestaurantNotFound       = errors.New("restaurant not found")
	ErrNameRequired             = errors.New("name is required")
	ErrInvalidNameLength        = errors.New("restaurant name must be less than 50 characters")
	ErrAddrRequired             = errors.New("address is required")
	ErrRestaurantAlreadyDeleted = errors.New("restaurant is already deleted")
	ErrStatusInvalid            = errors.New("status is invalid, must be in [active, inactive, deleted]")
)
