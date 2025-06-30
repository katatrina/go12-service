package categorymodel

import (
	"errors"
)

var (
	ErrCategoryNotFound       = errors.New("category not found")
	ErrNameRequired           = errors.New("name is required")
	ErrInvalidNameLength      = errors.New("category name must be less than 100 characters")
	ErrCategoryAlreadyDeleted = errors.New("category is already deleted")
	ErrStatusInvalid          = errors.New("status is invalid, must be in [active, inactive, deleted]")
)
