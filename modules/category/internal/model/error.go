package categorymodel

import (
	"errors"
)

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrNameRequired          = errors.New("name is required")
	ErrCategoryStatusInvalid = errors.New("status must be in (active, inactive, deleted)")
	ErrCategoryDeleted       = errors.New("category is deleted")
)
