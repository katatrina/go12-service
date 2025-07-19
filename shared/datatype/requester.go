package datatype

import (
	"github.com/google/uuid"
)

type Requester interface {
	Subject() uuid.UUID
	GetRole() UserRole
}
