package model

import (
	"errors"
)

var (
	ErrCategoryNotFound  = errors.New("category not found")
	ErrNameRequired      = errors.New("name is required")
	ErrInvalidNameLength = errors.New("category name must be less than 100 characters")
	ErrCategoryDeleted   = errors.New("category is deleted")
	ErrStatusInvalid     = errors.New("status is invalid, must be in [active, inactive, deleted]")
)
