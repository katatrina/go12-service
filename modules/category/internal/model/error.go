package categorymodel

import (
	"errors"
)

var (
	ErrCategoryNotFound  = errors.New("category not found")
	ErrNameRequired      = errors.New("name is required")
	ErrInvalidNameLength = errors.New("category name must be less than 100 characters")
	ErrStatusInvalid     = errors.New("status must be in (active, inactive, deleted)")
	ErrCategoryDeleted   = errors.New("category is deleted")
)
