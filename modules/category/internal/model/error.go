package categorymodel

import (
	"errors"
)

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryStatusInvalid = errors.New("category status must be in (active, inactive, deleted)")
)
