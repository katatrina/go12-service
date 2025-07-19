package sharedmodel

import (
	"errors"
)

var (
	ErrUserRoleNotAllowed = errors.New("user role not allowed to access this resource")
)
